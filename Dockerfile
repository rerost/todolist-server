FROM golang:1.11.1

## Base package
RUN apt-get update && apt-get install -y unzip 

## Install
RUN curl -Lo grapi https://github.com/izumin5210/grapi/releases/download/v0.2.2/grapi_linux_amd64 && chmod +x grapi && mv grapi /usr/local/bin

ENV PROTOC_ZIP=protoc-3.3.0-linux-x86_64.zip 
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.3.0/$PROTOC_ZIP && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && rm -f $PROTOC_ZIP

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN go get github.com/izumin5210/gex/cmd/gex

## Depend on src
ENV REPO_NAME=rerost/todolist-server
ENV SRC_DIR=${GOPATH}/src/github.com/${REPO_NAME}/
ENV PATH=${PATH}:${GOPATH}/bin
RUN mkdir -p ${SRC_DIR}
WORKDIR ${SRC_DIR}
ADD . ./

RUN dep ensure
RUN gex --build
RUN grapi build

CMD ["./bin/server"]
