package types

type MessageView struct {
	Sender   User      `json:"sender"`
	Messages []Message `json:"messages"`
	Receiver User      `json:"receiver"`
}
