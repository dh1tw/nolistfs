package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"

	nfs "github.com/dh1tw/nolistfs"
)

var port = flag.Int("p", 7474, "webserver port")
var ip = flag.String("i", "127.0.0.1", "webserver ip address")

//go:embed html
var htmlDir embed.FS

func main() {

	flag.Parse()

	// make the "html" folder our root directory
	htmlAssets, err := fs.Sub(htmlDir, "html")
	if err != nil {
		log.Panic(err)
	}

	// create the http.FileSystem for our assets directory
	fsWithListing := http.FS(htmlAssets)

	// wrap the http.FileSystem with our custom "no listing filesystem"
	fsNoListing := nfs.New(fsWithListing)

	// instanciate the file server
	fileSvr := http.FileServer(fsNoListing)

	// capture systems signal CTRL-C so that we can gracefully exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// exit when CTRL-C has been pressed
	go func() {
		<-c
		os.Exit(0)
	}()

	// start http server
	fmt.Printf("Listening on %s:%d\n", *ip, *port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", *ip, *port), fileSvr)
	if err != nil {
		log.Println(err)
	}

}
