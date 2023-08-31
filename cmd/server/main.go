package main

import (
	"github.com/SenRecep/redisclone/src/apiserver"
	"log"
	"os"
)

func main() {
	if err := apiserver.New(
		apiserver.WithServerEnv(os.Getenv("SERVER_ENV")),
		apiserver.WithLogLevel(os.Getenv("LOG_LEVEL")),
	); err != nil {
		log.Fatal(err)
	}
}
