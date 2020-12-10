package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Word"))
}

func startHttpServer(ctx context.Context, addr string) error {
	s := http.Server{
		Addr:    addr,
		Handler: &Handler{},
	}
	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Printf("%s Shutdown!\n", addr)
		s.Shutdown(ctx)
	}(ctx)
	return s.ListenAndServe()
}

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return startHttpServer(ctx, ":9990")
	})
	group.Go(func() error {
		return startHttpServer(ctx, ":9991")
	})

	group.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-quit:
			return fmt.Errorf("通过信号关闭http服务")
		case <-ctx.Done():
			return fmt.Errorf("http服务关闭")
		}

	})

	if err := group.Wait(); err != nil {
		fmt.Printf("err:%v \n", err)
		return
	}

}
