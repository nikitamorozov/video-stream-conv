package tools

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func HashGenerator() string {
	h := sha1.New()
	h.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
