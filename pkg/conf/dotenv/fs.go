package dotenv

import (
	"os"
	"path/filepath"
)

func toAbs(f string) string {
	abs, err := filepath.Abs(f)
	if err != nil {
		return f
	}
	return abs
}

func dir(f string) string {
	return filepath.Dir(toAbs(f))
}

func exists(f string) bool {
	_, e := os.Stat(toAbs(f))
	if e == nil {
		return true
	}
	return false
}
