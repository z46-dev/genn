package shared

import (
	cRand "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

var uniqueIDCounter atomic.Uint64

func init() {
	var (
		entropy  [8]byte
		hostname string
		err      error
		seed     uint64 = 14695981039346656037
	)

	if _, err = cRand.Read(entropy[:]); err != nil {
		panic(fmt.Sprintf("crypto/rand.Read: %v", err))
	}

	for _, value := range entropy {
		seed = (seed ^ uint64(value)) * 1099511628211
	}

	if hostname, err = os.Hostname(); err == nil {
		for index := 0; index < len(hostname); index++ {
			seed = (seed ^ uint64(hostname[index])) * 1099511628211
		}
	}

	seed ^= uint64(os.Getpid()) * 0x9e3779b97f4a7c15
	seed ^= uint64(time.Now().UnixNano())
	uniqueIDCounter.Store(seed)
}

// UniqueID generates a compact process-unique identifier.
func UniqueID() (id string) {
	var (
		raw     [12]byte
		encoded [16]byte
	)

	binary.BigEndian.PutUint32(raw[:4], uint32(time.Now().Unix()))
	binary.BigEndian.PutUint64(raw[4:], uniqueIDCounter.Add(1))

	base64.RawURLEncoding.Encode(encoded[:], raw[:])
	id = string(encoded[:])
	return
}
