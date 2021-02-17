package ggpk

import (
	"testing"
)

func TestGGPK(t *testing.T) {
	f, err := NewGGPK("G:\\Path of Exile\\Content.ggpk")
	if err != nil {
		t.Fatal(err)
	}
	if err = f.GetNumRecords(); err != nil {
		t.Fatal(err)
	}
	if err = f.GetOffsets(); err != nil {
		t.Fatal(err)
	}
}
