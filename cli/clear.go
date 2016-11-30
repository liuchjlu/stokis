package cli

import (
	"github.com/liuchjlu/stokis/global"
	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/cmd"
)

func Clear() {
	log.Infoln("cli.Clear(): Start checking and clearing the old jobs.")
	tag := TeamData.TeamInfo.TeamName + "-" + TeamData.ProjectInfo.ProjectType + "-" + TeamData.ProjectInfo.ProjectVersion + "-0"
	log.Infoln("tag:",tag)
	commandname := "kubectl"
	params := []string{SysConfig.K8sServer, "delete", "jobs", tag}
	if err := cmd.ExecCommand(commandname, global.Logpath, params); err != nil {
		log.Fatalf("cli.CreateTrain(): %+v\n", err)
	}
        log.Infoln("cli.Clear(): Clear completed.")
	return

}
