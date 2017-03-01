package main

import "github.com/sturfeeinc/cglog/context"

const PARAMETERS  = 6

func main() {
	println("get context")
	ctx := library.GetContext()

	go func() {
		println("send deffered data")
		library.SentDeffer(ctx, 98)
	}()

	println("start calculation")

	data := []float32{
		37.7922668457, 122.393577576, 3.3707399368, 15.1165, 94.8269, 0.0,
		37.7945899963, 122.393371582, 2.9586400986, 143.546, 103.36, 0.0,
	}

	library.StartCalc(ctx, data, len(data) / PARAMETERS, 70, 1920, 1080)
}
