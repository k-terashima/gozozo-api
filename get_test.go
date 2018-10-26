package gozozo

import "testing"

func TestGetRanking(t *testing.T) {
	p := new(SnapRanking)
	if err := p.GetRanking(); err != nil {
		t.Error(err)
	}
}
