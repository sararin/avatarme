package avatarme

import (
	"github.com/SaraTrawnik/avatarme/uniqhash"
	"github.com/SaraTrawnik/avatarme/drawimg"
	"encoding/base64"
	"os"
)

func Draw(b []byte) io.WriterCloser {
	a, _ := uniqhash.Encrypt(b)
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(drawimg.FromHash(a))
	encoder.Close()
}