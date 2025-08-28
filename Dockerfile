
FROM ubuntu:20.04 as build

# Disable Prompt During Packages Installation
ARG DEBIAN_FRONTEND=noninteractive

ARG TAG=main

# Env variables
ENV PATH="${PATH}:/usr/local/go/bin"
ENV LD_LIBRARY_PATH="${LD_LIBRARY_PATH}:/usr/lib64/"

COPY . .

# Install loxilb related packages
RUN mkdir -p /opt/loxilb && \
    mkdir -p /root/loxilb-io/loxiflow/ && \
    # Update Ubuntu Software repository
    apt-get update && apt-get install -y wget && \
    arch=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) && echo $arch && if [ "$arch" = "arm64" ] ; then apt-get install -y gcc-multilib-arm-linux-gnueabihf; else apt-get update && apt-get install -y  gcc-multilib;fi && \
    # Arch specific packages - GoLang
    wget https://go.dev/dl/go1.23.0.linux-${arch}.tar.gz && tar -xzf go1.23.0.linux-${arch}.tar.gz --directory /usr/local/ && rm go1.23.0.linux-${arch}.tar.gz && \
    # Dev and util packages
    apt-get install -y clang llvm libelf-dev libpcap-dev vim net-tools ca-certificates \
    elfutils dwarves git libbsd-dev bridge-utils wget unzip build-essential \
    bison flex sudo iproute2 pkg-config tcpdump iputils-ping curl bash-completion && \
    # Install openssl-3.3.1
    wget https://github.com/openssl/openssl/releases/download/openssl-3.3.1/openssl-3.3.1.tar.gz && tar -xvzf openssl-3.3.1.tar.gz && \
    cd openssl-3.3.1 && ./Configure enable-ktls '-Wl,-rpath,$(LIBRPATH)' --prefix=/usr/local/build && \
    make -j$(nproc) && make install_dev install_modules && cd - && \
    cp -a /usr/local/build/include/openssl /usr/include/ && \
    if [ -d /usr/local/build/lib64  ] ; then mv /usr/local/build/lib64  /usr/local/build/lib; fi && \
    cp -fr /usr/local/build/lib/* /usr/lib/ && ldconfig && \
    rm -fr openssl-3.3.1*  && \
    ./build.sh && \
    cp loxiflow /root/loxilb-io/loxiflow/.

FROM ubuntu:20.04

# LABEL about the loxiflow image
LABEL name="loxiflow" \
      vendor="loxilb.io" \
      version="0.1" \
      release="0.1" \
      summary="loxiflow docker image" \
      description="Loxilb flow exporter" \
      maintainer="inhogog2@netlox.io"

# Disable Prompt During Packages Installation
ARG DEBIAN_FRONTEND=noninteractive

# Env variables
ENV PATH="${PATH}:/usr/local/go/bin"
ENV LD_LIBRARY_PATH="${LD_LIBRARY_PATH}:/usr/lib64/"

RUN apt-get update && apt-get install -y --no-install-recommends sudo wget \
    libbsd-dev iproute2 tcpdump bridge-utils net-tools libllvm10 ca-certificates && \
    wget https://raw.githubusercontent.com/loxilb-io/tools/refs/heads/main/k8s/mkllb-url.sh && \
    chmod +x mkllb-url.sh && mv mkllb-url.sh /usr/local/sbin/mkllb-url && \
    rm -rf /var/lib/apt/lists/* && apt clean

COPY --from=build /usr/lib64/libbpf* /usr/lib64/
COPY --from=build /usr/local/build/lib/* /usr/lib64
COPY --from=build /usr/local/go/bin /usr/local/go/bin
COPY --from=build /opt/loxilb /opt/loxilb
COPY --from=build /loxiflow /root/loxilb-io/loxiflow/loxiflow

ENTRYPOINT ["/root/loxilb-io/loxiflow/loxiflow"]

# Expose Ports
EXPOSE 11111 22222 3784

