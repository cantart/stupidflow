package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	cfgPath := flag.String("config", "dag.yaml", "path to DAG config")
	enablePprof := flag.Bool("pprof", false, "enable pprof server")
	dockerImage := flag.String("docker-image", "", "run tasks in docker using this image")
	profileFile := flag.String("profile", "cpu.pprof", "write CPU profile to file")
	memProfileFile := flag.String("mem-profile", "mem.pprof", "write memory profile to file")
	flag.Parse()

	if *enablePprof {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}

	cfg, err := loadConfig(*cfgPath)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	ctx := context.Background()
	if *dockerImage != "" {
		if err := runDAGInDocker(ctx, *cfgPath, *dockerImage, *profileFile, *memProfileFile); err != nil {
			log.Fatalf("run DAG: %v", err)
		}
	} else {
		if err := runDAGWithProfiles(ctx, cfg, *profileFile, *memProfileFile); err != nil {
			log.Fatalf("run DAG: %v", err)
		}
	}

	os.Exit(0)
}
