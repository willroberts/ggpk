package ggpk

import "testing"

func TestGGPK(t *testing.T) {
	var path = "C:\\Program Files (x86)\\Grinding Gear Games\\Path of Exile\\Content.ggpk"
	if _, err := NewGGPK(path); err != nil {
		t.Fatal(err)
	}
}
