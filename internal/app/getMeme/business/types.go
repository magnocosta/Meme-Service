package business

import "meme_service/internal/shared/types"

type Input interface {
  GetCoordinate() types.Coordinate
  GetQuery() string
}

type Item interface {
  GetURL() string
}

type Output interface {
  GetItems() []Item
}

type item struct {
  url string
}

func (i item) GetURL() string { return i.url }

type output struct {
  items []Item
}

func (o output) GetItems() []Item { return o.items }
