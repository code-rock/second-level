package task_4

import (
	"sort"
)

func sortString(str string) string {
	arr := []rune(str)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return string(arr)
}

func sortUnique(arr []string) []string {
	var unique []string
	for i := 0; i < len(arr); i++ {
		if len(unique) == 0 || arr[i] != unique[len(unique)-1] {
			unique = append(unique, arr[i])
		}
	}
	return unique
}

func СreateAnogramDictionary(arr []string) map[string][]string {
	sortedWord := make(map[string]string)
	dic := make(map[string][]string)

	for i := 0; i < len(arr); i++ {
		sorted := sortString(arr[i])
		initial, ok := sortedWord[sorted]

		if !ok {
			sortedWord[sorted] = arr[i]
			initial = arr[i]
		}
		dic[initial] = append(dic[initial], arr[i])
	}

	for k, v := range dic {
		sort.Strings(v)
		dic[k] = sortUnique(v)
		if len(dic[k]) < 2 {
			delete(dic, k)
		}

	}

	return dic
}

func main() {}
