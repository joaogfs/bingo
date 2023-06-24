package parser

import (
	"fmt"
	"parser/buffer"
	"parser/bytecast"
	"reflect"
)

var buf buffer.Buffer

type Endian bool

const (
	Little Endian = true
	Big    Endian = false
)

type settings struct {
	Endianess Endian
}

var Settings = settings{Endianess: Little}

func Parse[T any](input []byte, ast *T) error {
	buf.Init(input)

	root_type := reflect.TypeOf(*ast)
	root_val := reflect.ValueOf(ast)

	fillStruct(buf, root_type, root_val)

	return nil
}

func fillStruct(buf buffer.Buffer, root_type reflect.Type, root_val reflect.Value) error {
	//pp.Printf("parser.fillStruct: root_type: %v root_val: %v\n", root_type.Name(), root_val)
	if root_type.Kind() != reflect.Struct {
		return fmt.Errorf("value of type %v is not a struct", root_type.Name())
	}

	for i := 0; i < root_type.NumField(); i++ {
		tl_index := []int{i}
		cur_field := root_type.FieldByIndex(tl_index)

		switch cur_field.Type.Kind() {
		case reflect.Pointer:
			if cur_field.Type.Elem().Kind() == reflect.Struct {

				fillStruct(buf, cur_field.Type.Elem(), root_val.Elem().FieldByIndex(tl_index))
			}
		case reflect.Array:
			arr_len := cur_field.Type.Len()
			elem_size := cur_field.Type.Elem().Size()

			for i := 0; i < arr_len; i++ {
				cur_bytes_val := buf.Next(int(elem_size))
				casted_val := bytecast.To(cur_field.Type.Elem().Kind(), cur_bytes_val)

				root_val.Elem().FieldByIndex(tl_index).Index(i).Set(reflect.ValueOf(casted_val))
			}

		default:
			cur_type_size := int(cur_field.Type.Size())
			cur_bytes_val := buf.Next(cur_type_size)
			if Settings.Endianess {
				cur_bytes_val = ReverseSlice(cur_bytes_val)
			}
			//fmt.Printf("parser.fillStruct: casting to %v\n", cur_field.Type.Kind())
			casted_val := bytecast.To(cur_field.Type.Kind(), cur_bytes_val)
			//pp.Printf("parser.fillStruct: is settable %v\n", root_val.Elem().Type().Kind())
			root_val.Elem().FieldByIndex(tl_index).Set(reflect.ValueOf(casted_val))
		}

		//fmt.Printf("parser.fillStruct: cur_field: %v %v\n", cur_field.Name, cur_field.Type.Name())
	}
	return nil
}
