package main_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

var p = fmt.Println
var uri = "http://localhost:5000/catchme"
var wg = sync.WaitGroup{}

func TestDo(t *testing.T) {
	p(":::::::::::::::::::::::::::::::::::::::::::::::::")
	defer p("END")
	t1 := time.Now()

	cnt := 0
	for i := 0; i < 6000; i++ {
		num := cnt
		wg.Add(1)
		go proc(num)
		cnt += 1
	}

	wg.Wait()
	p(time.Since(t1))
}

func proc(num int) {
	var err error
	defer wg.Done()

	resp, err := http.Get(uri)
	if err != nil {
		log.Panic(num, err)
	}

	res_byte, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Panic(num, err)
	}
	_ = res_byte
	//p(string(res_byte))
}

func _TestDo(t *testing.T) {
	p(":::::::::::::::::::::::::::::::::::::::::::::::::")
	var err error

	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}

	res_byte, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}

	p(string(res_byte))

}
