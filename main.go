package main

import (
	"fmt"
	"os"

	"github.com/bitrise-io/go-steputils/stepconf"
	"github.com/bitrise-io/go-utils/log"
)

// Config ...
type Config struct {
	ProjectLocation   string `env:"project_location,dir"`
	ReportPathPattern string `env:"report_path_pattern"`
	Variant           string `env:"variant"`
	Module            string `env:"module"`
	Arguments         string `env:"arguments"`
	CacheLevel        string `env:"cache_level,opt[none,only_deps,all]"`
	DeployDir         string `env:"BITRISE_DEPLOY_DIR,dir"`
}

var logger = log.NewLogger(false)

func failf(f string, args ...interface{}) {
	logger.Errorf(f, args...)
	os.Exit(1)
}

func main() {
	var config Config

	if err := stepconf.Parse(&config); err != nil {
		failf("Process config: couldn't create step config: %v\n", err)
	}

	stepconf.Print(config)

	// Get the DRONE_OUTPUT environment variable
	droneOutputPath := os.Getenv("DRONE_OUTPUT")
	if droneOutputPath == "" {
		failf("Environment variable DRONE_OUTPUT is not set.\n")
	}

	// Prepare the key-value pair content
	fileContent := fmt.Sprintf("EXPORTED_BITRISE_DEPLOY_DIR=%s\n", config.DeployDir)

	// Write to the file specified by DRONE_OUTPUT
	err := os.WriteFile(droneOutputPath, []byte(fileContent), 0644)
	if err != nil {
		failf("Failed to write to %s: %v\n", droneOutputPath, err)
	}

	logger.Infof("Written to file %s: %s", droneOutputPath, fileContent)
}
