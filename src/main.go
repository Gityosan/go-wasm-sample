package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

// var timer *time.Ticker
// var times int64
// var stopTime int64
var inputtingNum = ""
var result = 0
var operator = None
var afterEqual = false

const (
	Plus = iota
	Sub
	Mul
	Div
	None
)

func input(this js.Value, i []js.Value) interface{} {
	if afterEqual {
		operator = None
		result = 0
	}
	afterEqual = false
	// fmt.Println("operator", operator, operator == Plus, operator == 0)
	inputtingNum += strconv.Itoa(i[0].Int())
	fmt.Println("result", result, "inputtingNum", inputtingNum, "operator", operator)
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", inputtingNum)
	return "success"
}
func clean(this js.Value, i []js.Value) interface{} {
	fmt.Println("result", result, "inputtingNum", inputtingNum, "operator", operator)
	inputtingNum = ""
	result = 0
	operator = None
	afterEqual = false
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", "0")
	return "success"
}
func operation(this js.Value, i []js.Value) interface{} {
	fmt.Println("result", result, "inputtingNum", inputtingNum, "operator", operator)
	if operator != None {
		equal(this, i)
	} else {
		int1, _ := strconv.Atoi(inputtingNum)
		result += int1
	}
	inputtingNum = ""
	switch i[0].String() {
	case "plus":
		operator = Plus
	case "minus":
		operator = Sub
	case "multiply":
		operator = Mul
	case "division":
		operator = Div
	default:
		operator = None
	}
	afterEqual = false
	return "success"
}
func equal(this js.Value, i []js.Value) interface{} {
	fmt.Println("result", result, "inputtingNum", inputtingNum, "operator", operator)
	int1, _ := strconv.Atoi(inputtingNum)
	switch operator {
	case Plus:
		result = result + int1
	case Sub:
		result = result - int1
	case Mul:
		result = result * int1
	case Div:
		if int1 != 0 {
			result = result / int1
		}
	}
	inputtingNum = ""
	operator = None
	afterEqual = true
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", result)
	return "success"
}

// func start(this js.Value, args []js.Value) interface{} {
// 	startTime := time.Now().Unix()
// 	timer := time.NewTicker(time.Millisecond * 100)
// 	fmt.Println("Hello, WebAssembly!", startTime)
// 	for {
// 		select {
// 		case <-timer.C:
// 			now := time.Now().Unix()
// 			times := now - startTime + stopTime
// 			js.Global().Get("document").Call("getElementByClass", "display").Set("textContent", times)
// 		}
// 	}
// 	return "success"
// }
// func stop(this js.Value, args []js.Value) interface{} {
// 	fmt.Println("Hello, WebAssembly!", this, args)
// 	defer timer.Stop()
// 	stopTime := times
// 	js.Global().Get("document").Call("getElementById", "startButton").Set("textContent", "リスタート")
// 	return "stopTime"
// }
// func reset(this js.Value, args []js.Value) interface{} {
// 	fmt.Println("Hello, WebAssembly!", this, args)
// 	defer timer.Stop()
// 	js.Global().Get("document").Call("getElementsByClassName", "display").Set("innerHTML", "0.0")
// 	js.Global().Get("document").Call("getElementById", "startButton").Set("innerHTML", "リスタート")
// 	start()
// 	return "success"
// }
// func prints(this js.Value, i []js.Value) interface{} {
// 	fmt.Println(i)
// 	return "success"
// }
// func test(this js.Value, vs []js.Value) interface{} {
// 	timer := time.NewTicker(time.Millisecond * 100)
// 	fmt.Println("Hello, WebAssembly!", this, vs)
// 	fmt.Println(timer)
// 	fmt.Println(reflect.TypeOf(timer))
// 	fmt.Println(reflect.TypeOf(this))
// 	fmt.Println(reflect.TypeOf(timer))
// 	for {
// 		select {
// 		case <-timer.C:
// 			timer.Stop()
// 		}
// 	}

// 	<-timer.C
// 	fmt.Println(timer)
// 	fmt.Println(js.Global().Get("document").Call("getElementById", "display").Get("innerHTML"))
// 	js.Global().Get("document").Call("getElementById", "display").Set("innerHTML", "10.0")
// 	return "success"
// }

func setFunctions() {
	// js.Global().Set("prints", js.FuncOf(prints))
	// js.Global().Set("start", js.FuncOf(start))
	// js.Global().Set("stop", js.FuncOf(stop))
	// js.Global().Set("reset", js.FuncOf(reset))
	// js.Global().Set("test", js.FuncOf(test))
	js.Global().Set("operation", js.FuncOf(operation))
	js.Global().Set("equal", js.FuncOf(equal))
	js.Global().Set("input", js.FuncOf(input))
	js.Global().Set("clean", js.FuncOf(clean))
	// cc := time.NewTicker(5 * time.Second)
	// for range cc.C {
	// 	fmt.Println("tock")
	// }
	// cc.Stop()
}

func main() {
	fmt.Println("Hello, WebAssembly!")
	c := make(chan struct{})
	setFunctions()
	<-c
}
