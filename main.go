package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	cfgPath := flag.String("config", "dag.yaml", "path to DAG config")
	enablePprof := flag.Bool("pprof", false, "enable pprof server")
	dockerImage := flag.String("docker-image", "", "run tasks in docker using this image")
	report := flag.Bool("report", false, "show CPU profile report")
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
		out, err := runDAGInDocker(ctx, *cfgPath, *dockerImage, *report)
		if err != nil {
			log.Fatalf("run DAG: %v", err)
		}
		if *report {
			fmt.Print(out)
		}
	} else if *report {
		out, err := runDAGProfile(ctx, cfg)
		if err != nil {
			log.Fatalf("run DAG: %v", err)
		}
		fmt.Print(out)
	} else {
		if err := runDAG(ctx, cfg); err != nil {
			log.Fatalf("run DAG: %v", err)
		}
	}

	os.Exit(0)
}
