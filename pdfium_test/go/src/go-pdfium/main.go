package main

/*
#cgo CFLAGS: -I/pdfium/public
#cgo LDFLAGS: -lm -L/pdfium/out/Release/obj -L/pdfium/out/Release/obj/third_party -lpdfium -lfxcodec -lfpdfdoc -lfpdfapi -lfpdftext -lformfiller -lfxcrt -lfxedit -lfxge -lfx_freetype -ljpeg -lfx_lcms2 -lfx_libopenjpeg -lfx_zlib -lbigint -lfx_agg -ljavascript -lpdfwindow -lfdrm
#cgo LDFLAGS: -lstdc++

#include <stdlib.h>
#include <fpdfview.h>
*/
import "C"
import "fmt"
import "unsafe"

func init() {
	C.FPDF_InitLibrary()
}

func main() {
     cs := C.CString("/pdfium_test/foxit_sdk.pdf")
     doc := C.FPDF_LoadDocument(cs, nil)
     C.free(unsafe.Pointer(cs))
     pageCount := C.FPDF_GetPageCount(doc)
     fmt.Printf("I ran, it got %d pages\n", pageCount)
}