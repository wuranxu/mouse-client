package scene

import (
	"reflect"
	"unsafe"
)

func ToString(data []byte) string {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&data))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	return *(*string)(unsafe.Pointer(&bh))
}

func ToBytes(data string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&data))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
