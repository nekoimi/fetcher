package engine

import (
	"fmt"

	"github.com/nekoimi/fetcher/fetcher"
)

func worker(r Request) (ParseResult, error) {
	body, fetchErr := fetcher.Fetch(r.Url)
	if fetchErr != nil {
		fmt.Printf("Fetcher Error.  Url => [ %s ], Error => %v", r.Url, fetchErr)
		return ParseResult{}, fetchErr
	}
	return r.ParseFunc(body), nil
}
