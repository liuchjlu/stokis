package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/cmd"
	"github.com/liuchjlu/stokis/img"
)

func CreateTrain() {
	YamlPath, err := img.BuildTrainyaml(TeamData, SysConfig)
	if err != nil {
		log.Fatalf("cli.CreateTrain(): %+v\n", err)
		return
	}
	commandname := "kubectl"
	params := []string{"create", "-f", YamlPath}
	if err := cmd.ExecCommand(commandname, SysConfig.LogAddress, params); err != nil {
		log.Fatalf("cli.CreateTrain(): %+v\n", err)
	}

	Livejobtorun(Logsavetolocal)

}
