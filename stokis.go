// manage project manage.go
package main

import (
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/liuchjlu/stokis/cli"
)

func main() {
	logFilename := "/tmp/stokis.log"
	logFile, _ := os.OpenFile(logFilename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer logFile.Close()

	writers := []io.Writer{
		logFile,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)

	log.SetOutput(fileAndStdoutWriter)
	log.SetLevel(log.DebugLevel)

	log.Infoln("main.main():Start Stokis Main")

	cli.Run()
}
