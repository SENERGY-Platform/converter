FROM golang:1.22 AS builder

COPY . /go/src/app
WORKDIR /go/src/app

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/rest
RUN ./build_so.sh

RUN git log -1 --oneline > version.txt

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/app/app .
COPY --from=builder /go/src/app/converter.so .
COPY --from=builder /go/src/app/converter.h .
COPY --from=builder /go/src/app/converter_arm.so .
COPY --from=builder /go/src/app/converter_arm.h .
COPY --from=builder /go/src/app/version.txt .

EXPOSE 8080

ENTRYPOINT ["./app"]