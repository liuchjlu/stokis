package img

import (
	"github.com/liuchjlu/stokis/datatype"
	log "github.com/Sirupsen/logrus"
	//"os/exec"
)

func BuildTrainyaml(teamdata datatype.Teamdata, sysconfig datatype.SysConfig) (string, error) {
	return "", nil
}

func BuildTestyaml(teamdata datatype.Teamdata, sysconfig datatype.SysConfig) (string, error) {
	log.Debugf("Before running img.BuildScript")
	scriptpath, err := BuildScript(teamdata, sysconfig)
	log.Debugf("After running img.BuildScript")
	if err != nil {
		return "", err
	}
	log.Debugf("Before image build")
	imgname, err := BuildImage(teamdata, sysconfig, scriptpath)
	if err != nil {
		return "", err
	}
	log.Debugf("Before yaml build")
	yamlpath, err := BuildYaml(teamdata, imgname)
	//exec.Command("python", "img/yamlbuilder.py", "--docker_image", imgname, "--num_workers", teamdata.Frame.WorkerNums).Output()
	log.Debugf("After yaml build")
	return yamlpath, err
}
