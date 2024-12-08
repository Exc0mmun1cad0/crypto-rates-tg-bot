package telegram

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
	editMessageMthod = "editMessage"
)

const (
	errSendingMessage = "cannot send message"
	errDoingRequest   = "cannot do request"
	errGettingUpdates = "cannot get updates"
)

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

//* Only gor logging, for sending messages field 'Chat' is used
type From struct {
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}