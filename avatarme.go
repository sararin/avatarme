package avatarme

import (
	"github.com/SaraTrawnik/avatarme/uniqhash"
	"github.com/SaraTrawnik/avatarme/drawimg"
	"encoding/base64"
	"os"
)

func Draw(b []byte) (string, error) {
	a, err := uniqhash.Encrypt(b)
	if err != nil {
		return "", err
	}

	result, err := drawimg.FromHash(uint64(10555899939929309))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(result), nil
}