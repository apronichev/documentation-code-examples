package main

import "runtime"

// to check function calls in debugger, debug the main function (breakpoints are set programmatically)
func main() {
	t := T{"ok"}
	_ = FuncOneResult()
	_ = t.MethodOneResult()
	_, _, _ = t.MethodManyUnnamedResults()
	_, _, _ = t.MethodManyNamedResults()
	t.MethodNoResults()
	Panic(false)
	_ = Factorial(10, false)
	runtime.Breakpoint()
	// evaluate the following expressions in debugger:
	// FuncOneResult(), t.MethodOneResult() - the result of function or method is shown
	// t.MethodManyUnnamedResults() - ~r0, ~r1, ... are used as return value names
	// t.MethodManyNamedResults() - return value names specified in function signature are used
	// t.MethodNoResults() - result is undefined since function doesn't return anything
	// Panic(false) - the argument to the panic() call is shown as a result
	// Factorial(10, true) - notice that evaluation doesn't stop at breakpoints inside the evaluated function
	// UnusedFunc(), t.UnusedMethod() - evaluation fails, the error explains that unused function/method are not included into executable
	println("ok")
}

type T struct {
	f string
}

func (t T) MethodNoResults() {
	println("foo")
}

func (t T) MethodOneResult() string {
	return "ok" //position MethodOneResult
}

func (t T) MethodManyUnnamedResults() (string, int, bool) {
	return "ok", 42, false
}

func (t T) MethodManyNamedResults() (one string, two int, three bool) {
	return "ok", 42, false
}

func (t T) UnusedMethod() int {
	return 0
}

func UnusedFunc() int {
	return 0
}

func FuncOneResult() string {
	return "ok"
}

func Factorial(n int, withBreak bool) int {
	if withBreak {
		runtime.Breakpoint()
	}
	if n == 0 {
		return 1
	} else {
		return n * Factorial(n - 1, withBreak)
	}
}

func Panic(b bool) {
	if b {
		panic("panic value")
	}
}
