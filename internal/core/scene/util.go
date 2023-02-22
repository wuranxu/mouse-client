package scene

import "unsafe"

func ToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func ToBytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}
