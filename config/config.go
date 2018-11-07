package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// ConfigDB dados para conexão com banco de dados
type ConfigDB struct {
	Server   string
	Database string
}

// Read read the file with datas for connection
func (c *ConfigDB) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal("[ERROR CONNECTION]", err)
	}
}
