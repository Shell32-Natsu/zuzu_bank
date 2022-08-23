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

func (c *Config) IsAdmin(u int64) bool {
	us := strconv.FormatInt(u, 10)
	for _, admin := range c.AdminUsers {
		if admin == us {
			return true
		}
	}
	return false
}

func (c *Config) IsAllowedUser(u int64) bool {
	us := strconv.FormatInt(u, 10)
	for _, au := range c.AllowedUsers {
		if au == us {
			return true
		}
	}
	return false
}

func (c *Config) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("BotKey: %s\n", c.BotKey))
	builder.WriteString(fmt.Sprintf("AdminUsers: %s\n", c.AdminUsers))
	builder.WriteString(fmt.Sprintf("AllowedUsers: %s\n", c.AllowedUsers))
	return builder.String()
}

func NewConfig(p string) (*Config, error) {
	b, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %s", p, err)
	}

	log.Printf("config file:\n%s", string(b))

	var c Config
	json.Unmarshal(b, &c)

	log.Printf("parsed config:\n%s", c.String())
	return &c, nil
}
