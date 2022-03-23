package main

import (
	"fmt"
	"reflect"
	"syscall/js"
	"time"
)

var timer *time.Ticker
var times int64
var stopTime int64

func start() {
	startTime := time.Now().Unix()
	timer := time.NewTicker(time.Millisecond * 100)
	for {
		select {
		case <-timer.C:
			now := time.Now().Unix()
			times := now - startTime + times
			js.Global().Get("document").Call("getElementByClass", "display").Set("textContent", times)
		}
	}
}
func stop() {
	defer timer.Stop()
	js.Global().Get("document").Call("getElementById", "startButton").Set("textContent", "リスタート")
}
func reset() {
	defer timer.Stop()
	js.Global().Get("document").Call("getElementByClass", "display").Set("textContent", "0.0")
	js.Global().Get("document").Call("getElementById", "startButton").Set("textContent", "スタート")
	start()

}
func print(i []js.Value) {
	fmt.Println(i)
}
func test(this js.Value, vs []js.Value) interface{} {
	timer := time.NewTicker(time.Millisecond * 100)

	fmt.Println("Hello, WebAssembly!")
	fmt.Println(vs[0])
	fmt.Println(reflect.TypeOf(vs))
	fmt.Println(reflect.TypeOf(this))
	// fmt.Println(reflect.TypeOf(timer))
	timer.Stop()
	<-timer.C
	// fmt.Println(timer)
	// fmt.Println(reflect.TypeOf(timer))
	return "success"
}
func setFunctions() {
	// js.Global().Set("print", js.FuncOf(print))
	// js.Global().Set("start", js.FuncOf(start))
	// js.Global().Set("stop", js.FuncOf(stop))
	// js.Global().Set("reset", js.FuncOf(reset))
	js.Global().Set("test", js.FuncOf(test))
}

func main() {
	fmt.Println("Hello, WebAssembly!")
	c := make(chan struct{}, 0)
	setFunctions()
	<-c
}
