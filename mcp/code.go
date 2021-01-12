package mcp

import (
	"encoding/binary"
	"strconv"
)

// PLC Data communication code.
// This item is operating byte order.
type Code int

const (
	// Ascii code is normal mode.
	// Stored from upper byte to lower byte.
	Ascii Code = iota

	//　Binary code is approximately half the amount of communication data compared to communication using ASCII code
	// Stored from lower byte to upper byte.
	Binary
)

func (c Code) EncodeHex(s string) ([]byte, error) {
	if c == Ascii {
		return []byte(s), nil
	}

	hexCode, err := strconv.ParseUint(s, 16, 16)
	if err != nil {
		return nil, err
	}

	buff := make([]byte, 2)
	binary.LittleEndian.PutUint16(buff, uint16(hexCode))
	return buff, nil
}
