package tablify

import (
	"fmt"
	"reflect"
	"strconv"
)

type tablify struct {
	h string
	v string
	j string
}

type Opts struct {
	Horizontal string
	Vertical   string
	Join       string
}

func New() tablify {
	return NewFromOpts(Opts{})
}

func NewFromOpts(opts Opts) tablify {
	t := tablify{
		h: opts.Horizontal,
		v: opts.Vertical,
		j: opts.Join,
	}

	if t.h == "" {
		t.h = horizontalDefault
	}

	if t.v == "" {
		t.v = verticalDefault
	}

	if t.j == "" {
		t.j = joinDefault
	}

	return t
}

func (t tablify) Struct(vals interface{}) {
	t.printRowsStruct(interfaceToInterfaceSlice(vals))
}

func interfaceToInterfaceSlice(vals interface{}) []interface{} {
	r := reflect.ValueOf(vals)
	len := r.Len()
	fmt.Println(len)

	is := make([]interface{}, len)
	for i := 0; i < len; i++ {
		is[i] = r.Index(i).Interface()
	}

	return is
}

func (t tablify) createSeparator(lens reflect.Value) string {
	fieldCount := lens.NumField()

	s := t.j
	for i := 0; i < fieldCount; i++ {
		fieldLen := int(lens.Field(i).Int())
		for j := 0; j < int(fieldLen); j++ {
			s = s + t.h
		}
		s = s + t.j
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

func (t tablify) printTitle(values interface{}, l reflect.Value) {
	title := t.v
	row := reflect.ValueOf(values)
	for i := 0; i < row.NumField(); i++ {

		n, p := fieldTitle(row.Type().Field(i))
		title = title + t.padForLength(p, int(l.FieldByName(n).Int())) + t.v
	}

	fmt.Println(title)
}

func (t tablify) printRow(value interface{}, l reflect.Value) {
	val := t.v
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
		val = val + t.padForLength(v, int(l.FieldByName(n).Int())) + t.v
	}

	fmt.Println(val)
}

func (t tablify) printRowsStruct(values []interface{}) {
	if !t.validate(values) {
		return
	}

	l := lengths(values)
	sep := t.createSeparator(l)

	// Start the table.
	fmt.Println(sep)

	vs := make([]interface{}, len(values))
	for i, v := range values {
		vs[i] = v
	}

	t.printTitle(vs[0], l)

	for _, v := range values {
		fmt.Println(sep)
		t.printRow(v, l)
	}

	// End the table.
	fmt.Println(sep)
}

func (t tablify) padForLength(value string, length int) string {
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

func (t tablify) validate(values []interface{}) bool {
	if len(values) == 0 {
		fmt.Println("Unable to render table. No data exists.")
		return false
	}

	return true
}
