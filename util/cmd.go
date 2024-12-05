package util

import (
	"bafa/global"
	"os/exec"
	"strings"
)

func RunCommand(commandStr string) {
	parts := strings.Fields(commandStr)
	if len(parts) < 1 {
		global.SugarLogger.Error("empty command")
	}
	cmd := exec.Command(parts[0], parts[1:]...)
	global.SugarLogger.Debug("StartRun: ", cmd.String())
	err := cmd.Run()
	if err != nil {
		global.SugarLogger.Error(cmd.String(), err)
	}
	global.SugarLogger.Debug("Finish: ", cmd.String())
}
