package api

// Item provides implementation of the shop item
type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}

// String retruns name of the Item
func (itm *Item) String() string {
	return itm.Name
}
