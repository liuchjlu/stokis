package cli

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/ftp"
	"github.com/liuchjlu/stokis/ftpclient"
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
func SimpleSimplifynil(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), "")
}
func GetResult() {
	var err error
	var ftp *goftp.FTP
	var server = SysConfig.FtpAddress
	if ftp, err = goftp.Connect(server); err != nil {
		panic(err)
	}
	defer ftp.Close()
	log.Infoln("build connection success.")
	if err = ftp.Login(SysConfig.FtpUser, SysConfig.FtpPasswd); err != nil {
		panic(err)
	}
	// Get directory listing
	Path := TeamData.ProjectInfo.ProjectType + "/" + TeamData.TeamInfo.TeamName + "/" + TeamData.ProjectInfo.ProjectVersion
	files, err := ftp.List(Path)
	if err != nil {
		panic(err)
	}
	var names []string //list of line
	for _, line := range files {
		name := strings.Split(line, " ")[len(strings.Split(line, " "))-1]
		names = append(names, name)
	}

	var buf []byte
	var file string
	var data io.Reader
	var conn *ftpclient.ServerConn
	var f *os.File

	for _, nm := range names {
		nm = SimpleSimplifynil(nm)
		file = Path + "/" + nm
		//build connection to ftp
		conn, err = ftpclient.DialTimeout(server, 5*time.Second)
		if err != nil {
			return
		}
		err = conn.Login(SysConfig.FtpUser, SysConfig.FtpPasswd)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Quit()
		//get data from ftp
		data, _ = conn.Retr(file)
		buf, _ = ioutil.ReadAll(data)
		if checkFileIsExist(SysConfig.ResultAddress + nm) {
			//	fmt.Println("file exist")
			f, _ = os.OpenFile(SysConfig.ResultAddress+nm, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		} else {
			//	fmt.Println("file don't exist,creating.")
			f, _ = os.Create(SysConfig.ResultAddress + nm)
		}
		_, err = io.WriteString(f, string(buf))
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Close()
	}
	log.Infoln("cli.Getresult(): Result completed download with the path:", SysConfig.ResultAddress)
	return
}
