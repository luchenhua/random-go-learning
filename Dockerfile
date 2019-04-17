#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk add --no-cache git
COPY . .
RUN go mod tidy
RUN go mod verify
RUN go mod vendor
RUN go install -mod=vendor -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/random-learning-go /random-learning-go
ENV LIEN_API_ADDRESS=https://ual-uat-lien.decathlon.com.cn/trigger-msg/v1/api/event/query
ENTRYPOINT ./random-learning-go
LABEL Name=random-learning-go Version=0.0.1
EXPOSE 3000
