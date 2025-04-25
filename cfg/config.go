package cfg

import (
	"encoding/json"
	"os"
)

type Config struct {
	Ethereum struct {
		RPCURL string `json:"rpc_url"`
	} `json:"ethereum"`
}

func LoadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err!= nil {
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}