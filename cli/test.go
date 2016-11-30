package cli

import (
	"time"
	//"os"

	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/cmd"
	"github.com/liuchjlu/stokis/global"
	"github.com/liuchjlu/stokis/img"
)

func CreateTest() {
	Clear()
	log.Infoln("cli.CreateTest(): The test option is running.")
	YamlPath, err := img.BuildTestyaml(TeamData, SysConfig)
	log.Debugf("cli.CreateTest(): YamlPath=:", YamlPath)
	if err != nil {
		log.Fatalf("cli.CreateTest(): %+v\n", err)
		return
	}
	commandname := "kubectl"
	params := []string{SysConfig.K8sServer, "create", "-f", YamlPath}
	log.Infoln("Start creating jobs.")
	if err := cmd.ExecCommand(commandname, global.Logpath, params); err != nil {
		log.Fatalf("cli.CreateTest(): Job create faild. %+v\n", err)
	}
//	os.RemoveAll(TeamData.ProjectInfo.WorkPath + ".stokis")
	log.Infoln("cli.CreateTest():Job created success!")
        log.Infoln("Your program is running in daemon.")
        log.Infoln("Please don't quit!  Logs is synchronizing from standard output... ")
	time.Sleep(3*time.Minute)
//	t1 := time.Now()
	Livejobtorun(Logsavetolocal)
	log.Infoln("The option of model-test is completed.")
//	elspsed := time.Since(t1)
//	log.Infoln("Total time:", elspsed)
	return 
}
