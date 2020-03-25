package totp

import (
	"encoding/binary"
	"time"

	"github.com/imagi-wa/hotp"
)

var (
	X  int64 = 30
	T0 int64 = 0
)

func TOTP(k []byte) (value uint32) {
	t := uint64((time.Now().Unix() - T0) / X)
	buf := make([]byte, binary.MaxVarintLen64)
	binary.BigEndian.PutUint64(buf, t)
	value = hotp.HOTP(k, buf)
	return
}
