package ggpk

import (
	"bytes"
	"encoding/binary"
	"os"
)

func readUint32(f *os.File) (uint32, error) {
	buf := make([]byte, 4)
	if _, err := f.Read(buf); err != nil {
		return 0, err
	}
	var val uint32
	if err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func readUint64(f *os.File) (uint64, error) {
	buf := make([]byte, 8)
	if _, err := f.Read(buf); err != nil {
		return 0, err
	}
	var val uint64
	if err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func readTag(f *os.File) (string, error) {
	buf := make([]byte, 4)
	if _, err := f.Read(buf); err != nil {
		return "", err
	}
	return string(buf), nil
}
