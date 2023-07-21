package kbucket

import (
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareAddrPorts(t *testing.T) {
	addr1, _ := netip.ParseAddrPort("1.1.1.1:6881")
	addr2, _ := netip.ParseAddrPort("127.0.0.1:6881")
	addr3, _ := netip.ParseAddrPort("[::FFFF:127.0.0.1]:6881")

	assert.False(t, CompareAddrPorts(addr1, addr2))
	assert.False(t, CompareAddrPorts(addr1, addr3))

	assert.True(t, CompareAddrPorts(addr1, addr1))
	assert.True(t, CompareAddrPorts(addr2, addr3))
}

func TestGenerateId(t *testing.T) {
	id, err := GenerateId()
	if assert.NoError(t, err) {
		assert.NotEmpty(t, id)
		assert.Len(t, id, 20)
	}

	id1, _ := GenerateId()
	id2, _ := GenerateId()
	assert.NotEqual(t, id1, id2)
}

func TestGenerateRandomBytes(t *testing.T) {
	rb, err := GenerateRandomBytes(20)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, rb)
		assert.Len(t, rb, 20)
	}

	rbs, err := GenerateRandomBytes(10)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, rbs)
		assert.Len(t, rbs, 10)
	}

	r1, _ := GenerateRandomBytes(10)
	r2, _ := GenerateRandomBytes(10)
	assert.NotEqual(t, r1, r2)
}
