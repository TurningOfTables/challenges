package main

import (
	"fmt"
	"math/rand"
	"time"
)

var conveyor = make(chan int)
var workercapacity = make(chan struct{}, 3)
var qacapacity = make(chan struct{}, 3)
var done = make(chan int)

func main() {
	start := time.Now()
	reqProducts := 15
	manager(reqProducts)
	dur := time.Since(start).Seconds()
	fmt.Printf("Made %v products in %v seconds", reqProducts, dur)
}

func manager(products int) {
	for x := 1; x <= products; x++ {
		workercapacity <- struct{}{}
		go worker(x)
		qacapacity <- struct{}{}
		go qa(x)
	}

	for x := 1; x <= products; x++ {
		product := <-done
		fmt.Printf("Finished product # %v - [%v]\n", x, product)
	}
}

func worker(x int) {
	fmt.Printf("Started worker %v\n", x)
	time.Sleep(time.Second * 2)
	product := rand.Intn(100)
	fmt.Printf("Worker %v made [%v]\n", x, product)
	<-workercapacity
	conveyor <- product
}

func qa(x int) {
	fmt.Printf("Started QA %v\n", x)
	product := <-conveyor
	time.Sleep(time.Second * 2)
	fmt.Printf("QA %v completed testing [%v]\n", x, product)
	<-qacapacity
	done <- product
}
