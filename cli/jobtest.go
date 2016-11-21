package cli

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/cmd"
	"github.com/liuchjlu/stokis/img"
)

func CreateTest() {
	YamlPath, err := img.BuildTestyaml(TeamData, SysConfig)
	if err != nil {
		log.Fatalf("cli.CreateTest(): %+v\n", err)
		return
	}
	commandname := "kubectl"
	params := []string{"create", "-f", YamlPath}
	fmt.Println("Before kubectl create")
	if err := cmd.ExecCommand(commandname, SysConfig.LogAddress, params); err != nil {
		log.Fatalf("cli.CreateTest(): %+v\n", err)
	}
	log.InfoLn("cli.CreateTest():Job created success!")

	Livejobtorun(Logsavetolocal)
}
