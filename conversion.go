package interfaceconv

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Converter interface defines the conversion methods.
type Converter interface {
	ToFloat() (float64, error)
	ToInt() (int, error)
	ToString() (string, error)
	ToByte() ([]byte, error)
}

// DefaultConverter struct implements the Converter interface.
type DefaultConverter struct {
	Value interface{}
}

// NewConverter creates a new Converter object with the given value.
func NewConverter(value interface{}) Converter {
	return &DefaultConverter{
		Value: value,
	}
}

// ToFloat converts the input value to a float64.
func (c *DefaultConverter) ToFloat() (float64, error) {
	return toFloat(c.Value)
}

// ToInt converts the input value to an int.
func (c *DefaultConverter) ToInt() (int, error) {
	return toInt(c.Value)
}

// ToString converts the input value to a string.
func (c *DefaultConverter) ToString() (string, error) {
	return toString(c.Value)
}

// ToByte converts the input value to a byte slice.
func (c *DefaultConverter) ToByte() ([]byte, error) {
	return toByte(c.Value)
}

func toFloat(i interface{}) (float64, error) {
	switch val := i.(type) {
	case uintptr, int, int8, int16, int64, int32, uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(val)
		if strings.Contains(v.Type().String(), "uint") {
			return float64(v.Uint()), nil
		}
		return float64(v.Int()), nil
	case *uintptr, *int, *int8, *int16, *int64, *int32, *uint, *uint8, *uint16, *uint32, *uint64:
		v := reflect.ValueOf(val).Elem()
		if strings.Contains(v.Type().String(), "uint") {
			return float64(v.Uint()), nil
		}
		return float64(v.Int()), nil
	case string:
		floatVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0, err
		}
		return floatVal, nil
	case *string:
		return toFloat(*val)
	case float64:
		return val, nil
	case *float64:
		return *val, nil
	case float32:
		return float64(val), nil
	case *float32:
		return float64(*val), nil
	}
	return 0, fmt.Errorf("unsupported type: %T", i)
}

func toInt(i interface{}) (int, error) {
	switch val := i.(type) {
	case uintptr, int, int8, int16, int64, int32, uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(val)
		if strings.Contains(v.Type().String(), "uint") {
			return int(v.Uint()), nil
		}
		return int(v.Int()), nil
	case *uintptr, *int, *int8, *int16, *int64, *int32, *uint, *uint8, *uint16, *uint32, *uint64:
		v := reflect.ValueOf(val).Elem()
		if strings.Contains(v.Type().String(), "uint") {
			return toInt(v.Uint())
		}
		return toInt(v.Int())
	case string:
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return 0, err
		}
		return intVal, nil
	case *string:
		return toInt(*val)
	case float64:
		return int(val), nil
	case *float64:
		return int(*val), nil
	case float32:
		return int(val), nil
	case *float32:
		return int(*val), nil
	}
	return 0, fmt.Errorf("unsupported type: %T", i)
}

func toString(i interface{}) (string, error) {
	switch val := i.(type) {
	case uintptr, int, int8, int16, int64, int32, uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(val)
		if strings.Contains(v.Type().String(), "uint") {
			return fmt.Sprint(v.Uint()), nil
		}
		return fmt.Sprint(v.Int()), nil
	case string:
		return val, nil
	case float64, float32:
		return fmt.Sprint(val), nil
	case []rune:
		return string(val), nil
	case []byte:
		return string(val), nil
	case *float64:
		return toString(*val)
	case *string:
		if val != nil {
			return *val, nil
		}
	case *[]byte:
		if val != nil {
			return string(*val), nil
		}
	case *[]rune:
		if val != nil {
			return string(*val), nil
		}
	case *uintptr, *int, *int8, *int16, *int64, *int32, *uint, *uint8, *uint16, *uint32, *uint64:
		v := reflect.ValueOf(val).Elem()
		if strings.Contains(v.Type().String(), "uint") {
			return toString(v.Uint())
		}
		return toString(v.Int())
	}
	return "", fmt.Errorf("unsupported type: %T", i)
}

func toByte(i interface{}) ([]byte, error) {
	str, err := toString(i)
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
