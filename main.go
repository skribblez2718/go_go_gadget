package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	flag.Usage = displayUsage
	flag.Parse()

	args := getArgs()
	command := args[0]
	logPath := args[1]

	cmd := getCommand(command)
	tempFile := getTempFile(command)

	tabWriter := &TabWriter{Writer: tempFile}
	multiWriter := io.MultiWriter(os.Stdout, tabWriter)

	fmt.Fprint(tempFile, getLogEntryHeader(command))

	startTime := time.Now()

	cmd.Stdout = multiWriter

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(tempFile, getLogEntryFooter(startTime))
	tempFile.Close()

	waitForLogFileHandle(logPath, tempFile.Name())

	tempFileContent := getTempFileContent(tempFile.Name())
	logFile := getLogFile(logPath)
	writeToLogFile(tempFileContent, logFile, tempFile.Name())
	os.Remove(tempFile.Name())
}
