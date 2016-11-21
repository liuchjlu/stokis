from __future__ import absolute_import
from __future__ import division
from __future__ import print_function

import argparse
import sys

DEFAULT_DOCKER_IMAGE = '{.cmd.Exec[BuildImage(teamdata)]}'
DEFAULT_PORT = 2222

WORKER_JOB = (
    """apiVersion: batch/v1
kind: Job
metadata:
  name: tf-worker{worker_id}
spec:
  template:
    spec:
      containers:
      - name: tf-worker{worker_id}
        image: {docker_image}
        args:
          - --cluster_spec={cluster_spec}
          - --job_name=worker
          - --task_id={worker_id}
        ports:
        - containerPort: {port}
        imagePullPolicy: IfNotPresent
        command: ["/bin/bash", "/start.sh"]
        volumeMounts:
        - name: shared
          mountPath: /shared
      volumes:
      - name: shared
        hostPath:
          path: /shared
      restartPolicy: Never
      nodeSelector:
        kubernetes.io/hostname: 192.168.11.53
""")

WORKER_SVC = (
    """apiVersion: v1
kind: Service
metadata:
  name: tf-worker{worker_id}
  labels:
    tf-worker: "{worker_id}"
spec:
  ports:
  - port: {port}
    targetPort: {port}
  selector:
    tf-worker: "{worker_id}"
""")
WORKER_LB_SVC = (
    """apiVersion: v1
kind: Service
metadata:
  name: tf-worker{worker_id}
  labels:
    tf-worker: "{worker_id}"
spec:
  type: LoadBalancer
  ports:
  - port: {port}
  selector:
    tf-worker: "{worker_id}"
""")

def main():
  """Do arg parsing."""
  parser = argparse.ArgumentParser()
  parser.add_argument('--num_workers',
                      type=int,
                      default=2)
  parser.add_argument('--grpc_port',
                      type=int,
                      default=DEFAULT_PORT)
  parser.add_argument('--request_load_balancer',
                      type=bool,
                      default=False,
                      help='To request worker0 to be exposed on a public IP '
                      'address via an external load balancer, enabling you to '
                      'run client processes from outside the cluster')
  parser.add_argument('--docker_image',
                      type=str,
                      default=DEFAULT_DOCKER_IMAGE)
  args = parser.parse_args()

  if args.num_workers <= 0:
    sys.stderr.write('--num_workers must be greater than 0; received %d\n'
                     % args.num_workers)
    sys.exit(1)

  # Generate contents of yaml config
  yaml_config = GenerateConfig(args.num_workers,
                               args.grpc_port,
							     args.request_load_balancer,
                               args.docker_image)
  f = open('./job.yaml', 'w')
  f.write(yaml_config)
  f.close()

def GenerateConfig(num_workers,
                   port,
				   request_load_balancer,
                   docker_image):
  """Generate configuration strings."""
  config = ''
  for worker in range(num_workers):
    config += WORKER_JOB.format(
        port=port,
        worker_id=worker,
        docker_image=docker_image,
        cluster_spec=WorkerClusterSpecString(num_workers,
                                             port))
				
    config += '------------------------------\n'
    if request_load_balancer:
      config += WORKER_LB_SVC.format(port=port,
                                     worker_id=worker)
    else:
      config += WORKER_SVC.format(port=port,
                                  worker_id=worker)
      config += '-----------------------------\n'
  return config

def ClusterSpecString(num_workers,
                      port):
  """Generates general cluster spec."""
  spec = 'worker|'
  for worker in range(num_workers):
    spec += 'tf-worker%d:%d' % (worker, port)
    if worker != num_workers-1:
      spec += ';'
  return spec
  
def WorkerClusterSpecString(num_workers,
                            port):
  """Generates worker cluster spec."""
  return ClusterSpecString(num_workers, port)

if __name__ == '__main__':
    main()