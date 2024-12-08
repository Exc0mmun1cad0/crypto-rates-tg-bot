package telegram

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	coincapapi "github.com/telegram_bot/bot/api/coincap_api"
	e "github.com/telegram_bot/bot/lib/error_wrapping"
)

const (
	StartCmd  = "/start"
	HelpCmd   = "/help"
	RateCmd   = "/rates"
	DelCmd    = "/delete"
	DelAllCmd = "/deleteall"
	AddCmd    = "/add"
	ListCmd   = "/list"
)

func (p *EventProcessor) doCMD(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if p.state[username] != defaultState {
		switch p.state[username] {
		case stateAdd:
			return p.Add(chatID, username, text)
		case stateDel:
			return p.Del(chatID, username, text)
		}
	}

	switch text {
	case StartCmd:
		return p.sendHello(chatID)
	case HelpCmd:
		return p.sendHelp(chatID)
	case ListCmd:
		return p.tg.SendMessage(chatID, msgListInfo+msgList)
	case RateCmd:
		return p.sendRates(chatID, username)
	case DelCmd:
		return p.sendDelete(chatID, username)
	case DelAllCmd:
		return p.sendDeleteAll(chatID, username)
	case AddCmd:
		return p.sendAdd(chatID, username)
	default:
		return p.tg.SendMessage(chatID, msgUknownCommand)
	}
}

func (p *EventProcessor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func (p *EventProcessor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *EventProcessor) sendRates(chatID int, username string) error {
	tokens, err := p.db.GetAll(username)
	if err != nil {
		return e.Wrap(errGettingRates, err)
	}

	if len(tokens) == 0 {
		return p.tg.SendMessage(chatID, msgNoTokens)
	}

	response, err := coincapapi.GetLatestRates(tokens)
	if err != nil {
		return e.Wrap(errGettingRates, err)
	}

	// timeAndDate :=
	respFullTime := time.UnixMilli(response.Timestamp)
	y, m, d := respFullTime.Date()
	respTime := fmt.Sprintf("%d-%d-%d %d:%d", y, m, d, respFullTime.Hour(), respFullTime.Minute())
	latestRates := fmt.Sprintf(msgRates, respTime)

	for _, token := range response.Data {
		tokenPrice, err := strconv.ParseFloat(token.PriceUsd, 32)
		if err != nil {
			return e.Wrap(errGettingRates, err)
		}

		latestRates += fmt.Sprintf("- %s: %.2f\n", token.ID, tokenPrice)
	}

	return p.tg.SendMessage(chatID, latestRates)
}

func (p *EventProcessor) sendAdd(chatID int, username string) error {
	p.state[username] = stateAdd
	return p.tg.SendMessage(chatID, msgToAdd)
}

func (p *EventProcessor) sendDelete(chatID int, username string) error {
	p.state[username] = stateDel
	return p.tg.SendMessage(chatID, msgToDel)
}

func (p *EventProcessor) sendDeleteAll(chatID int, username string) error {
	p.db.DeleteAll(username)
	return p.tg.SendMessage(chatID, msgDelAll)
}

func (p *EventProcessor) Add(chatID int, username string, text string) error {
	err := p.db.Add(username, text)
	if err != nil {
		return e.Wrap(errAddingTokens, err)
	}

	p.state[username] = defaultState

	msg := msgAdd
	for _, token := range strings.Fields(text) {
		if isValid(token) {
			msg += fmt.Sprintf("\n- %s", token)
			fmt.Println(token, msg)
		}
	}

	return p.tg.SendMessage(chatID, msg)
}

func (p *EventProcessor) Del(chatID int, username string, text string) error {
	err := p.db.Delete(username, text)
	if err != nil {
		return e.Wrap(errDeletingTokens, err)
	}

	p.state[username] = defaultState

	msg := msgDel
	for _, token := range strings.Fields(text) {
		msg += fmt.Sprintf("\n- %s", token)
	}

	return p.tg.SendMessage(chatID, msg)
}

// Validate cryptocurrency name
func isValid(s string) bool {
	var list = []string{"bitcoin", "ethereum", "tether", "binance-coin", "usd-coin", "xrp", "solana", "cardano", "dogecoin", "ron", "ulti-collateral-dai", "olygon", "olkadot", "rapped-bitcoin", "itecoin", "itcoin-cash", "hainlink", "hiba-inu", "nus-sed-leo", "rueusd", "valanche", "tellar", "onero", "kb", "niswap", "thereum-classic", "inance-usd", "osmos", "itcoin-bep2", "ilecoin", "nternet-computer", "aker", "ido-dao", "rypto-com-coin", "echain", "uant", "ear-protocol", "ave", "tacks", "itcoin-sv", "he-graph", "lgorand", "edera-hashgraph", "ender-token", "ezos", "he-sandbox", "os", "njective-protocol", "heta", "xie-infinity", "lrond-egld", "infin-network", "horchain", "ecentraland", "antom", "ava", "cash", "axos-standard", "eo", "ynthetix-network-token", "low", "rust-wallet-token", "hiliz", "ucoin-token", "cash", "ei-protocol", "rax-share", "ota", "laytn", "urve-dao-token", "ocket-pool", "uobi-token", "ina", "onflux-network", "atetoken", "asper", "ydx", "tx-token", "ala", "ompound", "exo", "ootrade", "illiqa", "ash", "1inch", "asis-network", "asic-attention-token", "rweave", "nosis-gno", "ancakeswap", "nem", "tum", "elf", "emini-dollar", "olo", "just", "oopring", "onvex-finance", "elo", "njin-coin"}
	for _, val := range list {
		if val == s {
			return true
		}
	}
	return false
}
