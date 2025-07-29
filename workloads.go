package main

import (
	"context"
	"crypto/rand"
	"os"
	"time"
)

func runSleep(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func runCPUBurn(ctx context.Context, d time.Duration) error {
	end := time.Now().Add(d)
	for time.Now().Before(end) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}
	return nil
}

func runMemorySpike(ctx context.Context, sizeMB int, d time.Duration) error {
	b := make([]byte, sizeMB*1024*1024)
	for i := range b {
		b[i] = byte(i)
	}
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func runFakeIO(ctx context.Context, sizeMB int) error {
	f, err := os.CreateTemp("", "stupidflow")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	data := make([]byte, 1024*1024)
	for i := 0; i < sizeMB; i++ {
		if _, err := rand.Read(data); err != nil {
			return err
		}
		if _, err := f.Write(data); err != nil {
			return err
		}
	}
	return f.Close()
}
