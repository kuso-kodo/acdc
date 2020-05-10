package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

const configFilePath = "config.toml"

type serverConfig struct {
	Title     string
	Postgres  postgresConfig `toml:"postgres"`
	RootUser  rootUserConfig `toml:"root-user"`
	JWTConfig jwtConfig      `toml:"jwt-middleware"`
}

type postgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type rootUserConfig struct {
	UserName string
	Password string
	Phone    string
}

type jwtConfig struct {
	Title string
	Key   string
}

func (cfg *postgresConfig) GetConnectionString() string {
	sb := strings.Builder{}
	sb.WriteString("host=" + cfg.Host + " ")
	sb.WriteString("port=" + strconv.Itoa(cfg.Port) + " ")
	sb.WriteString("user=" + cfg.User + " ")
	sb.WriteString("password=" + cfg.Password + " ")
	sb.WriteString("dbname=" + cfg.DbName)
	return sb.String()
}

var (
	config *serverConfig
	once   sync.Once
)

func GetConfig() *serverConfig {
	once.Do(func() {
		filePath, err := filepath.Abs(configFilePath)
		if err != nil {
			log.Panic(err)
		}
		log.Println("Config file found, path: " + filePath)
		if _, err := toml.DecodeFile(filePath, &config); err != nil {
			log.Panic(err)
		}
	})
	return config
}
