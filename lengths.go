package tablify

import (
	"reflect"
	"strconv"
)

func lengths(values []interface{}) reflect.Value {
	// Convert to struct
	s := reflect.ValueOf(values[0])
	fieldCount := s.NumField()

	fs := make([]reflect.StructField, fieldCount)

	for i := 0; i < fieldCount; i++ {
		n := s.Type().Field(i).Name
		fs[i] = reflect.StructField{Name: n, Type: reflect.TypeOf(1)}
	}

	typ := reflect.StructOf(fs)
	fin := reflect.New(typ).Elem()

	// Set the initial value to be the length of the field name + 2
	for i := 0; i < fieldCount; i++ {
		n, p := fieldTitle(s.Type().Field(i))
		fin.FieldByName(n).SetInt(int64(len(p) + 2))
	}

	// Iterate through field values and set the length to be the length of it, if greater
	// for i := 0; i < values.
	for _, v := range values {
		for i := 0; i < fieldCount; i++ {
			val := reflect.ValueOf(v)
			k := val.Type().Field(i).Name

			var valLen int64
			switch val.Field(i).Kind() {
			case reflect.String:
				valLen = int64(len(val.Field(i).String()) + 2)
			case reflect.Int:
				valLen = int64(len(strconv.Itoa(int(val.Field(i).Int()))) + 2)
			case reflect.Bool:
				f := val.Field(i).Bool()
				if f {
					valLen = 6 // ` True `
				} else {
					valLen = 7 // ` False `
				}
			}

			if valLen > fin.FieldByName(k).Int() {
				fin.FieldByName(k).SetInt(valLen)
			}
		}
	}

	return fin
}
