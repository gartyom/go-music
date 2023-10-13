package helpers

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	id := uuid.New().String()
	id = strings.Replace(id, "-", "", -1)[3:14]
	return id
}
