package avatarme

import "hash/fnv"

// Encrypt transforms given data into a hash, which will be later transformed into an Identicon.
func Encrypt(b []byte) (uint64, error) {
	newHash := fnv.New64a()
	_, err := newHash.Write(b)
	if err != nil {
		return uint64(0), err
	}
	return newHash.Sum64(), nil
}