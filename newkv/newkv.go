package newkv

import (
	"strings"

	"github.com/karthik-vijayan93/go-operator-kv/kv"
)

func Kvprint(s string) string {
	kv.Kv()
	return strings.ToUpper(s)
}
