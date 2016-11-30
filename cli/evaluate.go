package cli

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/datatype"
)

func Anyresult() {

}
func Evaluate() {
	result := datatype.Evaluate{
		TeamData.TeamInfo.TeamName,
		TeamData.ProjectInfo.ProjectType,
		SysConfig.FtpAddress+ "/" + TeamData.ProjectInfo.ProjectType + "/" + TeamData.TeamInfo.TeamName + "/" + TeamData.ProjectInfo.ProjectVersion + "/result.txt",
		"v1",
		TeamData.ProjectInfo.ProjectVersion,
		"0:0:0",
	}
	fmt.Println(result)
	jsonresult, err := json.Marshal(result)
	url := SysConfig.EvaluateServer
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonresult))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Infoln("response Status:", resp.Status)
	log.Infoln("response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)

	log.Infoln("The evaluate result:", string(body))
}
