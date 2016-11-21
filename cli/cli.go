package cli

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/datatype"
)

/*
manage train yamlpath
manage test yamlpath
manage evaluate yamlpath
manage getlog yamlpath
*/
var TeamData datatype.Teamdata = datatype.Teamdata{}
var SysConfig datatype.SysConfig = datatype.SysConfig{}
var UserConfig datatype.UserConfig = datatype.UserConfig{}
var YamlPath string

func Run() {
	if len(os.Args) <= 2 || os.Args[1] == "help" || len(os.Args) > 3 {
		help()
		return
	}
	var err error
	if TeamData, err = datatype.UnmarshalTeamdata(os.Args[2], TeamData); err != nil {
		log.Errorf("cli.Run():%+v\n", err)
	}
	if SysConfig, err = datatype.UnmarshalSysconfig("config/sysconfig.yaml", SysConfig); err != nil {
		log.Errorf("cli.Run():%+v\n", err)
	}
	if UserConfig, err = datatype.UnmarshalUserconfig("config/userconfig.yaml", UserConfig); err != nil {
		log.Errorf("cli.Run():%+v\n", err)
	}

	log.Debugf("cli.Run(): cli args:%+v\n", os.Args)
	log.Debugf("teamdata: %+v %+v %+v", TeamData, SysConfig, UserConfig)
	switch command := os.Args[1]; command {
	case "train":
		CreateTrain()
	case "test":
		CreateTest()
	case "evaluate":
		Evaluate()
	case "getresult":
		GetResult()
	case "destroy":
		DeleteJob()
	default:
		help()
	}
}
