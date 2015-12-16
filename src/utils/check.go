package utils

import (
	"crypto/md5"
	"encoding/hex"
	"utils/logs"
)

func DataIntegrityCheck(data, sum []byte) bool {
	csum := GetMd5(data)

	if len(sum) != len(csum) {
		return false
	}

	for i := 0; i < len(csum); i++ {
		if csum[i] != sum[i] {
			return false
		}
	}

	return true
}

func GetMd5(data interface{}) []byte {
	c := md5.New()

	switch data.(type) {
	case []byte:
		c.Write(data.([]byte))
	case string:
		c.Write([]byte(data.(string)))
	default:
		logs.Logger.Debug("unsupported types")
	}
	return c.Sum(nil)
}

func GetMd5String(data []byte) string {
	sum := GetMd5(data)
	return hex.EncodeToString(sum)
}
