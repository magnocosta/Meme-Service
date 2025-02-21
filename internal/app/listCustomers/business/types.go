package business

type Item interface {
  GetID() string
  GetName() string
  GetEmail() string
}

type Output interface {
  GetItems() []Item
}

type output struct {
  items []Item
}

func (o output) GetItems() []Item { return o.items }
