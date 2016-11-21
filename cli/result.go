package cli

import (
	"fmt"
	"github.com/liuchjlu/combine/ftp"
	"github.com/liuchjlu/combine/ftpclient"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func GetResult() {
	var err error
	var ftp *goftp.FTP
	var server = SysConfig.FtpAddress
	if ftp, err = goftp.Connect(server); err != nil {
		panic(err)
	}
	defer ftp.Close()
	fmt.Println("Successfully onnected to", server)
	if err = ftp.Login(SysConfig.FtpUser, SysConfig.FtpPasswd); err != nil {
		panic(err)
	}
	// Get directory listing
	Path := TeamData.ProjectInfo.ProjectType + "/" + TeamData.TeamInfo.TeamName + "/" + TeamData.ProjectInfo.ProjectVersion
	files, err := ftp.List(Path)
	if err != nil {
		panic(err)
	}
	fmt.Println("files:", files)
	var names []string
	for _, line := range files {
		name := strings.Split(line, " ")[len(strings.Split(line, " "))-1]
		names = append(names, name)
	}
	fmt.Println(names)

	conn, err := ftpclient.DialTimeout(server, 5*time.Second)
	if err != nil {
		return
	}
	defer conn.Quit()
	err = conn.Login(SysConfig.FtpUser, SysConfig.FtpPasswd)
	if err != nil {
		return
	}
	for nm := range names {
		file := Path + "/" + nm
		data, _ := conn.Retr(file)
		buf, _ := ioutil.ReadAll(data)
		var f *os.File
		if checkFileIsExist(SysConfig.ResultAddress + "/" + nm) { //如果文件存在
			f, _ = os.OpenFile(SysConfig.ResultAddress+nm, os.O_APPEND, 0666) //打开文件
		} else {
			f, _ = os.Create(SysConfig.ResultAddress + nm) //创建文件
		}
		_, err = io.WriteString(f, string(buf))
		if err != nil {
			return
		}
	}

	//fmt.Println(string(buf))

}
