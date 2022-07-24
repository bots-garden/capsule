FROM gitpod/workspace-full

RUN sudo apt-get update && \
    sudo apt-get install gettext libncurses5 libxkbcommon0 libtinfo5 -y

USER gitpod

RUN brew tap suborbital/subo && \
    brew install subo && \
    brew install httpie && \
    brew install bat && \
    brew install exa && \
    brew install hey

# ------------------------------------
# Install TinyGo
# ------------------------------------
ARG TINYGO_VERSION="0.24.0"
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo_${TINYGO_VERSION}_amd64.deb
RUN sudo dpkg -i tinygo_${TINYGO_VERSION}_amd64.deb
RUN rm tinygo_${TINYGO_VERSION}_amd64.deb

# ------------------------------------
# Install Sat (Suborbital)
# ------------------------------------
RUN git clone --depth=1 https://github.com/suborbital/sat.git && \
    cd sat && \
    go build -o .bin/sat -tags netgo,wasmtime . && \
    sudo cp .bin/sat /usr/local/bin/sat && \
    cd .. && \
    rm -rf sat

# ------------------------------------
# Install Wagi (Deislab)
# ------------------------------------
RUN mkdir tmp-wagi && \
    cd tmp-wagi && \
    wget https://github.com/deislabs/wagi/releases/download/v0.8.1/wagi-v0.8.1-linux-amd64.tar.gz && \
    tar -zxf wagi-v0.8.1-linux-amd64.tar.gz && \
    sudo cp wagi /usr/local/bin/wagi && \
    cd .. && \
    rm -rf tmp-wagi

# ------------------------------------
# Install Spin (Fermyon)
# ------------------------------------
RUN wget https://github.com/fermyon/spin/releases/download/v0.2.0/spin-v0.2.0-linux-amd64.tar.gz && \
    tar xfv spin-v0.2.0-linux-amd64.tar.gz && \
    sudo cp spin /usr/local/bin/spin && \
    rm spin; rm readme.md; rm LICENSE; rm spin-v0.2.0-linux-amd64.tar.gz

# ------------------------------------
# Install Wasi Runtimes
# ------------------------------------
RUN curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.10.0 && \
    curl https://get.wasmer.io -sSfL | sh && \
    curl https://wasmtime.dev/install.sh -sSf | bash
