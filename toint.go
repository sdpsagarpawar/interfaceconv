package interfaceconv

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

func ToInt(i interface{}) int {
	if i == nil {
		return 0
	}
	elm := reflect.ValueOf(i)
	switch i.(type) {
	case uintptr, int, int8, int16, int64, int32, uint, uint8, uint16, uint32, uint64:
		if strings.Contains(elm.Type().String(), "uint") {
			return int(elm.Uint())
		}
		return int(elm.Int())
	case *uintptr, *int, *int8, *int16, *int64, *int32, *uint, *uint8, *uint16, *uint32, *uint64:
		if strings.Contains(elm.Type().String(), "uint") {
			return ToInt(elm.Elem().Uint())
		}
		return ToInt(elm.Elem().Int())
	case string:
		intVal, err := strconv.Atoi(elm.String())
		if err != nil {
			log.Panic(err)
		}
		return intVal
	case *string:
		return ToInt(elm.Elem().String())
	case float64, float32:
		return int(elm.Float())
	case *float64, *float32:
		return int(elm.Elem().Float())
	}
	return 0
}
