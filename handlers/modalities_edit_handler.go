package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Action string

const (
	Delete Action = "Delete"
	Update Action = "Update"
)

type ModalitiesEdit struct {
	ModalityId uuid.UUID `json:"modality-id"`
	Action     Action    `json:"action"`
}

var strategies = map[Action]interface{}{
	Delete: removeModalityFromUsers,
	Update: warnUsersAboutChanging,
}

func Handle(content []byte) bool {
	var msg ModalitiesEdit

	if err := json.Unmarshal(content, &msg); err != nil {
		return false
	}

	strategy := strategies[msg.Action]
	err := strategy.(func(*uuid.UUID) error)(&msg.ModalityId)

	if err != nil {
		fmt.Printf("Error: %v ===> %v", msg, err)
	}

	log.Printf("read %v", msg)

	return true
}

func removeModalityFromUsers(modalityId *uuid.UUID) error {
	return nil
}

func warnUsersAboutChanging(modalityId *uuid.UUID) error {
	return nil
}

// func getUsersByModalityId(modalityId *uuid.UUID) (usersId []uuid.UUID, err error) {

// }
