package wechatutil

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
)

func AuthMessage(token string,timestamp string,nonce string,signature string) bool {
	var arr = []string{timestamp,nonce,token}
	sort.Strings(arr)
	var str = ""
	for t := 0;t< len(arr);t++ {
		str+=arr[t]
	}
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	hex_string_data := hex.EncodeToString(bs)
	if hex_string_data==signature {
		return true
	}else {
		return false
	}
}
