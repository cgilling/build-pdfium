
default: pdfium_test

pdfium_test: main.cc
	g++ -o pdfium_test -I/pdfium/public -L/pdfium/out/Release/obj -L/pdfium/out/Release/obj/third_party main.cc -pthread -lpdfium -lfxcodec -lfpdfdoc -lfpdfapi -lfpdftext -lformfiller -lfxcrt -lfxedit -lfxge -lfx_freetype -ljpeg -lfx_lcms2 -lfx_libopenjpeg -lfx_zlib -lbigint -lfx_agg -ljavascript -lpdfwindow -lfdrm
