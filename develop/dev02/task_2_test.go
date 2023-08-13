package main

import (
	"fmt"
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
		// Возможно нужно поменять логику впо чтобы повторение любых
		// чисел без слеша даввало пустую подстроку???{
		// 	entrance: "a10",
		// 	want:     "a",
		// },
	}

	for id, ts := range tastCases {
		got := StringUnpacking(ts.entrance)
		fmt.Println(got, ts.want, reflect.DeepEqual(got, ts.want))
		if !reflect.DeepEqual(got, ts.want) {
			t.Errorf("%d string unpacking() = %v, want %v", id, got, ts.want)
		}
	}
}
