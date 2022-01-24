package main

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func setEnvironmentVariables() {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
    	log.Panic("unable to get the current filename (e.g. __filename)")
	}

	dir := filepath.Dir(file)

	environmentPath := filepath.Join(dir, ".env")
	envErr := godotenv.Load(environmentPath)
	if envErr != nil {
		log.Fatal(envErr.Error())
	}
}