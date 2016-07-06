//sorted map and produce custom json string on the sorted map
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
)

type OrderedMap struct {
	M         map[string]interface{}
	ByKey     bool `description:"if false, then sort by Val"`
	Ascending bool
}

func (o OrderedMap) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}

	//sort
	lst := make([]interface{}, len(o.M))
	i := 0
	for k, v := range o.M {
		lst[i] = SortItem{k, v}
		i++
	}

	sortedList := []interface{}{}
	var err error
	if o.ByKey {
		sortedList, err = SortStructList(lst, "Key", o.Ascending)
	} else {
		sortedList, err = SortStructList(lst, "Val", o.Ascending)
	}
	check(err)

	//generate json string
	buf.WriteString("{")
	for i, v := range sortedList {
		if i != 0 {
			buf.WriteString(",")
		}

		key, err := json.Marshal(v.(SortItem).Key)
		if err != nil {
			return nil, err
		}
		buf.Write(key)
		buf.WriteString(":")

		value, err := json.Marshal(v.(SortItem).Val)
		if err != nil {
			return nil, err
		}
		buf.Write(value)
	}
	buf.WriteString("}")

	return buf.Bytes(), nil
}

func main() {
	m := OrderedMap{
		M: map[string]interface{}{
			"cake":   8,
			"fruit":  10,
			"banana": 8,
		},
		ByKey:     false,
		Ascending: false,
	}

	ln(m)
}

//====================================================================================================
type SortItem struct {
	Key string
	Val interface{}
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
