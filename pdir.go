package ggpk

import "os"

type PDir struct {
	numRecords uint32
	name       string // uint16 utf-16le * uint32 length
	childCount uint32
	signature  uint8
	children   []*ChildNode
}

type ChildNode struct {
	hash   uint32
	offset uint64
}

func NewPDir(file *os.File, offset uint64) (*PDir, error) {
	return nil, nil
}

func (g *GGPK) FindPDirs(file *os.File) error {
	return nil
}
