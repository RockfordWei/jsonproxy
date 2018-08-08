FROM ubuntu:16.04

ENV GOPATH=/home
WORKDIR /home
RUN apt-get -y update && apt-get install -y golang git
RUN rm -rf /home/bin /home/src /home/pkg
RUN mkdir /home/bin && mkdir /home/pkg && mkdir /home/src && mkdir /home/src/jsonproxy
RUN go get github.com/gorilla/mux
COPY ./jsonproxy.go /home/src/jsonproxy
RUN cd /home/src/jsonproxy && go build && go install && cd /home

