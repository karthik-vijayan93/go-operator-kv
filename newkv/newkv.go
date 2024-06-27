package newkv

import (
	"fmt"
	"strings"

	"github.com/karthik-vijayan93/go-operator-kv/kv"
)

func Kvprint(s string) string {
	fmt.Println(kv.Kv())
	return strings.ToUpper(s)
}
