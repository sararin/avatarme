package uniqhash

import (
  
	"hash/fnv"

	"fmt"

)



type Hashed struct {

	Hash uint64

}



func (h Hashed) String() string {

	return fmt.Sprintf("%v", h.Hash)

}



func Encrypt(b []byte) (*Hashed, error) {

	newHash := fnv.New64a()

	_, err := newHash.Write(b)

	if err != nil {

		return &Hashed{}, err

	}

	return &Hashed{newHash.Sum64()}, nil

}