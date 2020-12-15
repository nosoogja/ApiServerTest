package main

import (
	"fmt"
	"net/http"
)

var p = fmt.Println
var checkNum = 0

func main() {
	p(":::::::::::::::::::::::::::::::::::::::::::::::::")

	http.HandleFunc("/catchme", PageCatchme)

	err := http.ListenAndServe("localhost:5000", nil)
	if err != nil {
		p(err)
	}

}

func PageCatchme(res http.ResponseWriter, req *http.Request) {

	checkNum += 1
	fmt.Fprintf(res, "%d", checkNum)

}
