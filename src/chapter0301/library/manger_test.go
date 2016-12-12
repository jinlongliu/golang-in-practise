package library

import (
	"testing"
	"fmt"
)

func TestMusicManagerOps(t *testing.T) {

	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty")
	}

	m0 := &Music{"1", "Love", "Tom", "http://www.baidu.com", "MP3"}
	mm.Add(m0)

	fmt.Println(mm.Get(0))

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed")
	}

	fmt.Println(mm.Find("Love"))
	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManger.Find() failed")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist || m.Name != m0.Name || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager.Find() failed")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicMangager.Remove() failed", err)
	}
}
