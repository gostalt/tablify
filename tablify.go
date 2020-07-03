package tablify

import (
	"fmt"
	"reflect"
	"strconv"
)

func separatorStruct(lens reflect.Value) string {
	fieldCount := lens.NumField()

	s := "+"
	for i := 0; i < fieldCount; i++ {
		fieldLen := int(lens.Field(i).Int())
		for j := 0; j < int(fieldLen); j++ {
			s = s + "-"
		}
		s = s + "+"
	}
	return s
}

func printTitle(values []interface{}, l reflect.Value) {
	title := "|"
	row := reflect.ValueOf(values[0])
	for i := 0; i < row.NumField(); i++ {
		n := row.Type().Field(i).Name
		disp := row.Type().Field(i).Tag.Get("tablify")
		if disp == "" {
			title = title + padForLength(n, int(l.FieldByName(n).Int())) + "|"
		} else {
			title = title + padForLength(disp, int(l.FieldByName(n).Int())) + "|"
		}
	}

	fmt.Println(title)
}

func printRow(value interface{}, l reflect.Value) {
	val := "|"
	row := reflect.ValueOf(value)
	for i := 0; i < row.NumField(); i++ {
		v := row.Field(i).String()
		if row.Field(i).Kind() == reflect.Int {
			v = strconv.Itoa(int(row.Field(i).Int()))
		}
		n := row.Type().Field(i).Name
		val = val + padForLength(v, int(l.FieldByName(n).Int())) + "|"
	}

	fmt.Println(val)
}

func PrintRowsStruct(values []interface{}) {
	if len(values) == 0 {
		fmt.Println("Unable to render table. No data exists.")
		return
	}
	l := lengths(values)
	sep := separatorStruct(l)

	fmt.Println(sep)

	vs := make([]interface{}, len(values))
	for i, v := range values {
		vs[i] = v
	}

	printTitle(vs, l)

	for _, v := range values {
		fmt.Println(sep)
		printRow(v, l)
	}

	fmt.Println(sep)
}

func padForLength(value string, length int) string {
	value = " " + value
	if len(value) == length {
		return value
	}

	spaces := length - len(value)
	for i := 0; i < spaces; i++ {
		value = value + " "
	}

	return value
}
