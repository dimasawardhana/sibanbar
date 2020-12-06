FROM golang:latest AS go_builder
ADD .
WORKDIR .
RUN go mod download
