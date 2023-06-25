FROM gitpod/workspace-base:latest

# ------------------------------------
# Install Go
# ------------------------------------
ENV GO_VERSION=1.20

ENV GOPATH=$HOME/go-packages
ENV GOROOT=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH
RUN curl -fsSL https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz | tar xzs \
    && printf '%s\n' 'export GOPATH=/workspace/go' \
                      'export PATH=$GOPATH/bin:$PATH' > $HOME/.bashrc.d/300-go

# ------------------------------------
# Install TinyGo
# ------------------------------------
ARG TINYGO_VERSION="0.28.1"
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo_${TINYGO_VERSION}_amd64.deb
RUN sudo dpkg -i tinygo_${TINYGO_VERSION}_amd64.deb
RUN rm tinygo_${TINYGO_VERSION}_amd64.deb

RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN go install -v golang.org/x/tools/gopls@latest
