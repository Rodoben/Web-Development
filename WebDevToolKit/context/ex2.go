package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = context.WithValue(ctx, "UserId", 777)
	ctx = context.WithValue(ctx, "fname", "ronald")
	results := dbAccess(ctx)
	fmt.Fprintln(res, results)
}

func dbAccess(ctx context.Context) int {
	uid := ctx.Value("UserId").(int)
	return uid
}

func bar(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(res, ctx)
}
