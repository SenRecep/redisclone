package main

import (
	"log"
	"os"

	"github.com/SenRecep/redisclone/src/apiserver"
)

func main() {
	if err := apiserver.New(
		apiserver.WithServerEnv(os.Getenv("SERVER_ENV")),
		apiserver.WithLogLevel(os.Getenv("LOG_LEVEL")),
	); err != nil {
		log.Fatal(err)
	}
}
