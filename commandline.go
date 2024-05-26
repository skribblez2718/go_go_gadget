package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func getArgs() []string {
	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		os.Exit(0)
	}

	return args
}

func getCommand(command string) *exec.Cmd {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	return cmd
}

func displayUsage() {
	usageMessage := "Description:\n"
	usageMessage += "\tRuns a command with sh or cmd.exe while logging the start and stop time of the command. Command output with time stamps is then written to specified file.\n"
	usageMessage += "Usage:\n"
	usageMessage += "\tgo_go_gadget \"<command>\" <output_file>\n"

	fmt.Print(usageMessage)
}
