interfaceconv
============

The `interfaceconv` package provides a set of functions and an interface for converting values between different types.

Installation
------------

To install the `interfaceconv` package, use the following command:

```shell
go get github.com/sdpsagarpawar/interfaceconv
```

## Usage
Import the interfaceconv package into your Go code:

```
import "github.com/sdpsagarpawar/interfaceconv"
```
1. Create a new Converter object using the NewConverter function and pass in the value you want to convert:
```
value := 42
converter := interfaceconv.NewConverter(value)
```
2. Use the methods provided by the Converter interface to perform the desired type conversion:
```
floatVal, err := converter.ToFloat()
if err != nil {
    // handle error
}
```

## Supported Conversion Methods
The Converter interface provides the following conversion methods:

ToFloat() (float64, error): Converts the value to a float64.
ToInt() (int, error): Converts the value to an int.
ToString() (string, error): Converts the value to a string.
ToByte() ([]byte, error): Converts the value to a []byte.

## Examples
Here are some examples of how to use the interfaceconv package:
```
package main

import (
    "fmt"
    "github.com/sdpsagarpawar/interfaceconv"
)

func main() {
    // Convert an integer to a float64
    value := 42
    converter := interfaceconv.NewConverter(value)
    floatVal, err := converter.ToFloat()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Float:", floatVal)
    }

    // Convert a string to an integer
    str := "123"
    converter = interfaceconv.NewConverter(str)
    intVal, err := converter.ToInt()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Int:", intVal)
    }

    // Convert a float32 to a string
    float32Val := float32(3.14)
    converter = interfaceconv.NewConverter(float32Val)
    strVal, err := converter.ToString()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("String:", strVal)
    }

    // Convert a byte slice to a string
    bytes := []byte{97, 98, 99}
    converter = interfaceconv.NewConverter(bytes)
    strVal, err = converter.ToString()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("String:", strVal)
    }
}

```

## License
This project is licensed under the MIT License. See the LICENSE file for more information.
```
Feel free to customize the README file according to your project's needs and update the import path with your actual username and repository.
```
