package encryption

import (
	"crypto/md5"
	"encoding/hex"
)

func DoMd5(value string) string {
	md5Ret := md5.New()
	md5Ret.Write([]byte(value))
	result := hex.EncodeToString(md5Ret.Sum(nil))
	return result
}
