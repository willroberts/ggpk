package ggpk

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// GGPK is a bin-packed container for game data.
type GGPK struct {
	file       *os.File
	numRecords uint32
	childCount uint32
	offsets    []uint64
}

// NewGGPK loads and initializes a GGPK file.
func NewGGPK(filename string) (*GGPK, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	g := &GGPK{file: f}

	numRecords, err := readUint32(f)
	if err != nil {
		return nil, err
	}
	g.numRecords = numRecords
	log.Print("records:", g.numRecords) // debug

	tag, err := readTag(f)
	if err != nil {
		return nil, err
	}
	if tag != "GGPK" {
		msg := fmt.Sprintf("expected GGPK tag, got %s", tag)
		return nil, errors.New(msg)
	}

	childCount, err := readUint32(f)
	if err != nil {
		return nil, err
	}
	g.childCount = childCount
	log.Print("child count:", childCount) // debug

	offsets := make([]uint64, 0)
	for i := uint32(0); i < g.numRecords; i++ {
		offset, err := readUint64(f)
		if err != nil {
			return nil, err
		}
		offsets = append(offsets, offset)
	}
	g.offsets = offsets
	log.Print("offsets:", offsets) // debug

	return g, nil
}
