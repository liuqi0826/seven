package gl

//#include <stdlib.h>
import "C"
import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func boolToInt(value bool) int8 {
	if value {
		return int8(1)
	}
	return int8(0)
}
func intToBool(value int8) bool {
	if value != 0 {
		return true
	}
	return false
}
func stringToUint8(str string) *uint8 {
	if !strings.HasSuffix(str, "\x00") {
		panic("str argument missing null terminator: " + str)
	}
	header := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return (*uint8)(unsafe.Pointer(header.Data))
}
func stringToChar(strs ...string) (cstrs **uint8, free func()) {
	if len(strs) == 0 {
		panic("Strs: expected at least 1 string")
	}

	for _, v := range strs {
		v += "\x00"
	}
	// Allocate a contiguous array large enough to hold all the strings' contents.
	n := 0
	for i := range strs {
		n += len(strs[i])
	}
	data := C.malloc(C.size_t(n))

	// Copy all the strings into data.
	dataSlice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(data),
		Len:  n,
		Cap:  n,
	}))
	css := make([]*uint8, len(strs)) // Populated with pointers to each string.
	offset := 0
	for i := range strs {
		copy(dataSlice[offset:offset+len(strs[i])], strs[i][:]) // Copy strs[i] into proper data location.
		css[i] = (*uint8)(unsafe.Pointer(&dataSlice[offset]))   // Set a pointer to it.
		offset += len(strs[i])
	}

	return (**uint8)(&css[0]), func() { C.free(data) }
}
func charToString(cstr *uint8) string {
	return C.GoString((*C.char)(unsafe.Pointer(cstr)))
}
func numberToPointer(value interface{}) unsafe.Pointer {
	if value == nil {
		return unsafe.Pointer(nil)
	}
	v := reflect.ValueOf(value)
	return unsafe.Pointer(v.UnsafeAddr())
}
func sliceToPointer(value interface{}) unsafe.Pointer {
	if value == nil {
		return unsafe.Pointer(nil)
	}
	v := reflect.ValueOf(value)
	return unsafe.Pointer(v.Index(0).UnsafeAddr())
}
func offsetPointer(ptr int32) unsafe.Pointer {
	return unsafe.Pointer(uintptr(ptr))
}
func ptr(data interface{}) unsafe.Pointer {
	if data == nil {
		return unsafe.Pointer(nil)
	}
	var addr unsafe.Pointer
	v := reflect.ValueOf(data)
	switch v.Type().Kind() {
	case reflect.Ptr:
		e := v.Elem()
		switch e.Kind() {
		case
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			addr = unsafe.Pointer(e.UnsafeAddr())
		default:
			panic(fmt.Errorf("unsupported pointer to type %s; must be a slice or pointer to a singular scalar value or the first element of an array or slice", e.Kind()))
		}
	case reflect.Uintptr:
		addr = unsafe.Pointer(v.Pointer())
	case reflect.Slice:
		addr = unsafe.Pointer(v.Index(0).UnsafeAddr())
	default:
		panic(fmt.Errorf("unsupported type %s; must be a slice or pointer to a singular scalar value or the first element of an array or slice", v.Type()))
	}
	return addr
}

type DebugProc func(
	source uint32,
	gltype uint32,
	id uint32,
	severity uint32,
	length int32,
	message string,
	userParam unsafe.Pointer)

var userDebugCallback DebugProc

//export debugCallback
func debugCallback(
	source uint32,
	gltype uint32,
	id uint32,
	severity uint32,
	length int32,
	message *uint8,
	userParam unsafe.Pointer) {
	if userDebugCallback != nil {
		userDebugCallback(source, gltype, id, severity, length, charToString(message), userParam)
	}
}
