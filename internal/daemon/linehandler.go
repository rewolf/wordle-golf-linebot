package daemon

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/line/line-bot-sdk-go/v7/linebot/httphandler"
	"log"
	"net/http"
)

func newLineEventHandler(channelSecret, channelToken string) (http.Handler, error) {

	handler, err := httphandler.New(
		channelSecret,
		channelToken,
	)
	if err != nil {
		return nil, err
	}

	// Setup HTTP Server for receiving requests from LINE platform
	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
		bot, err := handler.NewClient()
		if err != nil {
			log.Print(err)
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					log.Println(message.Text)
					log.Println(message.ID)
					log.Println(message.Mention)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
					if _, err = bot.PushMessage(event.Source.GroupID, linebot.NewTextMessage("hi\ntest<br>test2*cool*\n⬛️\U0001F7E8\U0001F7E8⬛️⬛️\n\U0001F7E8\U0001F7E9\U0001F7E8\U0001F7E8\U0001F7E8\n\U0001F7E9\U0001F7E9\U0001F7E9\U0001F7E9\U0001F7E9")).Do(); err != nil {
						log.Print(err)
					}

					container := &linebot.BubbleContainer{
						Type: linebot.FlexContainerTypeBubble,
						Body: &linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeHorizontal,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type: linebot.FlexComponentTypeText,
									Text: "Hello,",
								},
								&linebot.TextComponent{
									Type: linebot.FlexComponentTypeText,
									Text: "World!",
								},
							},
						},
					}
					flexMsg := linebot.NewFlexMessage("You can't view this type of message...", container)
					if _, err = bot.PushMessage(event.Source.GroupID, flexMsg).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	return handler, nil
}
