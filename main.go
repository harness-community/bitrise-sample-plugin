package main

import (
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

	// Export DeployDir as an environment variable
	err := os.Setenv("EXPORTED_DEPLOY_DIR", config.DeployDir)
	if err != nil {
		failf("Failed to set EXPORTED_BITRISE_DEPLOY_DIR: %v\n", err)
	}

	logger.Infof("Exported DeployDir as EXPORTED_BITRISE_DEPLOY_DIR=%s", config.DeployDir)

	logger.Donef("Done")
}
