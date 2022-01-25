package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func count(start int, end int) chan int {
	ch := make(chan int)
	closeChan := make(chan bool)

	go func(ch chan int, clCh chan bool) {
		for i := start; i <= end/2; i++ {
			ch <- i + rand.Int()
		}

		clCh <- true
	}(ch, closeChan)

	go func(ch chan int, clCh chan bool) {
		for i := end/2 + 1; i <= end; i++ {
			ch <- i + rand.Int()
		}

		clCh <- true
	}(ch, closeChan)

	go func(ch chan int, clCh chan bool) {
		<-clCh
		<-clCh

		close(clCh)
		close(ch)
	}(ch, closeChan)

	return ch
}

func main() {
	r := 99
	out := bufio.NewWriter(os.Stdout)

	err := writeLine("calculating 100 pseudorandom numbers", out)
	if err != nil {
		log.Println(err)
	}

	for i := range count(1, r) {
		msg := fmt.Sprintf("number %d was generated", i)

		err = writeLine(msg, out)
		if err != nil {
			log.Println(err)
		}
	}

	err = writeLine("calculation finished", out)
	if err != nil {
		log.Println(err)
	}
}

func writeLine(s string, out *bufio.Writer) error {
	s += "\n"

	_, err := out.WriteString(s)
	if err != nil {
		return fmt.Errorf("writing: %w", err)
	}

	err = out.Flush()
	if err != nil {
		return fmt.Errorf("flushing: %w", err)
	}

	return nil
}
