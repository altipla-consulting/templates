package funcs

import (
	"strconv"
	"strings"

	"github.com/ernestoalejo/aeimagesflags"
	"github.com/juju/errors"
)

func Thumbnail(servingURL string, strFlags string) (string, error) {
	if servingURL == "" || strFlags == "" {
		return "", nil
	}

	flags := aeimagesflags.Flags{
		ExpiresDays: 365,
	}
	for _, part := range strings.Split(strFlags, ";") {
		strFlag := strings.Split(part, "=")
		if len(strFlag) != 2 {
			return "", errors.New("all flags should be in the form key=value")
		}

		switch strings.TrimSpace(strFlag[0]) {
		case "width":
			n, err := strconv.ParseUint(strFlag[1], 10, 64)
			if err != nil {
				return "", errors.Trace(err)
			}
			flags.Width = n

		case "height":
			n, err := strconv.ParseUint(strFlag[1], 10, 64)
			if err != nil {
				return "", errors.Trace(err)
			}
			flags.Height = n

		case "square-crop":
			flags.SquareCrop = (strFlag[1] == "true")

		case "smart-square-crop":
			flags.SmartSquareCrop = (strFlag[1] == "true")

		case "original":
			flags.Original = (strFlag[1] == "true")

		default:
			return "", errors.Errorf("unknown flag: %s", strFlag[0])
		}
	}

	servingURL = strings.Replace(servingURL, "http://", "https://", 1)
	return aeimagesflags.Apply(servingURL, flags), nil
}
