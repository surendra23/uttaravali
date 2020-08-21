package main

import (
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/unrolled/render"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const local string = "LOCAL"

func main() {
	var (
		// environment variables
		env      = os.Getenv("ENV")      // LOCAL, DEV, STG, PRD
		port     = os.Getenv("PORT")     // server traffic on this port
		version  = os.Getenv("VERSION")  // path to VERSION file
		fixtures = os.Getenv("FIXTURES") // path to fixtures file
	)
	if env == "" || env == local {
		// running from localhost, so set some default values
		env = local
		port = "3001"
		dir, _ := os.Getwd()
		version = dir + "/VERSION"
		fixtures = dir + "/fixtures.json"
	}
	// reading version from file
	version, err := ParseVersionFile(version)
	if err != nil {
		log.Fatal(err)
	}
	// load fixtures data into mock database
	db, err := LoadFixturesIntoMockDatabase(fixtures)
	if err != nil {
		log.Fatal(err)
	}
	// initialse application context
	ctx := AppContext{
		Render:  render.New(),
		Version: version,
		Env:     env,
		Port:    port,
		DB:      db,
	}
	// start application
	StartServer(ctx)
}
