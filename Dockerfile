FROM golang:1.17-alpine
RUN apk update && apk add --no-cache build-base bash
RUN mkdir /go/src/hello
WORKDIR /go/src/hello

#Copy the source code
COPY . .

#ADD . /go/src/hello

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Build
RUN go build -o /hello

#CMD["/hello"]
