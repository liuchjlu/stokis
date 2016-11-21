package img

import (
	"fmt"
	"os"
	"text/template"

	"github.com/liuchjlu/stokis/datatype"
)

func BuildScript(teamdata datatype.Teamdata) (scriptpath string, err error) {
	filepath := "/.stokis/" + teamdata.ProjectInfo.ProjectType + "/" + teamdata.TeamInfo.TeamName + "/"
	os.MkdirAll(teamdata.ProjectInfo.WorkPath+filepath, 0755)
	fileName := teamdata.ProjectInfo.WorkPath + filepath + "podscript.sh"
	dstFile, err := os.Create(fileName)
	defer dstFile.Close()
	check(err)
	t, err := template.New("podscript.template").ParseFiles("template/podscript.template")
	check(err)

	data := struct {
		ScriptPath  string
		ProjectType string
		TeamName    string
		UserName    string
		PassWord    string
		FtpServer   string
		Version     string
	}{
		ScriptPath:  teamdata.ProjectInfo.ScriptPath,
		ProjectType: teamdata.ProjectInfo.ProjectType,
		TeamName:    teamdata.TeamInfo.TeamName,
		Version:     teamdata.ProjectInfo.ProjectVersion,
		UserName:    "stokis",
		PassWord:    "111111",
		FtpServer:   "192.168.11.51",
	}
	err = t.Execute(dstFile, data)
	check(err)
	fmt.Println(fileName)
	return filepath + "podscript.sh", nil
}
