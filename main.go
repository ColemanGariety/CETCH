package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"path"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

var Env = os.Getenv("env")

func main() {
	config := make(map[string]map[string]string)
	data, _ := ioutil.ReadFile(path.Join(utils.BasePath, "db/dbconf.yml"))
	_ = yaml.Unmarshal([]byte(data), &config)

	utils.InitTemplates()
	models.InitDB(config[Env]["open"])

	log.Println("Whispering...")
	log.Fatal(http.ListenAndServe(":" + config[Env]["port"], NewRouter()))
}
