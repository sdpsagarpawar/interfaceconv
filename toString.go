package interfaceconv

import (
	"fmt"
	"reflect"
	"strings"
)

func ToString(i interface{}) string {
	if i == nil {
		return ""
	}
	elm := reflect.ValueOf(i)
	switch v := i.(type) {
	case uintptr, int, int8, int16, int64, int32, uint, uint8, uint16, uint32, uint64:
		if strings.Contains(elm.Type().String(), "uint") {
			return fmt.Sprint(elm.Uint())
		}
		return fmt.Sprint(elm.Int())
	case string:
		return v
	case float64, float32:
		return fmt.Sprint(elm.Float())
	case []rune, []byte:
		return string(elm.Bytes())
	case *float64, *float32:
		return ToString(elm.Elem().Float())
	case *string:
		return elm.Elem().String()
	case *[]byte, *[]rune:
		return string(elm.Elem().String())
	case *uintptr, *int, *int8, *int16, *int64, *int32, *uint, *uint8, *uint16, *uint32, *uint64:
		if strings.Contains(elm.Elem().Type().String(), "uint") {
			return ToString(elm.Elem().Uint())
		}
		return ToString(elm.Elem().Int())
	}
	return ""
}
