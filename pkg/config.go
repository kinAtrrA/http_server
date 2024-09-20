package src

import (
	"bufio"
	"io"
	"log"
	"os"
	"sync"

	"github.com/pelletier/go-toml/v2"
)

type ServerConfig struct {
	Port int
	Host string
}

var lock = &sync.Mutex{}

var singleton *ServerConfig

func GetInstance() *ServerConfig {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		file, err := os.Open("config.toml")
		if err != nil {
			log.Fatalf("config.toml was not found in the folder.")
		}
		defer file.Close()
		stat, err := file.Stat()
		if err != nil {
			log.Fatalf("Cannot get file size of config.toml")
		}
		byteSlice := make([]byte, stat.Size())
		_, err = bufio.NewReader(file).Read(byteSlice)
		if err != nil && err != io.EOF {
			log.Fatalf("Cannot read config.toml into bytearray.")
		}
		err = toml.Unmarshal([]byte(byteSlice), &singleton)
		if err != nil {
			panic(err)
		}
	}
	return singleton
}
