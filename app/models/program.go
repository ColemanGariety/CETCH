package models

import (
	"os/exec"
	"path"
	"io/ioutil"
	"strings"
	"strconv"
	"errors"

	"github.com/JacksonGariety/cetch/app/utils"
)

func ProgramResultAndExecTime(codeString string, language string) (*float64, *float64, error) {
	var runnerPath string
	switch language {
	case "go":
		runnerPath = "./runners/go.sh"
	case "haskell":
		runnerPath = "./runners/haskell.sh"
	}

	runner := exec.Command(path.Join(utils.BasePath, runnerPath), codeString)
	runnerOut, _ := runner.StdoutPipe()
	runnerErr, _ := runner.StderrPipe()
	runner.Start()
	outputRead, _ := ioutil.ReadAll(runnerOut)
	errorsRead, _ := ioutil.ReadAll(runnerErr)
	outputString := string(outputRead)
	errorsString := string(errorsRead)
	runner.Wait()

	outputArray := strings.Split(outputString, "\n")
	errorsArray := strings.Split(errorsString, "\n")

	var result *float64
	var execTime *float64
	var error error

	if len(errorsArray) != 2 {
		result = nil
		execTime = nil
		error = errors.New(errorsString)
	} else {
		// last element is an empty string
		// second to last is time in format 0.00
		// the rest are real errors
		// need a safer way to do this
		r, err := strconv.ParseFloat(strings.Trim(outputArray[len(outputArray)-2], "\n\r"), 64)
		et, _ := strconv.ParseFloat(errorsArray[len(errorsArray) - 2], 64)
		result = &r
		execTime = &et
		if err != nil {
			error = err
		}
	}


	return result, execTime, error
}
