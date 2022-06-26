package pkg

import "os"

func Env(key, def string) string {
	if x := os.Getenv(key); x != "" {
		return x
	}
	return def
}
