package persist

import (
	"testing"

	"github.com/nekoimi/fetcher/engine"
)

func TestItemSave(t *testing.T) {
	goods := engine.Item{
		Title:   "goods title",
		Price:   "233",
		Details: []string{""},
	}
	saveItem(goods)
}
