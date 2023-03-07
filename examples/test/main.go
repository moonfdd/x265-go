package main

import (
	"encoding/json"
	"fmt"
	"os"
	"unsafe"

	"github.com/moonfdd/x265-go/libx265"
	"github.com/moonfdd/x265-go/libx265common"
)

func main() {
	os.Setenv("Path", os.Getenv("Path")+";./lib")
	libx265common.SetLibx265Path("./lib/libx265.dll")
	fmt.Println(libx265.AVC_INFO)
	a := libx265.X265ApiGet207(0)
	fmt.Println(a)
	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	fmt.Println(unsafe.Sizeof(libx265.X265Zone{}))
	fmt.Println(unsafe.Sizeof(libx265.X265AnalysisData{}))
}
