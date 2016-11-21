package img

import (
	"os"
	"text/template"

	"github.com/liuchjlu/stokis/datatype"
)

func BuildYaml(teamdata datatype.Teamdata, imgname string) (string, error) {
	filepath := "/.stokis/" + teamdata.ProjectInfo.ProjectType + "/" + teamdata.TeamInfo.TeamName + "/"
	os.MkdirAll(teamdata.ProjectInfo.WorkPath+filepath, 0755)
	filename := teamdata.ProjectInfo.WorkPath + filepath + "job.yaml"
	dstFile, err := os.Create(filename)
	defer dstFile.Close()
	check(err)
	for index := 0; index < teamdata.Frame.WorkerNums; index++ {
		t, err := template.New("job.template").ParseFiles("template/job.template")
		check(err)

		data := struct {
			WorkerID    int
			ClusterSpec string
			DockerImage string
			ProjectType string
			TeamName    string
		}{
			WorkerID:    index,
			ClusterSpec: "",
			DockerImage: imgname,
			ProjectType: teamdata.ProjectInfo.ProjectType,
			TeamName:    teamdata.TeamInfo.TeamName,
		}
		err = t.Execute(dstFile, data)
		check(err)
		dstFile.WriteString("---------\n")
	}
	return filename, nil
}
