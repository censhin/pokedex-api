package utils

import (
    "encoding/json"
    "io/ioutil"
    "log"
)

type Config struct {
    Db struct {
	    Host string
	    Port string
    }
}

var config *Config = nil

func GetConfig() *Config {

    if config == nil {
        config = new(Config)

        file, err := ioutil.ReadFile("../config.json")
        if err != nil {
            log.Fatal(err)
        }

        err = json.Unmarshal(file, &config)
        if err != nil {
            log.Fatal(err)
        }
    }

    return config
}
