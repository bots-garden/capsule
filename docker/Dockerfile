FROM gitpod/workspace-full

#RUN sudo apt-get update && \
#    sudo apt-get install gettext libncurses5 libxkbcommon0 libtinfo5 -y

USER gitpod

RUN brew install httpie && \
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
