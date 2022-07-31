package conf

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	Config  CONFIG
	Usermap map[string]string
)

type CONFIG struct {
	Name     string   `yaml:"Name"`
	Addr     string   `yaml:"Addr"`
	WeCom    WECOM    `yaml:"WeCom"`
	UserData []string `yaml:"UserData"`
	Gitlab   GITLAB   `yaml:"Gitlab"`
}

type GITLAB struct {
	AccessToken string `yaml:"AccessToken"`
}

type WECOM struct {
	DefaultAddr string `yaml:"DefaultAddr"`
}

// LoadConfig loads specified file into CONFIG struct
func LoadConfig(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	err = yaml.Unmarshal(content, &Config)
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	if Config.Gitlab.AccessToken == "" {
		log.Fatalf("access_token is empty")
	}

	if Config.WeCom.DefaultAddr == "" {
		log.Fatalf("default robot addr is empty")
	}
}

// LoadUserData will load users' data at userid.csv into user map
func LoadUserData() {
	Usermap = make(map[string]string)
	//1.1 check
	for _, path := range Config.UserData {
		_, err := os.Stat(path)
		if err != nil || os.IsNotExist(err) {
			log.Fatalf("config file has not found at valid address")
		}

		err = parseUserData(path, &Usermap)
		if err != nil {
			log.Fatalf("parse user data path err:%s", err)
		}
	}
}

//1.2 parse userid.csv to userMap
func parseUserData(path string, userMap *map[string]string) error {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return err
	}

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	// Ignore first row, which is the header
	for _, row := range data[1:] {
		(*userMap)[row[0]] = row[1]
	}

	return nil
}
