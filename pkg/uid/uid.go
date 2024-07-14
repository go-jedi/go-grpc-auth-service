package uid

import (
	"crypto/rand"
	"math/big"
)

var (
	chars   = []rune("0123456789ABCDEFHKLMNPQRSTUVWXYZabcdefghkmnpqrstuvwxyz")
	special = []rune(" !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
)

func GenUID(cnt int) (string, error) {
	// characters used in UID
	uid := make([]rune, cnt)

	for i := range uid {
		// generate a random index and select a character from chars
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}

		uid[i] = chars[num.Int64()]
	}

	return string(uid), nil
}

func GenSpecialUID(cnt int) (string, error) {
	// characters used in UID
	uid := make([]rune, cnt)
	sc := []rune(string(chars) + string(special))

	for i := range uid {
		// generate a random index and select a character from chars
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(sc))))
		if err != nil {
			return "", err
		}

		uid[i] = sc[num.Int64()]
	}

	return string(uid), nil
}
