package cli

import (
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/cmd"
)

func Livejobtorun(f func()) {
	commandname := "kubectl"
	params := []string{SysConfig.K8sServer, "get", "pods"}
	var PodLine []string
	var line string
	var live bool
	var err error
	tag := TeamData.TeamInfo.TeamName + "-" + TeamData.ProjectInfo.ProjectType + "-" + TeamData.ProjectInfo.ProjectVersion + "-0"
	for {
		live = false
		PodLine, err = cmd.GetLines(commandname, params)
		if err != nil {
                	log.Fatalf("cli.Livejobtorun(): %+v\n", err)
               		return
       		}

		for _, line = range PodLine[1:] {
			if strings.Contains(line, tag) {
				live = true
			}
		}
		go f()
		time.Sleep(1 * time.Minute)
		if !live {
			break
		}
	}
	
	return
}
