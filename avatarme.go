// Package avatarme provides easy way to create very simple Identicons.
// It can be used both to return a base64 encoded string and save image in PNG format.
package avatarme

import (
	"encoding/base64"
	"os"
	"image"
	"image/png"
)

// Draw transforms given data into identicon and saves it as 5x5 PNG file with given filename.
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

// DrawToBase64 transforms given data into a base64 encoded string.
// It empty string nothing upon failure.	
func DrawToBase64(b []byte) (string, error) {
	h, err := Encrypt(b)
	if err != nil {
		return "", err
	}

	result, err := FromHash(h)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(result), nil
}