package main

import (
	"context"
	"log"

	"github.com/Krab1o/meebin/internal/app"
)

func main() {

	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	a.Run()
}
