package cli

import (
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/cmd"
)

func Livejobtorun(f func()) {
	commandname := "kubectl"
	params := []string{"get", "pods"}
	PodLine, err := cmd.GetLines(commandname, params)
	if err != nil {
		log.Fatalf("cli.Livejobtorun(): %+v\n", err)
		return
	}
	var live bool = false
	tag := TeamData.ProjectInfo.ProjectType + "-" + TeamData.TeamInfo.TeamName + "-" + TeamData.ProjectInfo.ProjectVersion
	for {
		for _, line := range PodLine[1:] {
			if strings.Contains(line, tag) {
				live = true
				go f()
			}
		}
		if !live {
			break
		}
		time.Sleep(30 * time.Second)
	}
	return
}
