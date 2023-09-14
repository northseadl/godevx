package validatex

import "regexp"

var (
	regexEmail       = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	regexPhoneNumber = regexp.MustCompile(`^\+[0-9]{1,3}-[0-9]{1,14}$`)
	regexURL         = regexp.MustCompile(`^(http|https|ftp)://[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	regexIP          = regexp.MustCompile(`^([0-9]{1,3}\.){3}[0-9]{1,3}$`)
)
