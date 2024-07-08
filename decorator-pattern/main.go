package main

import (
	"fmt"
	"os"

	"github.com/eltonjr/go-experiments/decorator-pattern/cmd/server"
	"github.com/eltonjr/go-experiments/decorator-pattern/internal/item"
	"github.com/eltonjr/go-experiments/decorator-pattern/pkg/log"
)

func main() {
	addr1 := "localhost:8080"
	go serveItemModule(addr1)
	fmt.Println("Quiet server running at", addr1)

	addr2 := "localhost:8081"
	go serveItemModuleLogger(addr2)
	fmt.Println("Logged server running at", addr2)

	// wait forever
	ch := make(chan struct{})
	<-ch

}

// serveItemModule builds the module with regular
// implementations and injecting it to nested components
func serveItemModule(addr string) {
	r, err := item.NewRepository()
	if err != nil {
		panic(err)
	}

	m, err := item.NewItemModule(r)
	if err != nil {
		panic(err)
	}

	err = server.NewServer(addr, m)
	if err != nil {
		panic(err)
	}
}

// serveItemModuleLogger builds the regular components
// and adapt them with a logger, keeping the same interface
func serveItemModuleLogger(addr string) {
	l := log.NewLogger(log.LevelDebug, os.Stdout)

	r, err := item.NewRepository()
	if err != nil {
		panic(err)
	}

	rl := item.NewRepoLogger(l, r)

	// use the logger repository as a normal repository
	m, err := item.NewItemModule(rl)
	if err != nil {
		panic(err)
	}

	ml := item.NewItemLogger(l, m)

	// use the logger module as a normal module
	err = server.NewServer(addr, ml)
	if err != nil {
		panic(err)
	}
}
