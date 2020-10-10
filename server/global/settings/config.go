package settings

import (
	"log"

	"github.com/go-ini/ini"
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Server struct {
	Address string
	DataDir string
}

var ServerSetting = &Server{}

func Setup() {
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("load ini 'config/app.ini': %v", err)
	}

	if err = cfg.Section("database").MapTo(DatabaseSetting); err != nil {
		log.Fatalf("mapping to database: %v", err)
	}
	if err = cfg.Section("server").MapTo(ServerSetting); err != nil {
		log.Fatalf("mapping to server: %v", err)
	}
}
