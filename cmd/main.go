package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/blck-snwmn/gorabin"
)

func main() {
	var n int64
	flag.Int64Var(&n, "n", 2, "is prime")
	flag.Parse()

	ok, err := gorabin.IsPrime(n)
	if err != nil {
		log.Fatal(err)
	}
	msg := "is prime number"
	if !ok {
		msg = "is not prime number"
	}
	fmt.Printf("%d %s\n", n, msg)
}
