package controllers

import (
	"net/http"
	"io/ioutil"
	"os/exec"
	"path"
	"fmt"
	"strconv"
	"strings"
	"github.com/go-zoo/bone"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func EntryShow(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(bone.GetValue(r, "id"))
	entry := models.Entry{}

	if !models.ExistsById(&entry, id) || err != nil {
		utils.NotFound(w, r)
		return
	}

	competition := new(models.Competition)
	models.DB.Model(&entry).Related(&competition)

	utils.Render(w, r, "entry.html", &utils.Props{
		"entry": entry,
		"competition": competition,
	})
}

func EntryNew(w http.ResponseWriter, r *http.Request) {
	comp, _ := (&models.Competition{}).Current()
	utils.Render(w, r, "enter.html", &utils.Props{
		"competition": comp,
	})
}

func EntryCreate(w http.ResponseWriter, r *http.Request) {
	comp, _ := (&models.Competition{}).Current()
	current_user := (*r.Context().Value("data").(*utils.Props))["current_user"]
	entry := &models.Entry{
		UserID: current_user.(models.User).ID,
		CompetitionID: comp.ID,
	}
	if models.Exists(&entry) {
		utils.BadRequest(w, r)
		return
	}

	// read the file
	reader, _ := r.MultipartReader()
	part, _ := reader.NextPart()
	code, _ := ioutil.ReadAll(part)
	codeString := string(code)

	// pass it to the runner
	runner := exec.Command(path.Join(utils.BasePath, "./runners/go.sh"), codeString)
	runnerOut, _ := runner.StdoutPipe()
	runnerErr, _ := runner.StderrPipe()
	runner.Start()
	output, _ := ioutil.ReadAll(runnerOut)
	errors, _ := ioutil.ReadAll(runnerErr)
	outputString := string(output)
	errorsString := string(errors)
	runner.Wait()

	errorArray := strings.Split(errorsString, "\n")

	var timeResult float64;

	if len(errorArray) != 2 {
		utils.Render(w, r, "enter.html", &utils.Props{
			"competition": comp,
			"stderrError": true,
		})
		return
	} else {
		// last element is an empty string
		// second to last is time in format 0.00
		// the rest are real errors
		// need a safer way to do this
		timeResult, _ = strconv.ParseFloat(errorArray[len(errorArray)-2], 64)

	}

	result, err := strconv.ParseFloat(strings.Trim(outputString, "\n\r"), 64)
	if result == comp.Solution && err == nil {
		user := (*r.Context().Value("data").(*utils.Props))["current_user"]
		entry := models.Entry{
			CompetitionID: comp.ID,
			UserID: user.(models.User).ID,
			Language: "go",
			Code: codeString,
			ExecTime: timeResult,
		}

		models.Create(&entry)

		http.Redirect(w, r, fmt.Sprintf("/entry/%v", entry.ID), 307)
	} else {
		utils.Render(w, r, "enter.html", &utils.Props{
			"competition": comp,
			"stdoutError": true,
		})
	}
}
