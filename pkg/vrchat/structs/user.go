package structs

type (
	User struct {
		LimitedUser
		BioLinks           []string `json:"bioLinks"`
		Status             string   `json:"status"`
		StatusDescription  string   `json:"statusDescription"`
		State              string   `json:"state"`
		LastLogin          string   `json:"last_login"`
		AllowAvatarCopying bool     `json:"allowAvatarCopying"`
		FriendKey          string   `json:"friendKey"`
		Location           string   `json:"location"`
		WorldID            string   `json:"worldId"`
		InstanceID         string   `json:"instanceId"`
	}

	LimitedUser struct {
		ID                             string   `json:"id"`
		Username                       string   `json:"username"`
		DisplayName                    string   `json:"displayName"`
		Bio                            string   `json:"bio"`
		CurrentAvatarImageURL          string   `json:"currentAvatarImageUrl"`
		CurrentAvatarThumbnailImageURL string   `json:"currentAvatarThumbnailImageUrl"`
		LastPlatform                   string   `json:"last_platform"`
		Tags                           []string `json:"tags"`
		DeveloperType                  string   `json:"developerType"`
		IsFriend                       bool     `json:"isFriend"`
	}

	CurrentUser struct {
		User
		PastDisplayNames       []PastDisplayName `json:"pastDisplayNames"`
		HasEmail               bool              `json:"hasEmail"`
		HasPendingEmail        bool              `json:"hasPendingEmail"`
		Email                  string            `json:"email"`
		ObfuscatedEmail        string            `json:"obfuscatedEmail"`
		ObfuscatedPendingEmail string            `json:"obfuscatedPendingEmail"`
		EmailVerified          bool              `json:"emailVerified"`
		HasBirthday            bool              `json:"hasBirthday"`
		Unsubscribe            bool              `json:"unsubscribe"`
		Friends                []string          `json:"friends"`
		FriendGroupNames       []interface{}     `json:"friendGroupNames"`
		CurrentAvatar          string            `json:"currentAvatar"`
		CurrentAvatarAssetURL  string            `json:"currentAvatarAssetUrl"`
		AccountDeletionDate    interface{}       `json:"accountDeletionDate"`
		AcceptedTOSVersion     int               `json:"acceptedTOSVersion"`
		SteamID                string            `json:"steamId"`
		SteamDetails           SteamDetails      `json:"steamDetails"`
		OculusID               string            `json:"oculusId"`
		HasLoggedInFromClient  bool              `json:"hasLoggedInFromClient"`
		HomeLocation           string            `json:"homeLocation"`
		TwoFactorAuthEnabled   bool              `json:"twoFactorAuthEnabled"`
		Feature                Feature           `json:"feature"`
		OnlineFriends          []string          `json:"onlineFriends"`
		ActiveFriends          []string          `json:"activeFriends"`
		OfflineFriends         []string          `json:"offlineFriends"`

		RequiresTwoFactorAuth []string `json:"requiresTwoFactorAuth"`
	}

	Feature struct {
		TwoFactorAuth bool `json:"twoFactorAuth"`
	}

	SteamDetails struct {
		SteamID                  string `json:"steamid"`
		CommunityVisibilityState int    `json:"communityvisibilitystate"`
		ProfileState             int    `json:"profilestate"`
		PersonaName              string `json:"personaname"`
		ProfileURL               string `json:"profileurl"`
		Avatar                   string `json:"avatar"`
		AvatarMedium             string `json:"avatarmedium"`
		AvatarFull               string `json:"avatarfull"`
		PersonaState             int    `json:"personastate"`
		RealName                 string `json:"realname"`
		PrimaryClanID            string `json:"primaryclanid"`
		TimeCreated              int    `json:"timecreated"`
		PersonaStateFlags        int    `json:"personastateflags"`
		AvatarHash               string `json:"avatarhash"`
		GameExtraInfo            string `json:"gameextrainfo"`
		GameID                   string `json:"gameid"`
	}

	PastDisplayName struct {
		DisplayName string `json:"displayName"`
		UpdatedAt   string `json:"updated_at"`
	}
)
