package avatarme

import (
	"hash/fnv"
	"errors"
)

// Encrypt transforms given data into a hash, which will be later transformed into an Identicon.
func Encrypt(b []byte) (uint64, error) {
	if len(b) == 0 {
		return uint64(0), errors.New("empty byte array")
	}
	newHash := fnv.New64a()
	_, err := newHash.Write(b)
	if err != nil {
		return uint64(0), err
	}
	return newHash.Sum64(), nil
}