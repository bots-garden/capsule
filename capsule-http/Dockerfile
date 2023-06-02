FROM golang:1.20-1-buster as build-env

FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ADD capsule-http ./
