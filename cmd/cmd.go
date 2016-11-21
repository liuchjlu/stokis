package cmd

import (
	"bufio"
	"io"
	"os"
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

func ExecCommand(commandName string, logpath string, params []string) error {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	//fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatalf("cmd.ExecCommand():%+v\n", err)
		return err
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	fo, err := os.Create(logpath)
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	writer := bufio.NewWriter(fo) //创建输出缓冲流

	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if n2, err := writer.Write(buf[:n]); err != nil {
			panic(err)
		} else if n2 != n {
			panic("error in writing")
		}
	}

	if err = writer.Flush(); err != nil {
		panic(err)
	}

	// //实时循环读取输出流中的一行内容
	// for {
	// 	line, err2 := reader.ReadString('\n')
	// 	if err2 != nil || io.EOF == err2 {
	// 		break
	// 	}

	// 	n, err := writer.WriteString(line)
	// 	if err != nil || n == -1{
	// 		panic("error in writing")
	// 	}

	// 	fmt.Println(line)
	// }

	cmd.Wait()
	return nil
}

func GetLines(commandName string, params []string) ([]string, error) {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	//fmt.Println(cmd.Args)
	var Line []string
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatalf("cmd.ExecCommand():%+v\n", err)
		return Line, err
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		Line = append(Line, line)

	}
	cmd.Wait()
	return Line, nil
}
