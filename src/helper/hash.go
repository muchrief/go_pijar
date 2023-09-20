package helper

import "fmt"

func Hash(value string) string {
	enc := ""
	for _, v := range value[:] {
		s := fmt.Sprintf("%x", v)
		enc += s
	}

	return enc
}
