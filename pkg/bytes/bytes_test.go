package bytes_test

import (
	"bytes"
	"encoding/hex"
	"math/rand"
	"testing"

	cbytes "github.com/joincivil/go-common/pkg/bytes"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// TestByte32Utils tests encoding, then decoding from byte[32] to string
func TestByte32Utils(t *testing.T) {
	hexStr, _ := randomHex(32)
	testBytes := []byte(hexStr)
	testFixed := [32]byte{}
	copy(testFixed[:], testBytes)

	result := cbytes.Byte32ToHexString(testFixed)
	if result == "" {
		t.Error("Empty result for string")
	}

	bys, err := cbytes.HexStringToByte32(result)
	if err != nil {
		t.Errorf("Should not have returned an error: err: %v", err)
	}

	if !bytes.Equal(testFixed[:], bys[:]) {
		t.Error("[32]bytes are not equal")
	}
}