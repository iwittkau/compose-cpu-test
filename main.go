package main

import (
	"flag"
	"log"
)

func main() {
	var cpus int
	flag.IntVar(&cpus, "cpus", 1, "set the number of cpus to occupy")
	flag.Parse()

	stop := make(chan struct{})
	for i := 0; i < cpus; i++ {
		go func(i int) {
			log.Printf("started worker %d/%d", i+1, cpus)
			for {
				select {
				case <-stop:
					return
				default:
				}
			}
		}(i)
	}

	<-stop

}
