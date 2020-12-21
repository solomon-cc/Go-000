package main

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	g, ctx := errgroup.WithContext(context.Background())

	s := NewHttpServer(":80")

	g.Go(func() error {
		g.Go(func() error {
			<-ctx.Done()
			logrus.Info("http ctx done")
			return s.Shutdown(context.TODO())
		})

		return s.ListenAndServe()
	})

	// notify system sigals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	g.Go(func() error {
		for {
			logrus.Info("signal sentinels")
			select {
			case <-ctx.Done():
				logrus.Info("signal ctx done")
				return ctx.Err()
			case <-sigs:
				logrus.Infof("recv sigal, Now exiting...")
				return errors.New("signal exit")

			}
		}
	})

	err := g.Wait()
	logrus.Error(err)

}
