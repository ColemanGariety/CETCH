package models

import (
	"os/exec"
	"path"
	"io/ioutil"
	"strings"
	"strconv"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/JacksonGariety/cetch/app/utils"
)

type Entry struct {
	gorm.Model
	Language        string
	Code            string
	ExecTime        float64
	UserID          uint
	CompetitionID   uint
	User            User `gorm:"ForeignKey:UserID"`
	Competition     Competition `form:"ForeignKey:CompetitionID"`
}

type Entries []Entry

func (entries *Entries) FindByUserId(id uint) *Entries {
	DB.Where("user_id = ?", id).Find(&entries)
	return entries
}

func (entry *Entry) TimesFaster() float64 {
	comp := new(Competition)
	DB.Model(&entry).Related(comp)
	return comp.AverageExecTime() / entry.ExecTime
}

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

func (entries Entries) Len() int {
	return len(entries)
}

func (entries Entries) Less(i, j int) bool {
	e := entries
	return utils.TimesFaster(e[i].ExecTime, e[i].Competition.AverageExecTime()) < utils.TimesFaster(e[j].ExecTime, e[i].Competition.AverageExecTime())
}

func (entries Entries) Swap(i, j int) {
	e := entries
	e[i], e[j] = e[j], e[i]
}
