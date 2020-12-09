package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func main() {
	stop := make(chan struct{}) // stop chan

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

	// notify system sigals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	go func() {
		sig := <-sigs
		logrus.Infof("recv sigal: %v, Now exiting...", sig)
		cancel()

	}()

	go func() {
		<-ctx.Done()
		s1.Shutdown(ctx) // ignore err
		s2.Shutdown(ctx) // ignore err

		close(stop)

	}()

	<-stop
}
