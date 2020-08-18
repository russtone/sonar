package http_test

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/bi-zone/sonar/pkg/server/http"
)

var (
	srv    *http.Server
	srvTLS *http.Server
)

func TestMain(m *testing.M) {
	if err := setupGlobals(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	ret := m.Run()

	os.Exit(ret)
}

func setupGlobals() error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	srv = http.New("localhost:1080",
		http.NotifyStartedFunc(func() {
			wg.Done()
		}),
	)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(fmt.Errorf("fail to start server: %w", err))
		}
	}()

	cert, err := tls.LoadX509KeyPair("test/cert.pem", "test/key.pem")
	if err != nil {
		log.Fatal(err)
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	srvTLS = http.New("localhost:1443",
		http.TLSConfig(tlsConfig),
		http.NotifyStartedFunc(func() {
			wg.Done()
		}),
	)

	go func() {
		if err := srvTLS.ListenAndServe(); err != nil {
			log.Fatal(fmt.Errorf("fail to start server: %w", err))
		}
	}()

	if waitTimeout(&wg, 30*time.Second) {
		return errors.New("timeout waiting for server to start")
	}

	return nil
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
