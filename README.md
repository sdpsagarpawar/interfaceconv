# Package interfaceconv

The zip package provides a fast and simple interface conversion.

- ToString
- ToInt
- ToFloat
- ToByte

## Features
- Simply call function to covert interface

## Tech
```
- golang >= 1.17
```
## Installation
```
go get -u github.com/sdpsagarpawar/interfaceconv
```
## Uses

```sh
a:="1"
intValue:=interfaceconv.ToInt(a)
a:=1
intValue:=interfaceconv.ToString(a)
a:=1
intValue:=interfaceconv.ToFloat(a)
a:="john doe"
intValue:=interfaceconv.ToByte(a)
```