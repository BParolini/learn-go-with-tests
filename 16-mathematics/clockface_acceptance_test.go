package clockface

import (
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := Point{X: 150, Y: 150 + 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
