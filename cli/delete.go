package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/cmd"
	"github.com/liuchjlu/stokis/global"
)

func DeleteJob() {
	commandname := "kubectl"
	params := []string{SysConfig.K8sServer, "delete", "jobs", TeamData.TeamInfo.TeamName + "-" + TeamData.ProjectInfo.ProjectType + "-" + TeamData.ProjectInfo.ProjectVersion + "-0"}
	if err := cmd.ExecCommand(commandname, global.Logpath, params); err != nil {
		log.Fatalf("cli.CreateTrain(): %+v\n", err)
	}
	log.Infoln("The Current job has bean deleted.")
	return

}
