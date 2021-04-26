FROM golang:1.16 AS builder

COPY . /go/src/app
WORKDIR /go/src/app

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/rest
RUN GOOS=linux go build -o converter.so -buildmode=c-shared ./cmd/sharedlib

RUN git log -1 --oneline > version.txt

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/app/app .
COPY --from=builder /go/src/app/converter.so .
COPY --from=builder /go/src/app/converter.h .
COPY --from=builder /go/src/app/cmd/sharedlib/converter.py .
COPY --from=builder /go/src/app/version.txt .

EXPOSE 8080

ENTRYPOINT ["./app"]