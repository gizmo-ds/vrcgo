package structs

type (
	WebSocketMessage struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	}

	WebSocketFriendAdd struct {
		UserID string `json:"userId"`
		User   User   `json:"user"`
	}

	WebSocketFriendDelete struct {
		UserID string `json:"userId"`
	}

	WebSocketFriendActive struct {
		UserID string `json:"userId"`
		User   User   `json:"user"`
	}

	WebSocketUserUpdate struct {
		UserID string `json:"userId"`
		User   User   `json:"user"`
	}

	WebSocketFriendLocation struct {
		UserID           string `json:"userId"`
		User             User   `json:"user"`
		World            World  `json:"world"`
		Location         string `json:"location"`
		Instance         string `json:"instance"`
		CanRequestInvite bool   `json:"canRequestInvite"`
	}

	WebSocketFriendOffline struct {
		UserID string `json:"userId"`
	}

	WebSocketFriendOnline struct {
		UserID string `json:"userId"`
		User   User   `json:"user"`
	}

	WebSocketFriendUpdate struct {
		UserID string `json:"userId"`
		User   User   `json:"user"`
	}
)
