package uuid

import (
	"crypto/rand"
	"fmt"
)

func MustNewUUID() string {
	uuid := make([]byte, 16)

	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		panic(err.Error())
	}

	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
