package img

import (
	"os"
	"strings"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/datatype"
)

func BuildScript(teamdata datatype.Teamdata, sysconfig datatype.SysConfig) (scriptpath string, err error) {
	filepath := ".stokis/" + teamdata.ProjectInfo.ProjectType + "/" + teamdata.TeamInfo.TeamName + "/"
	os.MkdirAll(teamdata.ProjectInfo.WorkPath + "/" + filepath, 0755)
	fileName := teamdata.ProjectInfo.WorkPath + "/" + filepath + "podscript.sh"
	dstFile, err := os.Create(fileName)
	defer dstFile.Close()
	check(err)
	t, err := template.New("podscript.template").ParseFiles("/etc/.config/template/podscript.template")
	check(err)
	ftpserver := strings.Split(sysconfig.FtpAddress, ":")[0]
	ftpserverbak := strings.Split(sysconfig.FtpAddressBak, ":")[0]

	data := struct {
		ScriptPath     string
		ProjectType    string
		TeamName       string
		UserName       string
		PassWord       string
		FtpServer      string
                UserNameBak    string
                PassWordBak    string
                FtpServerBak   string
		Version        string
		ResultPath     string
	}{
		ScriptPath:     teamdata.ProjectInfo.ScriptPath,
		ProjectType:    teamdata.ProjectInfo.ProjectType,
		TeamName:       teamdata.TeamInfo.TeamName,
		Version:        teamdata.ProjectInfo.ProjectVersion,
		UserName:       sysconfig.FtpUser,
		PassWord:       sysconfig.FtpPasswd,
		FtpServer:      ftpserver,
                UserNameBak:    sysconfig.FtpUserBak,
                PassWordBak:    sysconfig.FtpPasswdBak,
                FtpServerBak:   ftpserverbak,
		ResultPath:	teamdata.ProjectInfo.ResultPath,
	}
	err = t.Execute(dstFile, data)
	check(err)
	log.Debugf("ScriptFilePath:"+filepath + "podscript.sh")
	return filepath + "podscript.sh", nil
}
