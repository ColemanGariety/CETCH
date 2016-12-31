package controllers

import (
	"net/http"
	"io/ioutil"
	"os/exec"
	"path"
	"fmt"
	"strconv"
	"strings"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func EntryShow(w http.ResponseWriter, r *http.Request) {
	entry, _ := (&models.Entry{}).Find()
	utils.Render(w, r, "entry.html", &utils.Props{
		"entry": entry,
	})
}

func EntryNew(w http.ResponseWriter, r *http.Request) {
	comp, _ := (&models.Competition{}).Current()
	utils.Render(w, r, "enter.html", &utils.Props{
		"competition": comp,
	})
}

func EntryCreate(w http.ResponseWriter, r *http.Request) {
	// read the file
	reader, _ := r.MultipartReader()
	part, _ := reader.NextPart()
	code, _ := ioutil.ReadAll(part)
	codeString := string(code)

	// pass it to the runner
	runner := exec.Command(path.Join(utils.BasePath, "./runners/go.sh"), codeString)
	runnerOut, _ := runner.CombinedOutput()

	comp, _ := (&models.Competition{}).Current()
	result, err := strconv.ParseFloat(strings.Trim(string(runnerOut), "\n\r"), 64)
	if result == comp.Solution && err == nil {
		user := (*r.Context().Value("data").(*utils.Props))["current_user"]
		entry := (&models.Entry{
			CompetitionID: comp.ID,
			UserID: user.(*models.User).ID,
			Language: "go",
			Code: codeString,
		})

		entry.Create()

		http.Redirect(w, r, fmt.Sprintf("/entry/%v", entry.ID), 307)
	} else {
		fmt.Fprintf(w, "%s", string(runnerOut))
	}

}
