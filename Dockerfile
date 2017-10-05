FROM golang:alpine
LABEL version=0.0.2 author="Alexis Daciuk <adaciuk@gmail.com"

RUN apk add --update\
    git 

WORKDIR /go/src/app
RUN go-wrapper download golang.org/x/tools/cmd/present
RUN go-wrapper install golang.org/x/tools/cmd/present


EXPOSE 3999

ENTRYPOINT /go/bin/present -http="0.0.0.0:3999" -orighost="127.0.0.1"