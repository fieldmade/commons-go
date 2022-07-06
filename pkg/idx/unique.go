package idx

import (
	"encoding/hex"
	"github.com/google/uuid"
)

func GenerateUniqueId() string {
	id := uuid.New()
	return hex.EncodeToString(id[:])
}

func GenerateUniqueToken() string {
	u1 := GenerateUniqueId()
	u2 := GenerateUniqueId()

	return u1 + u2
}
