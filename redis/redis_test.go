package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestIncr(t *testing.T) {
	Init("")
	ctx := context.Background()
	r, e := Incr(ctx, "tmp")
	fmt.Print(r, e)
}

func TestDecr(t *testing.T) {
	Init("")
	ctx := context.Background()
	r, e := Decr(ctx, "tmp")
	fmt.Print(r, e)
}

func TestGet(t *testing.T) {
	Init("")
	ctx := context.Background()
	r, e := Get(ctx, "tmp")
	fmt.Printf("%s %v", r, e)
}
