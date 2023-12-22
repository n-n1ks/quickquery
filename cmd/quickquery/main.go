package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"quickquery/internal/config"
	"quickquery/internal/database"
	"quickquery/internal/database/compute"
	"quickquery/internal/database/storage"
	"quickquery/internal/database/storage/engine/memengine"
	"quickquery/internal/initialization"
)

func main() {
	config.LoadEnvs()

	logger := initialization.NewLogger(config.GetConfigLevel())

	comp, err := compute.NewCompute(compute.NewParser(), compute.NewAnalyzer(), logger)
	if err != nil {
		log.Fatalln(err)
	}

	engine := memengine.NewEngine()
	store, err := storage.NewStorage(engine, logger)
	if err != nil {
		log.Fatalln(err)
	}

	database, err := database.NewDatabase(comp, store, logger)
	if err != nil {
		log.Fatalln(err)
	}

	logger.Info("database started")
	for {
		fmt.Print("> ")

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		if text == "exit\n" {
			os.Exit(0)
		}

		result := database.HandleQuery(context.TODO(), text)
		fmt.Printf("-> %s\n", result)
	}
}

func MustEnv(key string) string {
	if os.Getenv(key) == "" {
		log.Panicf("unknown %s param in ENV", key)
	}
	return os.Getenv(key)
}

func Env(key, defaultValue string) string {
	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
