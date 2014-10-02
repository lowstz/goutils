package uuid

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

func GenUUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func GenIntUUID() int64 {
	id := time.Now().UnixNano()
	return id
}

func GenMd5(org string, key string) string {
	h := md5.New()
	io.WriteString(h, org+key)
	md := fmt.Sprintf("%x", h.Sum(nil))
	return md
}
