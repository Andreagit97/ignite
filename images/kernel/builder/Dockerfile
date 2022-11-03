FROM ubuntu:22.04 AS builder

ARG GCC_VERSION="gcc-7"

RUN echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu focal main universe" >> /etc/apt/sources.list

# Install dependencies
RUN apt-get update -y && \
    apt-get install -y --no-install-recommends \
	bc                    \
	bison                 \
	build-essential       \
	ccache                \
	flex                  \
	${GCC_VERSION}        \
	git                   \
	kmod                  \
	libelf-dev            \
	libncurses-dev        \
	libssl-dev            \
	wget                  \
	ca-certificates    && \
    update-alternatives --install /usr/bin/gcc gcc /usr/bin/${GCC_VERSION} 10

# Install crosscompilers for non-amd64 arches
# RUN apt-get install -y --no-install-recommends \
# 	binutils-multiarch \
# 	${GCC_VERSION}-aarch64-linux-gnu && \
# 	ln -s /usr/bin/aarch64-linux-gnu-${GCC_VERSION} /usr/bin/aarch64-linux-gnu-gcc

RUN apt update && apt install fakeroot rsync cpio python3 libdw-dev cmake liblz4-tool -y

RUN git clone  --branch v1.24 --single-branch https://github.com/acmel/dwarves.git
RUN cd dwarves && mkdir build && cd build && cmake -D__LIB=lib ../ && make install
RUN  ln -s /usr/local/lib/libdwarves.so.1 /usr/lib/
RUN  ln -s /usr/local/lib/libdwarves_reorganize.so.1 /usr/lib/
RUN  ln -s /usr/local/lib/libdwarves_emit.so.1 /usr/lib/
ENV  PATH "$PATH:/usr/local"
ENV  LD_LIBRARY_PATH "$LD_LIBRARY_PATH:/usr/local/lib/"