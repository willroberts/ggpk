package ggpk

import (
	"testing"
)

func TestGGPK(t *testing.T) {
	//f, err := NewGGPK("G:\\Path of Exile\\Content.ggpk")
	f, err := NewGGPK("C:\\Program Files (x86)\\Grinding Gear Games\\Path of Exile\\Content.ggpk")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(f)
}
