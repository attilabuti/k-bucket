package kbucket

import (
	"crypto/rand"
	"crypto/sha1"
	"net/netip"
)

// CompareAddrPorts compares two netip.AddrPorts.
// It will return true if the two AddrPorts are equal, false otherwise.
func CompareAddrPorts(a, b netip.AddrPort) bool {
	if a.Addr().Unmap().Compare(b.Addr().Unmap()) == 0 && a.Port() == b.Port() {
		return true
	}

	return false
}

// GenerateId generates a random 160-bit ID (SHA-1).
// It will return an error if the system's secure random number generator fails
// to function correctly, in which case the caller should not continue.
func GenerateId() ([]byte, error) {
	b, err := GenerateRandomBytes(20)
	if err != nil {
		return nil, err
	}

	h := sha1.New()
	h.Write(b)

	return h.Sum(nil), nil
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random number generator fails
// to function correctly, in which case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
