FROM 192.168.11.51:5000/tensorflow-gpu:centos7

MAINTAINER Aaron liu "254311953@qq.com"

RUN mkdir -p /user/lib/golang/src/github.com/liuchjlu/stokis/test:
COPY ./ /user/lib/golang/src/github.com/liuchjlu/stokis/test
COPY .stokis/facerecognition/liucaihong/podscript.sh  /start.sh

RUN chmod a+x /start.sh 
ENTRYPOINT [/start.sh]
