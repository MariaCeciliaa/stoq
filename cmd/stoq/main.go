package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/MariaCeciliaa/stoq/internal/config"
)

func main() {

	default_config := &config.Config{}

	if file_config := os.Getenv("STOQ_FILE_CONFIG"); file_config != "" {

		file, err := os.ReadFile(file_config)
		if err != nil {
			log.Panicln(err.Error())
		}

		err = json.Unmarshal(file, default_config)
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	conf := config.NewConfig(default_config)

	data, _ := json.Marshal(conf)

	fmt.Println(string(data))
}