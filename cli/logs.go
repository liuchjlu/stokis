package cli

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/cmd"
)

func SimpleSimplifyWhitespace(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}

func Logsavetolocal() {
	var err error
	commandname := "kubectl"
	params1 := []string{SysConfig.K8sServer, "get", "pods", "-a"}
	PodLine, err := cmd.GetLines(commandname, params1)
	if err != nil {
		log.Fatalf("cli.logsavetolocal(): %+v\n", err)
		return
	}
	var line string
	var params2 []string
	var rows []string
	tag := TeamData.TeamInfo.TeamName + "-" + TeamData.ProjectInfo.ProjectType + "-" + TeamData.ProjectInfo.ProjectVersion
	for _, line = range PodLine {
		if strings.Contains(line, tag) {
			rows = strings.Split(SimpleSimplifyWhitespace(line), " ")
			params2 = []string{SysConfig.K8sServer, "logs", rows[0]}
			if err = cmd.ExecCommand(commandname, SysConfig.PodLogAddress+rows[0]+".log", params2); err != nil {
				log.Fatalf("cli.logsavetolocal(): %+v\n", err)
			}
		}
	}
	return
}
func Logs() {
	var err error
	commandname := "kubectl"
        params1 := []string{SysConfig.K8sServer, "get", "pods", "-a"}
        PodLine, err := cmd.GetLines(commandname, params1)
        if err != nil {
                log.Fatalf("cli.logsavetolocal(): %+v\n", err)
                return
        }
	var line string
        var l string 
        var params2 []string
        var rows []string
	tag := TeamData.TeamInfo.TeamName + "-" + TeamData.ProjectInfo.ProjectType + "-" + TeamData.ProjectInfo.ProjectVersion
	for _, line = range PodLine {
                if strings.Contains(line, tag) {
                        rows = strings.Split(SimpleSimplifyWhitespace(line), " ")
                        params2 = []string{SysConfig.K8sServer, "logs", rows[0]}
                        if err = cmd.ExecCommand(commandname, SysConfig.PodLogAddress+rows[0]+".log", params2); err != nil {
                                log.Fatalf("cli.logsavetolocal(): %+v\n", err)
                        }
			l = SimpleSimplifyWhitespace(line)
		        log.Infoln("cli.Logs(): Logs of your program has been saved to /stokis/podlogs/"+SimpleSimplifynil(strings.Split(l," ")[0])+".log from standard output.")
                }
        }
	return
}
