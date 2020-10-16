package structs

type (
	AuthResponse struct {
		UserInfo
		PastDisplayNames []struct {
			DisplayName string `json:"displayName"`
			UpdatedAt   string `json:"updated_at"`
		} `json:"pastDisplayNames"`
		HasEmail               bool          `json:"hasEmail"`
		HasPendingEmail        bool          `json:"hasPendingEmail"`
		Email                  string        `json:"email"`
		ObfuscatedEmail        string        `json:"obfuscatedEmail"`
		ObfuscatedPendingEmail string        `json:"obfuscatedPendingEmail"`
		EmailVerified          bool          `json:"emailVerified"`
		HasBirthday            bool          `json:"hasBirthday"`
		Unsubscribe            bool          `json:"unsubscribe"`
		Friends                []string      `json:"friends"`
		FriendGroupNames       []interface{} `json:"friendGroupNames"`
		CurrentAvatar          string        `json:"currentAvatar"`
		CurrentAvatarAssetURL  string        `json:"currentAvatarAssetUrl"`
		AccountDeletionDate    interface{}   `json:"accountDeletionDate"`
		AcceptedTOSVersion     int           `json:"acceptedTOSVersion"`
		SteamID                string        `json:"steamId"`
		SteamDetails           struct {
			Steamid                  string `json:"steamid"`
			Communityvisibilitystate int    `json:"communityvisibilitystate"`
			Profilestate             int    `json:"profilestate"`
			Personaname              string `json:"personaname"`
			Profileurl               string `json:"profileurl"`
			Avatar                   string `json:"avatar"`
			Avatarmedium             string `json:"avatarmedium"`
			Avatarfull               string `json:"avatarfull"`
			Personastate             int    `json:"personastate"`
			Realname                 string `json:"realname"`
			Primaryclanid            string `json:"primaryclanid"`
			Timecreated              int    `json:"timecreated"`
			Personastateflags        int    `json:"personastateflags"`
			Avatarhash               string `json:"avatarhash"`
			Gameextrainfo            string `json:"gameextrainfo"`
			Gameid                   string `json:"gameid"`
		} `json:"steamDetails"`
		OculusID              string `json:"oculusId"`
		HasLoggedInFromClient bool   `json:"hasLoggedInFromClient"`
		HomeLocation          string `json:"homeLocation"`
		TwoFactorAuthEnabled  bool   `json:"twoFactorAuthEnabled"`
		Feature               struct {
			TwoFactorAuth bool `json:"twoFactorAuth"`
		} `json:"feature"`
		OnlineFriends  []string `json:"onlineFriends"`
		ActiveFriends  []string `json:"activeFriends"`
		OfflineFriends []string `json:"offlineFriends"`

		RequiresTwoFactorAuth []string `json:"requiresTwoFactorAuth"`
	}
)
