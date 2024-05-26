package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type TabWriter struct {
	Writer io.Writer
}

func (tw *TabWriter) Write(p []byte) (n int, err error) {
	lines := strings.Split(string(p), "\n")
	for _, line := range lines {
		if len(line) > 0 && !strings.HasPrefix(line, "#") {
			line = "\t" + line
		}
		_, err = tw.Writer.Write([]byte(line + "\n"))
		if err != nil {
			return n, err
		}
	}
	return len(p), nil
}

func getLogEntryHeader(command string) string {
	header := fmt.Sprintf("%s%s", strings.Repeat("-", 100), "\n")
	header += fmt.Sprintf("- Command: %s\n", command)
	header += "- Output:\n"

	return header
}

func getTempFile(command string) *os.File {
	commandParts := strings.Split(command, " ")
	commandName := commandParts[0]
	scrubbedCommandName := scrubCommandName(commandName)
	fileName := fmt.Sprintf("%s-*", scrubbedCommandName)

	f, err := os.CreateTemp("", fileName)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func scrubCommandName(commandName string) string {
	var scrubbedCommandName string
	if runtime.GOOS == "windows" {
		scrubbedCommandName = strings.ReplaceAll(commandName, "\\", "_")
	} else {
		scrubbedCommandName = strings.ReplaceAll(commandName, "/", "_")
	}

	return scrubbedCommandName
}

func getLogEntryFooter(startTime time.Time) string {
	endTime := time.Now()
	footer := fmt.Sprintf("- Start time: %s\n", startTime.Format(time.RFC3339))
	footer += fmt.Sprintf("- End time: %s\n", endTime.Format(time.RFC3339))
	footer += fmt.Sprintf("%s%s", strings.Repeat("-", 100), "\n")

	return footer
}

func waitForLogFileHandle(logFile string, tempFileName string) {
	for {
		if !isLogFileOpen(logFile) {
			break
		}
		fmt.Printf("\033[1;33mLog file currently in use. Log entry can be found at %s. Retrying in 10 seconds...\033[0m", tempFileName)
		time.Sleep(time.Second * 10)
	}
}

func isLogFileOpen(fileName string) bool {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return true
	}
	file.Close()

	return false
}

func getTempFileContent(fileName string) []byte {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err, "Command output located at ", fileName)
	}

	return content
}

func getLogFile(fileName string) *os.File {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("\033[1;33mLog filel currently in use. Retrying in 10 seconds...\033[0m")
		time.Sleep(time.Second * 10)
	}

	return f
}

func writeToLogFile(content []byte, logFile *os.File, tempFileName string) {
	if _, err := logFile.Write(content); err != nil {
		log.Fatal(err, "Command output located at ", tempFileName)
	}
	defer logFile.Close()
}
