package bcrypt

import "regexp"

var reFormatBcrypt = regexp.MustCompile(`^\$2[aby]\$[0-9]{2}\$[./A-Za-z0-9]{53}$`)

// IsBcryptHash check string is hash bcrypt.
func IsBcryptHash(in string) bool {
	return reFormatBcrypt.MatchString(in)
}
