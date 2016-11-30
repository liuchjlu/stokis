package img

import (
	"os"
	"text/template"

	"github.com/liuchjlu/stokis/datatype"
)

func BuildYaml(teamdata datatype.Teamdata, imgname string) (string, error) {
	filepath := ".stokis/" + teamdata.ProjectInfo.ProjectType + "/" + teamdata.TeamInfo.TeamName + "/"
	os.MkdirAll(teamdata.ProjectInfo.WorkPath+"/"+filepath, 0755)
	filename := teamdata.ProjectInfo.WorkPath + "/" + filepath + "job.yaml"
	dstFile, err := os.Create(filename)
	defer dstFile.Close()
	check(err)
	var gpunm int
	for index := 0; index < teamdata.Frame.WorkerNums; index++ {
		t, err := template.New("job.template").ParseFiles("/etc/.config//template/job.template")
		check(err)
		if teamdata.Frame.ComputeType == "gpu" {
			gpunm = 1
		} else {
			gpunm = 0
		}

		data := struct {
			WorkerID       int
			ClusterSpec    string
			DockerImage    string
			ProjectType    string
			ProjectVersion string
			TeamName       string
			GpuNm          int
			DataVersion    string
			ResultPath     string
		}{
			WorkerID:       index,
			ClusterSpec:    "",
			DockerImage:    imgname,
			ProjectType:    teamdata.ProjectInfo.ProjectType,
			ProjectVersion: teamdata.ProjectInfo.ProjectVersion,
			TeamName:       teamdata.TeamInfo.TeamName,
			GpuNm:          gpunm,
			DataVersion:    teamdata.ProjectInfo.DataVersion,
			ResultPath:     teamdata.ProjectInfo.ResultPath,
		}
		err = t.Execute(dstFile, data)
		check(err)
		dstFile.WriteString("\n")
		dstFile.WriteString("---\n")
	}
	return filename, nil
}

