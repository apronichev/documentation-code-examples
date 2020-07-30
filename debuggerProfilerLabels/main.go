package main

/*
Put a breakpoint at 'println("ok")' and debug the 'main' function. Once breakpoint is reached, check out a list of goroutines
in the debugger. Goroutine names include the following information: /api/profile, userId: <some number>. You can use this
information to find a particular goroutine during debugging and core dump analysis.

To learn more about pprof labels, see https://rakyll.org/profiler-labels/.
*/
import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		f(ctx)
	}
	time.Sleep(time.Second)
}

func f(ctx context.Context) {
	time.Sleep(time.Millisecond)
	println("ok")
}