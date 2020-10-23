package structs

import (
	"github.com/mitchellh/mapstructure"
)

type (
	Notification struct {
		ID             string      `json:"id"`
		Type           string      `json:"type"`
		SenderUserID   string      `json:"senderUserId"`
		SenderUsername string      `json:"senderUsername"`
		ReceiverUserID string      `json:"receiverUserId"`
		Message        string      `json:"message,omitempty"`
		Details        interface{} `json:"details,omitempty"`
		CreatedAt      string      `json:"created_at"`
		JobName        string      `json:"jobName,omitempty"`
		JobColor       string      `json:"jobColor,omitempty"`
	}

	NotificationInviteDetails struct {
		WorldID   string `json:"worldId"`
		WorldName string `json:"worldName"`
	}

	NotificationRequestInviteDetails struct {
		Platform string `json:"platform"`
	}

	NotificationVoteToTick struct {
		UserToKickID    string `json:"userToKickId"`
		InitiatorUserID string `json:"initiatorUserId"`
	}
)

func (v Notification) Invite() (result NotificationInviteDetails) {
	_ = mapstructure.Decode(v.Details, &result)
	return result
}

func (v Notification) RequestInvite() (result NotificationRequestInviteDetails) {
	_ = mapstructure.Decode(v.Details, &result)
	return result
}

func (v Notification) VoteToTick() (result NotificationVoteToTick) {
	_ = mapstructure.Decode(v.Details, &result)
	return result
}
