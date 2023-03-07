package libx265common

import (
	"sync"

	"github.com/ying32/dylib"
)

var libx265Dll *dylib.LazyDLL
var libx265DllOnce sync.Once

func GetLibx265Dll() (ans *dylib.LazyDLL) {
	libx265DllOnce.Do(func() {
		libx265Dll = dylib.NewLazyDLL(libx265Path)
	})
	ans = libx265Dll
	return
}

var libx265Path = "libx265.dll"

func SetLibx265Path(path0 string) {
	libx265Path = path0
}
