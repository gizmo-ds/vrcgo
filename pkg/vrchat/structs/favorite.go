package structs

type (
	FavoriteInfo struct {
		ID         string   `json:"id"`
		Type       string   `json:"type"`
		FavoriteId string   `json:"favoriteId"`
		Tags       []string `json:"tags"`
	}
)
