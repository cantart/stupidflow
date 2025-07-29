package main

import (
	"context"
	"os"
	"runtime"
	"runtime/pprof"
)

func runDAGWithProfiles(ctx context.Context, cfg Config, cpuFile, memFile string) error {
	var cpuF *os.File
	var err error
	if cpuFile != "" {
		cpuF, err = os.Create(cpuFile)
		if err != nil {
			return err
		}
		if err = pprof.StartCPUProfile(cpuF); err != nil {
			cpuF.Close()
			return err
		}
		defer func() {
			pprof.StopCPUProfile()
			cpuF.Close()
		}()
	}

	if err = runDAG(ctx, cfg); err != nil {
		return err
	}

	if memFile != "" {
		memF, err := os.Create(memFile)
		if err != nil {
			return err
		}
		defer memF.Close()
		runtime.GC()
		if err = pprof.WriteHeapProfile(memF); err != nil {
			return err
		}
	}
	return nil
}
