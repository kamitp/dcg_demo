# A hello world example with Go
FROM ubuntu:16.04
MAINTAINER Rabindra Nath Sharma <write.to.rabindra@gmail.com>
RUN apt-get update; apt-get -y upgrade;
RUN apt-get update; apt-get -y install git-core
# RUN apt-get update && apt-get install -y net-tools 
# RUN apt-get update && apt-get install -y iputils-ping
ADD https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz .
RUN tar -xvf go1.10.3.linux-amd64.tar.gz
RUN mv go /usr/local
RUN export PATH=$PATH:/usr/local/go/bin
RUN mkdir -p /gopath
ENV GOROOT=/usr/local/go
ENV GOPATH=/gopath
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH
RUN go get github.com/gorilla/mux
RUN go get github.com/dgrijalva/jwt-go
RUN go get golang.org/x/crypto/bcrypt
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
ADD . $GOPATH/src/dcg_demo/RecipeService
WORKDIR $GOPATH/src/dcg_demo/RecipeService
RUN cd $GOPATH/src/dcg_demo/RecipeService
#release build
RUN go build -ldflags "-w"
# enbale below line for debug build
# RUN go build
RUN go install 
CMD ["RecipeService"]
#CMD ["go", "run", "service.go"]
