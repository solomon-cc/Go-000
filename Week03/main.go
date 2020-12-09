package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, _ := errgroup.WithContext(ctx)

	s1 := NewHttpServer(":8080")
	s2 := NewHttpServer(":8081")

	g.Go(func() error {
		if err := s1.ListenAndServe(); err != nil {
			cancel()
			return err
		}

		return nil

	})

	g.Go(func() error {
		if err := s2.ListenAndServe(); err != nil {
			cancel()
			return err
		}

		return nil

	})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	go func() {
		sig := <-sigs
		logrus.Errorf("catch sigal: %v, exiting..", sig)
		cancel()
		
	}()
}
