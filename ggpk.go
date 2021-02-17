package ggpk

import (
	"bytes"
	"encoding/binary"
	"os"
)

// GGPK is a bin-packed container for game data.
type GGPK struct {
	file       *os.File
	numRecords int32
	offsets    []int64
}

// NewGGPK loads and initializes a GGPK file.
func NewGGPK(filename string) (*GGPK, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	g := &GGPK{file: f}

	if err = g.getNumRecords(); err != nil {
		return nil, err
	}

	if err = g.getOffsets(); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *GGPK) getNumRecords() error {
	v, err := readInt32(g.file)
	if err != nil {
		return err
	}
	g.numRecords = v
	return nil
}

func (g *GGPK) getOffsets() error {
	offsets := make([]int64, 0)
	for i := int32(0); i < g.numRecords; i++ {
		v, err := readInt64(g.file)
		if err != nil {
			return err
		}
		offsets = append(offsets, v)
	}
	g.offsets = offsets
	return nil
}

func readInt32(f *os.File) (int32, error) {
	buf := make([]byte, 4)
	if _, err := f.Read(buf); err != nil {
		return 0, err
	}
	var val int32
	if err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func readInt64(f *os.File) (int64, error) {
	buf := make([]byte, 8)
	if _, err := f.Read(buf); err != nil {
		return 0, err
	}
	var val int64
	if err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}
