FROM golang:latest as build-env

FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ADD capsule-http ./
