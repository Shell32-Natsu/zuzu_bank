package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	BotKey       string
	AdminUsers   []string
	AllowedUsers []string
	Debug        bool
}

func (c *Config) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("BotKey: %s\n", c.BotKey))
	builder.WriteString(fmt.Sprintf("AdminUsers: %s\n", c.AdminUsers))
	builder.WriteString(fmt.Sprintf("AllowedUsers: %s\n", c.AllowedUsers))
	return builder.String()
}

var config Config

func IsAdmin(u int64) bool {
	us := strconv.FormatInt(u, 10)
	for _, admin := range config.AdminUsers {
		if admin == us {
			return true
		}
	}
	return false
}

func IsAllowedUser(u int64) bool {
	us := strconv.FormatInt(u, 10)
	for _, au := range config.AllowedUsers {
		if au == us {
			return true
		}
	}
	return false
}

func SetDebug(d bool) {
	config.Debug = d
}

func IsDebug() bool {
	return config.Debug
}

func BotKey() string {
	return config.BotKey
}

func InitConfig(p string) error {
	b, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("failed to open config file %s: %s", p, err)
	}

	log.Printf("config file:\n%s", string(b))

	json.Unmarshal(b, &config)

	log.Printf("parsed config:\n%s", config.String())
	return nil
}
