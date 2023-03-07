package shcommands

import "os/exec"

func ShScriptLogUpdate() {
	cmdStr := "./log.sh "
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}
}
