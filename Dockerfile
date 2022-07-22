FROM golang:1.17.8 AS builder
WORKDIR /
COPY . .
RUN make

FROM ubuntu:20.04

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list

RUN apt-get update -y && \
    apt-get install iptables -y

WORKDIR /

COPY --from=builder /iptables-server .

ENTRYPOINT ["/iptables-server"]


