package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	ctx = context.WithValue(ctx, "UserId", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")
	result, err := dbAccess(ctx)
	if err != nil {
		http.Error(res, err.Error(), http.StatusRequestTimeout)
	}

	fmt.Fprintln(res, result)

}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	ch := make(chan int)

	go func() {
		uid := ctx.Value("UserId").(int)
		time.Sleep(10 * time.Second)

		if ctx.Err() != nil {
			return
		}
		ch <- uid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}

}

func bar(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(res, ctx)
}
