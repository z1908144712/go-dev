package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 12345
	}
	fmt.Println(ret)
	rets, ok := ctx.Value("session").(string)
	if !ok {
		rets = "gogogo"
	}
	fmt.Println(rets)
}
func main() {
	ctx := context.WithValue(context.Background(), "trace_id", 123456)
	ctx = context.WithValue(ctx, "session", "abcd")
	process(ctx)
}
