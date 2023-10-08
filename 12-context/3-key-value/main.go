package main

import (
	"context"
	"fmt"
)

// CONTEXT WITHVALUE

// the context can pass data side to side

// im will can pass metadada in side to side via context instead of parameters

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "token", "$23HsqWsfh#845")
	bookHotel(ctx, "Hotel Piazzetta")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")

	fmt.Println(token)
}
