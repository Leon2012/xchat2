package idgen

import (
	"encoding/base64"
	"encoding/binary"

	sf "github.com/tinode/snowflake"
	"golang.org/x/crypto/xtea"
)

type IdGenerator struct {
	seq    *sf.SnowFlake
	cipher *xtea.Cipher
}

func (idGen *IdGenerator) Init(workerId uint, key []byte) error {
	var err error

	if idGen.seq == nil {
		idGen.seq, err = sf.NewSnowFlake(uint32(workerId))
	}
	if idGen.cipher == nil {
		idGen.cipher, err = xtea.NewCipher(key)
	}

	return err
}

// Get generates a unique weakly encryped id it so ids are random-looking.
func (idGen *IdGenerator) Get() uint64 {
	buf, err := getIdBuffer(idGen)
	if err != nil {
		return 0
	}
	return uint64(binary.LittleEndian.Uint64(buf))
}

// GetStr generates a unique id then returns it as base64-encrypted string.
func (idGen *IdGenerator) GetStr() string {
	buf, err := getIdBuffer(idGen)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(buf)[:11]
}

func getIdBuffer(idGen *IdGenerator) ([]byte, error) {
	var id uint64
	var err error
	if id, err = idGen.seq.Next(); err != nil {
		return nil, err
	}
	var src = make([]byte, 8)
	var dst = make([]byte, 8)
	binary.LittleEndian.PutUint64(src, id)
	idGen.cipher.Encrypt(dst, src)
	return dst, nil
}
