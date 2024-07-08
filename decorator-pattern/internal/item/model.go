package item

import "fmt"

type Item struct {
	ID int
}

func (item *Item) JSON() string {
	return fmt.Sprintf(`{"id": "%d"}`, item.ID)
}
