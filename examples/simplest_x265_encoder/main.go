// https://github.com/leixiaohua1020/simplest_encoder/blob/master/simplest_x265_encoder/simplest_x265_encoder.cpp
package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/x265-go/libx265"
	"github.com/moonfdd/x265-go/libx265common"
)

func main0() ffcommon.FInt {

	var i, j ffcommon.FInt
	var ret ffcommon.FInt
	var y_size ffcommon.FInt
	var buff []byte

	//FILE* fp_src  = fopen("../cuc_ieschool_640x360_yuv444p.yuv", "rb");
	fp_src, _ := os.Open("./resources/cuc_ieschool_640x360_yuv420p.yuv")
	fp_dst_file := "./out/cuc_ieschool_640x360_yuv420p.h265"
	fp_dst, _ := os.Create(fp_dst_file)

	var pNals *libx265.X265Nal
	var iNal ffcommon.FUint32T = 0

	var pParam *libx265.X265Param
	var pHandle *libx265.X265Encoder
	var pPic_in *libx265.X265Picture

	//Encode 50 frame
	//if set 0, encode all frame
	var frame_num ffcommon.FInt = 0
	var csp ffcommon.FInt = libx265.X265_CSP_I420
	var width, height ffcommon.FInt = 640, 360

	//Check
	if fp_src == nil || fp_dst == nil {
		fmt.Printf("Error open files.\n")
		return -1
	}

	pParam = libx265.X265ParamAlloc()
	pParam.X265ParamDefault()
	pParam.BRepeatHeaders = 1 //write sps,pps before keyframe
	pParam.InternalCsp = csp
	pParam.SourceWidth = width
	pParam.SourceHeight = height
	pParam.FpsNum = 25
	pParam.FpsDenom = 1
	//Init
	pHandle = pParam.X265EncoderOpen207()
	if pHandle == nil {
		fmt.Printf("x265_encoder_open err\n")
		return 0
	}
	y_size = pParam.SourceWidth * pParam.SourceHeight

	pPic_in = libx265.X265PictureAlloc()
	pParam.X265PictureInit(pPic_in)
	if frame_num == 0 {
		fi, _ := fp_src.Stat()
		switch csp {
		case libx265.X265_CSP_I444:
			buff = make([]byte, y_size*3)
			frame_num = int32(fi.Size()) / (y_size * 3)
			pPic_in.Planes[0] = uintptr(unsafe.Pointer(&buff[0]))
			pPic_in.Planes[1] = uintptr(unsafe.Pointer(&buff[0])) + uintptr(y_size)
			pPic_in.Planes[2] = uintptr(unsafe.Pointer(&buff[0])) + uintptr(y_size*2)
			pPic_in.Stride[0] = width
			pPic_in.Stride[1] = width
			pPic_in.Stride[2] = width
		case libx265.X265_CSP_I420:
			buff = make([]byte, y_size*3/2)
			frame_num = int32(fi.Size()) / (y_size * 3 / 2)
			pPic_in.Planes[0] = uintptr(unsafe.Pointer(&buff[0]))
			pPic_in.Planes[1] = uintptr(unsafe.Pointer(&buff[0])) + uintptr(y_size)
			pPic_in.Planes[2] = uintptr(unsafe.Pointer(&buff[0])) + uintptr(y_size*5/4)
			pPic_in.Stride[0] = width
			pPic_in.Stride[1] = width / 2
			pPic_in.Stride[2] = width / 2
		default:
			fmt.Printf("Colorspace Not Support.\n")
			return -1
		}
	}

	//Loop to Encode
	for i = 0; i < frame_num; i++ {
		switch csp {
		case libx265.X265_CSP_I444:

			fp_src.Read(ffcommon.ByteSliceFromByteP((*byte)(unsafe.Pointer(pPic_in.Planes[0])), int(y_size))) //Y
			fp_src.Read(ffcommon.ByteSliceFromByteP((*byte)(unsafe.Pointer(pPic_in.Planes[1])), int(y_size))) //U
			fp_src.Read(ffcommon.ByteSliceFromByteP((*byte)(unsafe.Pointer(pPic_in.Planes[2])), int(y_size))) //V

		case libx265.X265_CSP_I420:

			fp_src.Read(ffcommon.ByteSliceFromByteP((*byte)(unsafe.Pointer(pPic_in.Planes[0])), int(y_size)))   //Y
			fp_src.Read(ffcommon.ByteSliceFromByteP((*byte)(unsafe.Pointer(pPic_in.Planes[1])), int(y_size/4))) //U
			fp_src.Read(ffcommon.ByteSliceFromByteP((*byte)(unsafe.Pointer(pPic_in.Planes[2])), int(y_size/4))) //V

		default:

			fmt.Printf("Colorspace Not Support.\n")
			return -1

		}

		ret = pHandle.X265EncoderEncode(&pNals, &iNal, pPic_in, nil)
		fmt.Printf("Succeed encode frame: %5d\n", i)

		for j = 0; j < int32(iNal); j++ {
			a := unsafe.Sizeof(libx265.X265Nal{})
			pNal := (*libx265.X265Nal)(unsafe.Pointer(uintptr(unsafe.Pointer(pNals)) + uintptr(a*uintptr(j))))
			fp_dst.Write(ffcommon.ByteSliceFromByteP(pNal.Payload, int(pNal.SizeBytes)))
		}
	}
	//flush encoder
	for {
		ret = pHandle.X265EncoderEncode(&pNals, &iNal, nil, nil)
		if ret == 0 {
			break
		}
		fmt.Printf("Flush 1 frame.\n")
		for j = 0; j < int32(iNal); j++ {
			a := unsafe.Sizeof(libx265.X265Nal{})
			pNal := (*libx265.X265Nal)(unsafe.Pointer(uintptr(unsafe.Pointer(pNals)) + uintptr(a*uintptr(j))))
			fp_dst.Write(ffcommon.ByteSliceFromByteP(pNal.Payload, int(pNal.SizeBytes)))
		}
		i++
	}
	pHandle.X265EncoderClose()
	pPic_in.X265PictureFree()
	pParam.X265ParamFree()

	fp_src.Close()
	fp_dst.Close()

	fmt.Printf("\nffplay %s\n", fp_dst_file)

	return 0
}

func main() {
	os.Setenv("Path", os.Getenv("Path")+";./lib")
	libx265common.SetLibx265Path("./lib/libx265.dll")
	main0()
}
