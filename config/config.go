package config

import (
	"fmt"
	"gopkg.in/go-ini/ini.v1"
	"log"
	"os"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config ConfigList

const targetEnvName = "GO_ENV"

//main関数より先に実行する
func init() {
	LoadConfig()
	//utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	var cfg *ini.File
	var err error
	if "" == os.Getenv(targetEnvName) {
		cfg, err = ini.Load("config.ini")
	}
	//cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("wev").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
}
