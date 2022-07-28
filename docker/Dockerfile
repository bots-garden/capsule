FROM gitpod/workspace-full

#RUN sudo apt-get update && \
#    sudo apt-get install gettext libncurses5 libxkbcommon0 libtinfo5 -y

# ------------------------------------
# Install Couchbase
# ------------------------------------
RUN curl -O https://packages.couchbase.com/releases/couchbase-release/couchbase-release-1.0-amd64.deb && \
    sudo dpkg -i ./couchbase-release-1.0-amd64.deb && \
    sudo apt-get update && \
    sudo apt-get install couchbase-server-community -y && \
    rm ./couchbase-release-1.0-amd64.deb

USER gitpod

RUN brew install httpie && \
    brew install bat && \
    brew install exa && \
    brew install hey && \
    brew install redis



# ------------------------------------
# Install TinyGo
# ------------------------------------
ARG TINYGO_VERSION="0.24.0"
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo_${TINYGO_VERSION}_amd64.deb
RUN sudo dpkg -i tinygo_${TINYGO_VERSION}_amd64.deb
RUN rm tinygo_${TINYGO_VERSION}_amd64.deb
