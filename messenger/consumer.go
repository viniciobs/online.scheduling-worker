package messenger

import (
	"context"
	"fmt"

	"github.com/online.scheduling-worker/config"
	"github.com/online.scheduling-worker/handlers"
	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.GetMessengerBroker()},
		Topic:   config.GetMessengerModalitiesEditTopic(),
		GroupID: config.GetMessengerModalitiesEditGroupId(),
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("could not read message " + err.Error())
		}

		if handlers.Handle(msg.Value) {
			r.CommitMessages(ctx, msg)
		}
	}
}
