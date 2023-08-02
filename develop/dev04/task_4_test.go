package task_4

import (
	"reflect"
	"testing"
)

func TestCreateAnogramDictionary(t *testing.T) {
	tastCases := []struct {
		entrance []string
		want     map[string][]string
	}{
		{
			entrance: []string{"ек", "", "ek", "y"},
			want:     map[string][]string{},
		},
		{
			entrance: []string{},
			want:     map[string][]string{},
		},
		{
			entrance: []string{"ек", "норм", "ek", "y", "нрмо"},
			want: map[string][]string{
				"норм": []string{"норм", "нрмо"},
			},
		},
		{
			entrance: []string{"ек", "норм", "ek", "y", "нрмо", "норм"},
			want: map[string][]string{
				"норм": []string{"норм", "нрмо"},
			},
		},
	}

	for id, ts := range tastCases {
		got := СreateAnogramDictionary(ts.entrance)
		if !reflect.DeepEqual(got, ts.want) {
			t.Errorf("%d string unpacking() = %v, want %v", id, got, ts.want)
		}
	}
}
