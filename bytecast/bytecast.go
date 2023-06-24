package bytecast

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
)

/*
* int uint
* float
* bool
* string
 */

func To(val_type reflect.Kind, b []byte) any {
	switch val_type {
	case reflect.Bool:
		return b[0] != 0
	case reflect.Int8:
		return int8(b[0])
	case reflect.Int16:
		return int16(binary.BigEndian.Uint16(b))
	case reflect.Int32:
		return int32(binary.BigEndian.Uint32(b))
	case reflect.Int64:
		return int64(binary.BigEndian.Uint64(b))
	case reflect.Uint8:
		return uint8(b[0])
	case reflect.Uint16:
		return binary.BigEndian.Uint16(b)
	case reflect.Uint32:
		return binary.BigEndian.Uint32(b)
	case reflect.Uint64:
		return binary.BigEndian.Uint64(b)
	case reflect.Float32:
		return math.Float32frombits(binary.BigEndian.Uint32(b))
	case reflect.Float64:
		return math.Float64frombits(binary.BigEndian.Uint64(b))
	case reflect.String:
		return string(b)
	default:
		panic(fmt.Errorf("bytecast: invalid type %v", val_type.String()))
	}
}
