package ggpk

import (
	"testing"
)

func TestGGPK(t *testing.T) {
	var (
		path = "G:\\Path of Exile\\Content.ggpk" // desktop
		//path = "C:\\Program Files (x86)\\Grinding Gear Games\\Path of Exile\\Content.ggpk" // laptop
	)

	if _, err := NewGGPK(path); err != nil {
		t.Fatal(err)
	}
}
