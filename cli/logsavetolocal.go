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
	commandname := "kubectl"
	params := []string{"get", "pods"}
	PodLine, err := cmd.GetLines(commandname, params)
	if err != nil {
		log.Fatalf("cli.logsavetolocal(): %+v\n", err)
		return
	}

	tag := TeamData.ProjectInfo.ProjectType + "-" + TeamData.TeamInfo.TeamName + "-" + TeamData.ProjectInfo.ProjectVersion
	for _, line := range PodLine {
		if strings.Contains(line, tag) {
			rows := strings.Split(SimpleSimplifyWhitespace(line), " ")
			if err := cmd.ExecCommand("kubectl", SysConfig.LogAddress, rows); err != nil {
				log.Fatalf("cli.logsavetolocal(): %+v\n", err)
			}
		}
	}

	return
}
