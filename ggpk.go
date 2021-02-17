package ggpk

import (
	"bytes"
	"encoding/binary"
	"os"
)

type GGPK struct {
	file       *os.File
	numRecords int32
	offsets    []int64
}

func NewGGPK(filename string) (*GGPK, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return &GGPK{file: f}, nil
}

func (g *GGPK) GetNumRecords() error {
	buf := make([]byte, 4)
	if _, err := g.file.Read(buf); err != nil {
		return err
	}
	var val int32
	if err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, &val); err != nil {
		return err
	}
	g.numRecords = val
	return nil
}

func (g *GGPK) GetOffsets() error {
	offsets := make([]int64, 0)
	for i := int32(0); i < g.numRecords; i++ {
		buf := make([]byte, 8)
		if _, err := g.file.Read(buf); err != nil {
			return err
		}
		var val int64
		if err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, &val); err != nil {
			return err
		}
		offsets = append(offsets, val)
	}
	g.offsets = offsets
	return nil
}
