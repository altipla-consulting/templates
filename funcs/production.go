package funcs

import (
	"os"
)

func Development() bool {
	return os.Getenv("VERSION") == ""
}
