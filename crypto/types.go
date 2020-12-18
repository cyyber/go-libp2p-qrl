package crypto

import (
	"github.com/libp2p/go-libp2p-core/crypto"
	pb "github.com/libp2p/go-libp2p-core/crypto/pb"
)

const (
	LATTICE = iota
)

var extendedKeyTypes = []int {
	LATTICE,
}

var isLoaded bool

var mapExtendedKeyTypes = map[int]pb.KeyType {
	// This will be filled when LoadAllExtendedKeyTypes are called
}

func LoadAllExtendedKeyTypes() {
	if isLoaded {
		return
	}
	keyLen := pb.KeyType(len(crypto.KeyTypes))
	for _, data := range extendedKeyTypes {
		absoluteExtendedKey := keyLen + pb.KeyType(data)
		crypto.PubKeyUnmarshallers[absoluteExtendedKey] = UnmarshalDilithiumPublicKey
		crypto.PrivKeyUnmarshallers[absoluteExtendedKey] = UnmarshalDilithiumPrivateKey
		mapExtendedKeyTypes[LATTICE] = absoluteExtendedKey
		keyLen++
	}
	isLoaded = true
}
