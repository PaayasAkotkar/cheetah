// Package example showcases useage of the package
package example

import (
	cheetah "app/cheetah/cheetha"
	"context"
	"log"
	"time"
)

func PlayCheetah() {
	log.Println("[basic cheetah example]")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	type Greet struct {
		greet string
	}
	che := cheetah.New[string, Greet](100)
	key := "greet"
	parcel := Greet{"hey mate!!!"}
	che.Publish(key, &parcel)

	time.Sleep(1 * time.Second)

	ch := che.Subscribe(key)
	defer che.Unsubscribe(key, ch)
	select {
	case token := <-ch:
		log.Printf("[recieved package: %s]", token.greet)
	case <-ctx.Done():
		return
	}
}
