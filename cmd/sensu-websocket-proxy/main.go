package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/jspaleta/sensu-websocket-proxy/websocketproxy"
)

var (
	flagBackend = flag.String("backend", "", "Backend URL for proxying")
	flagPort    = flag.String("port", "9081", "Proxy port for local proxy")
)

func main() {
	flag.Parse()
	u, err := url.Parse(*flagBackend)
	if err != nil {

		log.Fatalln(err)
	}
	fmt.Printf("backend: %v\n", *flagBackend)
	fmt.Printf("Url: %v\n", u)
	if len(u.Host) == 0 {
		log.Fatalln("no host defined")
	}

	log.Println("Listening for requests at ws://localhost:" + *flagPort + " for backend URL: " + *flagBackend)

	err = http.ListenAndServe(":"+string(*flagPort), websocketproxy.NewProxy(u))
	if err != nil {
		log.Fatalln(err)
	}
}
