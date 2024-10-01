package imgui

// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

type StringPool struct {
	data     []byte
	spill    int
	cleanups []*C.char
}

func (sp *StringPool) Reset() {
	for _, ptr := range sp.cleanups {
		C.free(unsafe.Pointer(ptr))
	}
	sp.cleanups = sp.cleanups[:0]

	if sp.spill > 0 {
		sp.data = make([]byte, 0, cap(sp.data)+sp.spill)
		sp.spill = 0
	} else {
		sp.data = sp.data[:0]
	}
}

func (sp *StringPool) StoreCString(input string) *C.char {
	length := len(input)

	if length+1 > cap(sp.data)-len(sp.data) {
		sp.spill += length + 1
		ptr := C.CString(input)
		sp.cleanups = append(sp.cleanups, ptr)
		return ptr
	}

	at := len(sp.data)
	sp.data = append(sp.data, input...)
	sp.data = append(sp.data, 0)
	s := sp.data[at : at+length+1]

	return (*C.char)(unsafe.Pointer(&s[0]))
}
