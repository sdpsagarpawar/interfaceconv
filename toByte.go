package interfaceconv

func ToByte(i interface{}) []byte {
	return []byte(ToString(i))
}
