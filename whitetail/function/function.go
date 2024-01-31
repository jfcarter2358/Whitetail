package function

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"whitetail/logger"

	"github.com/google/uuid"
)

type Function struct {
	Language  string   `yaml:"language"`
	Run       string   `yaml:"run"`
	Arguments []string `yaml:"arguments"`
}

func Run(f Function, args map[string]interface{}) error {
	logger.Debugf("", "Callback params:")
	logger.Tracef("", "  language : %s", f.Language)
	logger.Tracef("", "  run      : %s", f.Run)
	logger.Tracef("", "  arguments: %v", f.Arguments)
	logger.Debugf("", "Function args:")
	logger.Debugf("", "  args     : %v", args)
	id := uuid.New().String()
	logger.Debugf("", "Starting callback run with ID %s", id)

	runDir := fmt.Sprintf("/tmp/whitetail/%s", id)
	err := os.MkdirAll(runDir, 0777)
	if err != nil {
		logger.Errorf("", "Error creating run directory %s", err.Error())
		return err
	}

	script_path := fmt.Sprintf("%s/run.%s", runDir, f.Language)

	script_command := []string{}
	script_contents := ""

	switch f.Language {
	case "sh":
		script_command = []string{"/bin/bash", fmt.Sprintf("%s/run.sh", runDir)}
		script_contents += "script_directory=\"$(dirname \"$(readlink -fm \"$0\")\")\"\n"
	case "py":
		script_command = []string{"python3", fmt.Sprintf("%s/run.py", runDir)}
		script_contents += "import sys\n"
		script_contents += "import os\n"
		script_contents += "script_directory = os.path.dirname(os.path.abspath(sys.argv[0]))\n"
	default:
		logger.Errorf("", "Invalid language type: %s", f.Language)
		// return err
	}

	script_contents += f.Run

	// Write out our run script
	script_data := []byte(script_contents)
	err = os.WriteFile(script_path, script_data, 0777)
	if err != nil {
		logger.Errorf("", "Error writing run file %s", err.Error())
		return err
	}

	logger.Debugf("", "Running function at /tmp/whitetail/%s", id)
	if err := DoFunction(script_command, f.Arguments, args, id); err != nil {
		logger.Errorf("", "error on function execution: %s", err.Error())
		return err
	}
	return nil
}

func DoFunction(script_command, arguments []string, args map[string]interface{}, id string) error {
	for _, key := range arguments {
		val := args[key]
		script_command = append(script_command, fmt.Sprintf("%v", val))
	}
	logger.Tracef("", "Executing function command %v", script_command)
	output, err := exec.Command(script_command[0], script_command[1:]...).CombinedOutput()
	outputString := strings.TrimSuffix(string(output), "\n")
	logger.Tracef("", "Probe output: %s", outputString)
	if err != nil {
		logger.Errorf("", "Error running script for probe /tmp/whitetail/%s/: %s", id, err.Error())
	}
	return err
}
