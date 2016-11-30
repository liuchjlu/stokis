package cli

import (
/*
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/cmd"
	"github.com/liuchjlu/stokis/global"
	"github.com/liuchjlu/stokis/img"
*/
	log "github.com/Sirupsen/logrus"
)

func CreateTrain() {
/*
	Clear()
	YamlPath, err := img.BuildTrainyaml(TeamData, SysConfig)
	if err != nil {
		log.Fatalf("cli.CreateTrain(): %+v\n", err)
		return
	}
	commandname := "kubectl"
	params := []string{SysConfig.K8sServer, "create", "-f", YamlPath}
	if err := cmd.ExecCommand(commandname, global.Logpath, params); err != nil {
		log.Fatalf("cli.CreateTrain(): %+v\n", err)
	}
        log.Infoln("Your application is running in daemon.")
        log.Infoln("Please don't quit!  Logs is synchronizing from standard output... ")
	
	time.Sleep(3*time.Minute)

	Livejobtorun(Logsavetolocal)
        log.Infoln("The test option is completed.")
*/
	log.Infoln("The modle-train option is not provide.")
	return

}
