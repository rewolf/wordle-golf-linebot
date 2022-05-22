package daemon

import (
	"errors"
	"fmt"
	"github.com/rewolf/wordle-golf-linebot/internal/store"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Daemon interface {
	Run() error
}

type wordleGolfLineBotDaemon struct {
	store  store.Store
	config Config
}

func New() (Daemon, error) {
	daemon := &wordleGolfLineBotDaemon{}

	if err := daemon.initConfig(); err != nil {
		return nil, err
	}
	if err := daemon.initStore(); err != nil {
		return nil, err
	}
	if err := daemon.initMessageHandler(); err != nil {
		return nil, err
	}
	return daemon, nil
}

func (d *wordleGolfLineBotDaemon) Run() error {
	log.Printf("Listening on %d", d.config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", d.config.Port), nil); err != nil {
		return err
	}
	return nil
}

func (d *wordleGolfLineBotDaemon) initConfig() error {
	var (
		port   int
		secret string
		token  string
		err    error
	)
	port, err = strconv.Atoi(os.Getenv("WGLB_PORT"))
	if err != nil {
		return fmt.Errorf("failed to parse WGLB_PORT for config: %w", err)
	}
	secret = os.Getenv("WGLB_CHANNEL_SECRET")
	token = os.Getenv("WGLB_CHANNEL_TOKEN")
	if secret == "" || token == "" {
		return errors.New("WGLB_CHANNEL_SECRET and WGLB_CHANNEL_TOKEN must have non-empty values")
	}
	d.config = Config{
		ChannelSecret: secret,
		ChannelToken:  token,
		Port:          port,
	}
	return nil
}

func (d *wordleGolfLineBotDaemon) initStore() error {
	//return errors.New("Store not implemented")
	return nil
}

func (d *wordleGolfLineBotDaemon) initMessageHandler() error {
	lineEventHandler, err := newLineEventHandler(
		d.config.ChannelSecret,
		d.config.ChannelToken,
	)
	if err != nil {
		return fmt.Errorf("could not create LINE event handler: %w", err)
	}
	http.Handle("/event", lineEventHandler)
	return nil
}
