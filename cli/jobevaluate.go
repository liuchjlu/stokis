package cli

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/datatype"
)

func Anyresult() {

}
func Evaluate() {
	result := datatype.Evaluate{
		TeamData.TeamInfo.TeamName,
		TeamData.ProjectInfo.ProjectType,
		"192.168.12.41:21/result/result.txt",
		"V1",
		"V1",
		"7:12:34",
	}
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
