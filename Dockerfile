FROM golang:1.20.6-alpine
WORKDIR /root
COPY . ./
RUN CGO_ENABLED=0 go build -o app .

FROM alpine:latest
VOLUME /root/config
VOLUME /arpcache

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /root/app ./
CMD ["./app"]
