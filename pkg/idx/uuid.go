package idx

import (
	"encoding/base32"
	"github.com/google/uuid"
	"strings"
)

var base32Encoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567").
	WithPadding(base32.NoPadding)

func GenerateUuid() string {
	return uuid.NewString()
}

func GenerateUuid32() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

func GenerateUuidBase32() string {
	id := uuid.New()
	return base32Encoding.EncodeToString(id[:])
}
