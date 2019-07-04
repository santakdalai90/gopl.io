package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	newyorkURL := flag.String("NewYork", "localhost:8000", "Port for NewYork time")
	londonURL := flag.String("London", "localhost:8001", "Port for London time")
	tokyoURL := flag.String("Tokyo", "localhost:8002", "Port for Tokyo time")

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		newyorkConn, err := net.Dial("tcp", *newyorkURL)
		if err != nil {
			log.Fatal(err)
		}
		defer newyorkConn.Close()
		mustCopy(os.Stdout, newyorkConn)
	}()

	go func() {
		defer wg.Done()
		londonConn, err := net.Dial("tcp", *londonURL)
		if err != nil {
			log.Fatal(err)
		}
		defer londonConn.Close()
		mustCopy(os.Stdout, londonConn)
	}()

	go func() {
		defer wg.Done()
		tokyoConn, err := net.Dial("tcp", *tokyoURL)
		if err != nil {
			log.Fatal(err)
		}
		defer tokyoConn.Close()
		mustCopy(os.Stdout, tokyoConn)
	}()

	wg.Wait()
}
