package tablify

import (
	"fmt"
	"reflect"
	"strconv"
)

func createSeparator(lens reflect.Value) string {
	fieldCount := lens.NumField()

	s := join
	for i := 0; i < fieldCount; i++ {
		fieldLen := int(lens.Field(i).Int())
		for j := 0; j < int(fieldLen); j++ {
			s = s + horizontal
		}
		s = s + join
	}
	return s
}

// fieldTitle gets the name that should appear at the top of the column. If the struct
// field has a `tablify` tag attached, this will be used instead of the field name.
//
// Returns two values, `name` and `preferred`.
func fieldTitle(field reflect.StructField) (string, string) {
	if field.Tag.Get("tablify") != "" {
		return field.Name, field.Tag.Get("tablify")
	}

	return field.Name, field.Name
}

func printTitle(values interface{}, l reflect.Value) {
	t := vertical
	row := reflect.ValueOf(values)
	for i := 0; i < row.NumField(); i++ {

		n, p := fieldTitle(row.Type().Field(i))
		t = t + padForLength(p, int(l.FieldByName(n).Int())) + vertical
	}

	fmt.Println(t)
}

func printRow(value interface{}, l reflect.Value) {
	val := vertical
	row := reflect.ValueOf(value)
	for i := 0; i < row.NumField(); i++ {
		v := row.Field(i).String()
		if row.Field(i).Kind() == reflect.Int {
			v = strconv.Itoa(int(row.Field(i).Int()))
		}
		if row.Field(i).Kind() == reflect.Bool {
			if row.Field(i).Bool() {
				v = "true"
			} else {
				v = "false"
			}
		}
		n := row.Type().Field(i).Name
		val = val + padForLength(v, int(l.FieldByName(n).Int())) + vertical
	}

	fmt.Println(val)
}

func PrintRowsStruct(values []interface{}) {
	if !validate(values) {
		return
	}

	l := lengths(values)
	sep := createSeparator(l)

	// Start the table.
	fmt.Println(sep)

	vs := make([]interface{}, len(values))
	for i, v := range values {
		vs[i] = v
	}

	printTitle(vs[0], l)

	for _, v := range values {
		fmt.Println(sep)
		printRow(v, l)
	}

	// End the table.
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

func validate(values []interface{}) bool {
	if len(values) == 0 {
		fmt.Println("Unable to render table. No data exists.")
		return false
	}

	return true
}
