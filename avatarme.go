package avatarme

import (
	"github.com/SaraTrawnik/avatarme/uniqhash"
	"github.com/SaraTrawnik/avatarme/drawimg"
	"encoding/base64"
	"os"
	"image"
	"image/png"
)

func Draw(b []byte, filename string) error {
	base, err := DrawToBase64(b)
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base))
	m, _, err := image.Decode(reader)
	if err != nil { return err }
	
	f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE, 0777)
	if err != nil { return err }
	defer f.Close()
	
	err = png.Encode(f, m)
	if err != nil { return err }
	
	return nil
}
	
func DrawToBase64(b []byte) (string, error) {
	h, err := uniqhash.Encrypt(b)
	if err != nil {
		return "", err
	}

	result, err := drawimg.FromHash(h)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(result), nil
}