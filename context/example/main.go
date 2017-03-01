package main

import "github.com/sturfeeinc/cglog/context"

func main() {
	println("get context")
	ctx := library.GetContext()

	go func() {
		println("send deffered data")
		library.SentDeffer(ctx, 98)
	}()

	println("start calculation")
	library.StartCalc(ctx)
}
