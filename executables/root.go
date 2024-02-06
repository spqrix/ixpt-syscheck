package executables

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func CheckForExecutables(executables []string, verbose bool) ([]string, error) {
	var missing []string
	operating_system := runtime.GOOS
	if operating_system != "linux" {
		err := errors.New("This check is for Linux only (for now).")
		return missing, err
	}
	for _, c := range executables {
		script_command := fmt.Sprintf("type -ap %s", c)
		command := exec.Command("sh", "-c", script_command)
		var out strings.Builder
		command.Stdout = &out
		err := command.Run()
		if err != nil {
			missing = append(missing, c)
		} else {
			if verbose {
				fmt.Printf("Path for %s:\n\t%s", c,out.String())
			}
		}
	}
	return missing, nil
}
