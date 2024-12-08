package telegram

import (
	"errors"

	"github.com/telegram_bot/bot/events"
	e "github.com/telegram_bot/bot/lib/error_wrapping"
	"github.com/telegram_bot/bot/storage"
	telegram "github.com/telegram_bot/bot/tg_client"
)

const (
	defaultState = iota
	stateAdd
	stateDel
)

type EventProcessor struct {
	tg     *telegram.Client
	offset int
	db     storage.Storage
	state  map[string]int
}

// constructor for event processor
func New(client *telegram.Client, db storage.Storage) *EventProcessor {
	state := make(map[string]int, 0)
	return &EventProcessor{
		tg:    client,
		db:    db,
		state: state,
	}
}

func (p *EventProcessor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap(errGettingEvent, err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)

	res := events.Event{
		Type: fetchType(upd),
		Text: fetchText(upd),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			Username: upd.Message.From.Username,
		}
	}

	return res
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	} else {
		return events.Message
	}
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	} else {
		return upd.Message.Text
	}
}

func (p *EventProcessor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	default:
		return e.Wrap(errUknownEventType, errors.New(errProcessingMessage))
	}
}

func (p *EventProcessor) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap(errProcessingMessage, err)
	}

	if err := p.doCMD(event.Text, meta.ChatID, meta.Username); err != nil {
		return e.Wrap(errProcessingMessage, err)
	}

	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap(errGettingMeta, errors.New(errUknownMetaType))
	} else {
		return res, nil
	}
}
