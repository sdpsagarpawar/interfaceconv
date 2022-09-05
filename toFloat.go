package interfaceconv

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

func ToFloat(i interface{}) float64 {
	if i == nil {
		return 0
	}
	elm := reflect.ValueOf(i)
	switch i.(type) {
	case uintptr, int, int8, int16, int64, int32, uint, uint8, uint16, uint32, uint64:
		if strings.Contains(elm.Type().String(), "uint") {
			return float64(elm.Uint())
		}
		return float64(elm.Int())
	case *uintptr, *int, *int8, *int16, *int64, *int32, *uint, *uint8, *uint16, *uint32, *uint64:
		if strings.Contains(elm.Type().String(), "uint") {
			return float64(elm.Elem().Uint())
		}
		return float64(elm.Elem().Int())
	case string:
		floatVal, err := strconv.ParseFloat(elm.String(), 10)
		if err != nil {
			log.Panic(err)
		}
		return floatVal
	case *string:
		return ToFloat(elm.Elem().String())
	case float64, float32:
		return elm.Float()
	case *float64, *float32:
		return elm.Elem().Float()
	}
	return 0
}
