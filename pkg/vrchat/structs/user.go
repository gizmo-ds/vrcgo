package structs

type (
	UserInfo struct {
		ID                             string   `json:"id"`
		Username                       string   `json:"username"`
		DisplayName                    string   `json:"displayName"`
		Bio                            string   `json:"bio"`
		BioLinks                       []string `json:"bioLinks"`
		CurrentAvatarImageURL          string   `json:"currentAvatarImageUrl"`
		CurrentAvatarThumbnailImageURL string   `json:"currentAvatarThumbnailImageUrl"`
		Status                         string   `json:"status"`
		StatusDescription              string   `json:"statusDescription"`
		State                          string   `json:"state"`
		Tags                           []string `json:"tags"`
		DeveloperType                  string   `json:"developerType"`
		LastLogin                      string   `json:"last_login"`
		LastPlatform                   string   `json:"last_platform"`
		AllowAvatarCopying             bool     `json:"allowAvatarCopying"`
		IsFriend                       bool     `json:"isFriend"`
		FriendKey                      string   `json:"friendKey"`
		Location                       string   `json:"location"`
		WorldID                        string   `json:"worldId"`
		InstanceID                     string   `json:"instanceId"`
	}
)
