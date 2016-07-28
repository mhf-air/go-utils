//sort map by value
package util

func mainMapValueSorter() {
	m := map[string]interface{}{
		"cake":   8,
		"fruit":  10,
		"banana": 8,
	}
	lst := SortMap(m, false, false)
	Ln(lst)

}

type SortItem struct {
	Key string
	Val interface{}
}

/*
 *sort map by key or value
 */
func SortMap(m map[string]interface{}, byKey, ascending bool) []interface{} {
	//sort
	lst := make([]interface{}, len(m))
	i := 0
	for k, v := range m {
		lst[i] = SortItem{k, v}
		i++
	}

	sortedList := []interface{}{}
	if byKey {
		sortedList, _ = SortStructList(lst, "Key", ascending)
	} else {
		sortedList, _ = SortStructList(lst, "Val", ascending)
	}
	return sortedList
}
