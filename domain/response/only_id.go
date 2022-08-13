package response

type OnlyID struct {
	ID int `json:"id"`
}

type OnlyIDs struct {
	ID []int `json:"id"`
}
