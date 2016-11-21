package img

import (
	"fmt"
	"github.com/liuchjlu/stokis/datatype"
	//"os/exec"
)

func BuildTrainyaml(teamdata datatype.Teamdata, sysconfig datatype.SysConfig) (string, error) {
	return "", nil
}

func BuildTestyaml(teamdata datatype.Teamdata, sysconfig datatype.SysConfig) (string, error) {
	fmt.Println("Before running img.BuildScript")
	fmt.Println("Teamdata:", teamdata)
	scriptpath, err := BuildScript(teamdata)
	fmt.Println("After running img.BuildScript")
	if err != nil {
		return "", err
	}
	fmt.Println("Before image build")
	imgname, err := BuildImage(teamdata, sysconfig, scriptpath)
	if err != nil {
		return "", err
	}
	fmt.Println("Before yaml build")
	yamlpath, err := BuildYaml(teamdata, imgname)
	//exec.Command("python", "img/yamlbuilder.py", "--docker_image", imgname, "--num_workers", teamdata.Frame.WorkerNums).Output()
	fmt.Println("After yaml build")
	return yamlpath, err
}
