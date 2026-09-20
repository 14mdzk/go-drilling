package main

import "testing"

// go test -bench=.
// go test -bench=. -benchmem
//
// Repeat using concatenation will resulting in go copying the memory since string is immutable
// RepeatBuilder using string builder to avoid go to copying the memory
//
// result:
// BenchmarkRepeat-10           	19526188	        61.01 ns/op	      16 B/op	       4 allocs/op
// BenchmarkRepeatBuilder-10    	55500952	        22.13 ns/op	       8 B/op	       1 allocs/op

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

func BenchmarkRepeatBuilder(b *testing.B) {
	for b.Loop() {
		RepeatBuilder("a")
	}
}
