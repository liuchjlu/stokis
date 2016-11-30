package datatype

import (
	// "fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	// "reflect"
)

type SysConfig struct {
	K8sServer     	  string `K8sServer`
	EvaluateServer    string `EvaluateServer`
	FtpAddress        string `FtpAddress`
	FtpUser           string `FtpUser`
	FtpPasswd     	  string `FtpPasswd`
	FtpAddressBak     string `FtpAddressBak`
        FtpUserBak        string `FtpUserBak`
        FtpPasswdBak      string `FtpPasswdBak`
	ImgRepository     string `ImgRepository`
	LogSize           string `LogSize` //单位为Mb
	PodScriptPath     string `PodScriptPath`
	ResultAddress     string `ResultAddress`
	PodLogAddress     string `PodLogAddress`
}
type UserConfig struct {
	MaxCommitNum string `MaxCommitNum`
	Cpu          string `Cpu`
	Memory       string `Memory` //单位为Gb
}
type Teamdata struct {
	TeamInfo    TeamInfo    `TeamInfo`
	ProjectInfo ProjectInfo `ProjectInfo`
	Frame       Frame       `Frame`
}
type TeamInfo struct {
	TeamID   string `TeamID`
	TeamName string `TeamName`
}
type ProjectInfo struct {
	ProjectVersion string `ProjectVersion`
	DataVersion    string `DataVersion`
	ProjectType    string `ProjectType`
	WorkPath       string `WorkPath`
	ScriptPath     string `ScriptPath`
	ResultPath     string `ResultPath`
	OS             string `OS`
}
type Frame struct {
	FrameType   string `FrameType`
	ComputeType string `ComputeType` //gpu/cpu
	WorkerNums  int    `WorkerNums`
	PsNums      int    `PsNums`
}

type Evaluate struct {
	UserName       string
	ProgramName    string
	ResultAddress  string
	DataVersion    string
	ProjectVersion string
	DetectTime     string
}

/* 数据说明：
   TeamID                  参赛队伍唯一id
   TeamName                参赛英文队名
   ProjectType             参与项目类型名(人脸识别：faceRecognition
                                          视频拷贝检测：videoCopydetection
                                          视频文本关键词检测：videoKeywords
                                          说话人识别：speakerRecognition
                                          特定视频识别：videoRecognition
                                          音频比对：audioCopydetection
                                          音频文本关键词：audioKeywords
                                          )
   WorkPath                项目程序包路径
   StartScriptPath         启动入口脚本
   GpuNums                 0/1,0不使用，1使用
*/

func UnmarshalTeamdata(path string, data Teamdata) (Teamdata, error) {
	in, err := ioutil.ReadFile(path)
	if err != nil {
		return data, err
	}
	err = yaml.Unmarshal(in, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func UnmarshalSysconfig(path string, data SysConfig) (SysConfig, error) {
	in, err := ioutil.ReadFile(path)
	if err != nil {
		return data, err
	}
	err = yaml.Unmarshal(in, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func UnmarshalUserconfig(path string, data UserConfig) (UserConfig, error) {
	in, err := ioutil.ReadFile(path)
	if err != nil {
		return data, err
	}
	err = yaml.Unmarshal(in, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
