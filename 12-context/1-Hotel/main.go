package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)

	// its good practice at the end of a cancel.
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// select is similar to the CASE but works in a way asynchronously, he stay waiting the result when the result arrives he performs the action.
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
