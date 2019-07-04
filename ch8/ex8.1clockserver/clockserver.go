package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func handleConn(c net.Conn, tz string) {
	defer c.Close()
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatal(err)
	}
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format(tz+"--> 15:04:05\n"))
		if err != nil {
			return //client disconnected
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {
	port := flag.Int("port", 8000, "Port to run the clock server")
	timezone := flag.String("TZ", "US/Eastern", "Timezone for the clock server")

	flag.Parse()

	log.Printf("======== Starting clock server at port %d for timezone %s ========\n", *port, *timezone)

	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, *timezone)
	}
}
