//sort map by value
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
)

func main() {
	m := map[string]interface{}{
		"cake":   8,
		"fruit":  10,
		"banana": 8,
	}
	lst := SortMap(m, false, false)
	ln(lst)

}

type SortItem struct {
	Key string
	Val interface{}
}

/*
 *sort map by value
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

func SortStructList(lst []interface{}, fieldName string, ascending bool) ([]interface{}, error) {
	var err error
	if len(lst) == 0 {
		return nil, errors.New("empty list")
	}
	v := reflect.ValueOf(lst[0])
	fieldMap := make(map[string]bool, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		name := reflect.TypeOf(lst[0]).Field(i).Name
		fieldMap[name] = true
	}
	if fieldMap[fieldName] == false {
		return nil, errors.New("wrong field name")
	}

	s := StructValueSorter{
		lst,
		func(a, b interface{}) bool {
			va := reflect.ValueOf(a).FieldByName(fieldName).Interface()
			vb := reflect.ValueOf(b).FieldByName(fieldName).Interface()
			ka := reflect.TypeOf(va).Kind()
			kb := reflect.TypeOf(vb).Kind()
			if ka != kb {
				err = errors.New("that shouldn't happen!")
			} else {
				switch ka {
				case reflect.Int:
					if va.(int) < vb.(int) {
						return true
					}
				case reflect.Int64:
					if va.(int64) < vb.(int64) {
						return true
					}
				case reflect.Float32:
					if va.(float32) < vb.(float32) {
						return true
					}
				case reflect.Float64:
					if va.(float64) < vb.(float64) {
						return true
					}
				case reflect.String:
					if va.(string) < vb.(string) {
						return true
					}
				default:
					err = errors.New("sorry, sortStructList() now only support struct fields whose type are int or string")
				}
			}

			return false
		},
	}

	if ascending {
		sort.Sort(s)
	} else {
		sort.Sort(sort.Reverse(s))
	}
	return s.Content, err
}

//====================================================================================================
type StructValueSorter struct {
	Content   []interface{}
	FrontFunc func(a, b interface{}) bool
}

func (s StructValueSorter) Len() int {
	return len(s.Content)
}

func (s StructValueSorter) Swap(i, j int) {
	s.Content[i], s.Content[j] = s.Content[j], s.Content[i]
}

func (s StructValueSorter) Less(i, j int) bool {
	return s.FrontFunc(s.Content[i], s.Content[j])
}

//====================================================================================================
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ln(a ...interface{}) {
	for _, v := range a {
		b, err := json.MarshalIndent(v, "", "  ")
		check(err)

		//ignore the case when \n is in " "
		lines := strings.Split(string(b), "\n")
		lst := []string{}
		for _, l := range lines {
			pureline := strings.TrimSpace(l)
			frontBlank := strings.Repeat(" ", len(l)-len(pureline))
			if strings.HasPrefix(pureline, "]") || strings.HasPrefix(pureline, "}") {
				continue
			}
			if strings.HasSuffix(pureline, ",") {
				pureline = pureline[:len(pureline)-1]
			}
			lst = append(lst, frontBlank+pureline+"\n")
		}
		str := strings.Join(lst, "")
		fmt.Printf("%s", str)
	}
}
