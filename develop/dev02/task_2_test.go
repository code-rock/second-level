package main

import (
	"reflect"
	"testing"
)

func TestSimpleStringUnpacking(t *testing.T) {
	tastCases := []struct {
		entrance string
		want     string
	}{
		{
			entrance: "a4bc2d5e",
			want:     "aaaabccddddde",
		}, {
			entrance: `qwe\4\5`,
			want:     "qwe45",
		},
		{
			entrance: `qwe\45`,
			want:     "qwe44444",
		},
		{
			entrance: `qwe\\5`,
			want:     `qwe\\\\\`,
		},
		{
			entrance: "abcd",
			want:     "abcd",
		},
		{
			entrance: "45",
			want:     "",
		},
		{
			entrance: "",
			want:     "",
		},
	}

	for id, ts := range tastCases {
		s := SStringUnpackage{}
		got := s.unpacking(ts.entrance)
		if !reflect.DeepEqual(got, ts.want) {
			t.Errorf("%d string unpacking() = %v, want %v", id, got, ts.want)
		}
	}
}
