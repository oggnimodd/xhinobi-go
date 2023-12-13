package helpers

import (
	"os/exec"
	"xhinobi-go/constants"
)

func OpenTempFileInCode(file string) (*exec.Cmd, error) {
	var cmd *exec.Cmd
	if constants.IsGoogleCloud {
		cmd = exec.Command("cloudshell", "open", file)
	} else {
		cmd = exec.Command("code", file)
	}

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
