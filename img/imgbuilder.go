package img

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"

	"github.com/liuchjlu/stokis/datatype"
)

func BuildImage(teamdata datatype.Teamdata, sysconfig datatype.SysConfig, scriptpath string) (imgname string, err error) {
	dockerfile, err := buildDockerfile(teamdata, sysconfig, scriptpath)
	baseImgName := sysconfig.ImgRepository + "/" + teamdata.Frame.FrameType + "-" + teamdata.Frame.ComputeType + ":" + teamdata.ProjectInfo.OS
	imgname = sysconfig.ImgRepository + "/" + teamdata.TeamInfo.TeamName + ":" + teamdata.ProjectInfo.ProjectType
	fmt.Printf("baseImgName:%+v dockerfile:%+v\n", baseImgName, dockerfile)
	out, err := exec.Command("docker", "pull", baseImgName).Output()
	fmt.Println(string(out), err)
	fmt.Printf("Before build %+v \n", imgname)
	out, err = exec.Command("docker", "build", "-t", imgname, dockerfile).Output()
	fmt.Println(string(out), err)
	exec.Command("docker", "push", imgname).Run()
	return imgname, err
}

func buildDockerfile(teamdata datatype.Teamdata, sysconfig datatype.SysConfig, scriptpath string) (dockerfile string, err error) {
	filepath := teamdata.ProjectInfo.WorkPath
	fileName := filepath + "Dockerfile"
	dstFile, err := os.Create(fileName)
	defer dstFile.Close()
	check(err)

	t, err := template.New("dockerfile.template").ParseFiles("template/dockerfile.template")
	check(err)

	data := struct {
		Framework     string
		Genre         string
		Os            string
		Repository    string
		Workpath      string
		Scriptpath    string
		PodScriptpath string
	}{
		Framework:     teamdata.Frame.FrameType,
		Genre:         teamdata.Frame.ComputeType,
		Os:            teamdata.ProjectInfo.OS,
		Repository:    sysconfig.ImgRepository,
		Workpath:      teamdata.ProjectInfo.WorkPath,
		Scriptpath:    scriptpath,
		PodScriptpath: sysconfig.PodScriptPath,
	}
	fmt.Printf("dockerfile data %+v\n", data)
	err = t.Execute(dstFile, data)
	check(err)
	return filepath, nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}