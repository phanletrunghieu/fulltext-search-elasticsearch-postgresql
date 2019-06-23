package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/config/database/elastic"
	"github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/config/database/pg"
	httpServ "github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/http"
)

func main() {
	closeDB := pg.Init()
	defer closeDB()

	elastic.Init()

	r := httpServ.NewHTTPHandler()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errs <- r.Run(":3000")
	}()

	log.Println("exit", <-errs)
}
