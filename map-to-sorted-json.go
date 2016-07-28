//sorted map and produce custom json string on the sorted map
package util

import (
	"bytes"
	"encoding/json"
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
	Check(err)

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

func localMain() {
	m := OrderedMap{
		M: map[string]interface{}{
			"cake":   8,
			"fruit":  10,
			"banana": 8,
		},
		ByKey:     false,
		Ascending: false,
	}

	Ln(m)
}
