package structs

type (
	Favorite struct {
		ID         string   `json:"id"`
		Type       string   `json:"type"`
		FavoriteId string   `json:"favoriteId"`
		Tags       []string `json:"tags"`
	}
)
