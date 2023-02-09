package messenger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/online.scheduling-worker/config"
	"github.com/segmentio/kafka-go"
)

type deleteObjects struct {
	UserId     uuid.UUID
	ModalityId uuid.UUID
}

var uri = config.Schedule_Api_Uri + "/api/schedules"

func Work_RemoveSchedulesFrom(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.GetMessengerBroker()},
		Topic:   "DELETED_OBJECTS",
		GroupID: "DELETED_OBJECTS_CONSUMER",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("could not read message " + err.Error())
		}

		if removeSchedulesFrom(msg.Value) {
			r.CommitMessages(ctx, msg)
		}
	}
}

func removeSchedulesFrom(content []byte) bool {
	var data deleteObjects

	if err := json.Unmarshal(content, &data); err != nil {
		return false
	}

	bodyRequest, _ := json.Marshal(map[string]uuid.UUID{
		"user-id":     data.UserId,
		"modality-id": data.ModalityId,
	})

	payload := bytes.NewBuffer(bodyRequest)
	resp, err := http.Post(uri, "application/json", payload)

	if err != nil {
		log.Printf("[Deleted_objects_consumer] SUCCESS %s\n%s", bodyRequest, err)
		return false
	}

	if resp.StatusCode != http.StatusNoContent {
		log.Printf("[Deleted_objects_consumer] FAIL %s", bodyRequest)
		return false
	}

	log.Printf("[Deleted_objects_consumer] SUCCESS %s", bodyRequest)

	return true
}
