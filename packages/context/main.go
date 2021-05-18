package main

import (
	"context"
	"fmt"
)

// Package context defines the Context type, which carries deadlines,
// cancellation signals, and other request-scoped values across API
// boundaries and between processes.
// Incoming requests to a server should create a Context, and outgoing calls
// to servers should accept a Context.
func main() {
	// Basic Type
	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

}
