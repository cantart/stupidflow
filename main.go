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
	dockerImage := flag.String("docker-image", "", "docker image used to run tasks")
	profileFile := flag.String("profile", "cpu.pprof", "write CPU profile to file")
	memProfileFile := flag.String("mem-profile", "mem.pprof", "write memory profile to file")
	flag.Parse()

	if *enablePprof {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}

	ctx := context.Background()
	if err := execute(ctx, *cfgPath, *dockerImage, *profileFile, *memProfileFile); err != nil {
		log.Fatalf("run DAG: %v", err)
	}

	os.Exit(0)
}

var (
	runDAGWithProfilesFunc = runDAGWithProfiles
	runDAGInDockerFunc     = runDAGInDocker
)

func execute(ctx context.Context, cfgPath, dockerImage, cpuFile, memFile string) error {
	cfg, err := loadConfig(cfgPath)
	if err != nil {
		return err
	}
	if dockerImage != "" {
		return runDAGInDockerFunc(ctx, cfgPath, dockerImage, cpuFile, memFile)
	}
	return runDAGWithProfilesFunc(ctx, cfg, cpuFile, memFile)
}
