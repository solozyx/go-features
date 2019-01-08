package main

import (
	"context"
	"fmt"
)

// ret:11111
// session:sdlkfjkaslfsalfsafjalskfj

func process(ctx context.Context) {
	ret,ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 22222
	}
	fmt.Printf("ret:%d\n", ret)

	s , _ := ctx.Value("session").(string)
	fmt.Printf("session:%s\n", s)
}

func main() {
	// context.Background() 会默认生成1个ctx
	ctx := context.WithValue(context.Background(), "trace_id", 11111)
	// ctx树状结构 可继承
	ctx = context.WithValue(ctx, "session", "sdlkfjkaslfsalfsafjalskfj")
	process(ctx)
}