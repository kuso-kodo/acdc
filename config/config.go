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
	Title        string
	Postgres     postgresConfig `toml:"postgres"`
	RootUser     rootUserConfig `toml:"root"`
	JWTConfig    jwtConfig      `toml:"jwt"`
	TicketConfig ticketConfig   `toml:"ticket"`
	HotelConfig  hotelConfig    `toml:"hotel"`
	AirConfig    AirConfig      `toml:"air"`
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
}

type jwtConfig struct {
	Title string
	Key   string
}

type ticketConfig struct {
	PageSize int `toml:"page_size"`
}

type hotelConfig struct {
	MaxRoom int `toml:"max_room"`
}

type AirConfig struct {
	MaxServeSize          int     `toml:"max_serve_size" json:"max_serve_size"`
	LowFanSpeedFeeRate    float32 `toml:"low_fan_speed_fee_rate" json:"low_fan_speed_fee_rate"`
	MediumFanSpeedFeeRate float32 `toml:"medium_fan_speed_fee_rate" json:"medium_fan_speed_fee_rate"`
	HighFanSpeedFeeRate   float32 `toml:"high_fan_speed_fee_rate" json:"high_fan_speed_fee_rate"`
	LowPriorityFactor     float32 `toml:"low_priority_factor" json:"low_priority_factor"`
	MediumPriorityFactor  float32 `toml:"medium_priority_factor" json:"medium_priority_factor"`
	HighPriorityFactor    float32 `toml:"high_priority_factor" json:"high_priority_factor"`
	Period                int     `toml:"period" json:"period"`
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
