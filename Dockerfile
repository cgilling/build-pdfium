FROM centos:7
MAINTAINER "Chris Gilling" cgilling@gmail.com

RUN yum -y update; yum clean all;
RUN yum install -y \
    git \
    gcc-c++ \
    freetype-devel
RUN git clone https://chromium.googlesource.com/chromium/tools/depot_tools.git
ENV PATH /depot_tools:"$PATH"
ENV GYP_DEFINES pdf_enable_v8=0
RUN gclient config --unmanaged https://pdfium.googlesource.com/pdfium.git; \
    gclient sync; \
    cd pdfium; \
    build/gyp_pdfium; \
    ninja -C out/Release