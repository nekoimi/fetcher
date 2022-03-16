package persist

import (
	"fmt"
	"log"
)

func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Count #%d : %v", itemCount, item)
			saveItem(item)
			itemCount++
		}
	}()
	return out
}

func saveItem(item interface{}) {
	fmt.Printf("SaveItem: %v \n", item)
}
