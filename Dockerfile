FROM golang:1.18.8-alpine3.16 AS builder

RUN mkdir -p /go/src/quiz1
COPY ./ /go/src/quiz1
WORKDIR /go/src/quiz1
ENV  GO111MODULE=on
RUN cd /go/src/quiz1 && go build -o quiz -mod vendor

EXPOSE 80

FROM alpine:3.11.6
COPY --from=builder /go/src/quiz1/quiz /go/src/quiz1/config.yaml ./
ENTRYPOINT ["./quiz"]
