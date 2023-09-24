package validate_subsequence

import (
	"bytes"
	"reflect"
	"unsafe"
)

func IsValidSubsequenceUnsafe(array []int, sequence []int) bool {
	la, ls := len(array), len(sequence)
	if la == 0 || ls == 0 || la < ls {
		return false
	}

	sz := int(unsafe.Sizeof(array[0]))
	a, s := a2b(array, sz), a2b(sequence, sz)
	_, _ = a[len(a)-1], s[len(s)-1]
	for {
		p := bytes.Index(a, s[:sz])
		if p == -1 {
			return false
		}
		a = a[p+sz:]
		if s = s[sz:]; len(s) == 0 {
			return true
		}
	}
}

func a2b(x []int, sz int) []byte {
	hx := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	h := reflect.SliceHeader{
		Data: hx.Data,
		Len:  hx.Len * sz,
		Cap:  hx.Cap * sz,
	}
	return *(*[]byte)(unsafe.Pointer(&h))
}
