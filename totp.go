package totp

import (
	"time"

	"github.com/imagi-wa/hotp"
)

var (
	X  int64 = 30
	T0 int64 = 0
)

func TOTP(k []byte) (value uint32) {
	t := uint64((time.Now().Unix() - T0) / X)
	value = hotp.HOTP(k, t)
	return
}
