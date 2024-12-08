package telegram

type Meta struct {
	ChatID   int
	Username string
}

const (
	errUknownEventType   = "unknown event type"
	errUknownMetaType    = "unknown meta type"
	errGettingEvent      = "cannot get event"
	errProcessingMessage = "cannot process message"
	errGettingMeta       = "cannot get meta"
	errGettingRates      = "cannot get rates"
	errAddingTokens      = "cannot add token(s) to user's list"
	errDeletingTokens    = "cannot delete token(s) from user's list"
)
