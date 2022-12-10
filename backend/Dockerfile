FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY . .


RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./cmd/main.go

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
RUN mkdir -p /usr/bin/images/profiles
RUN mkdir -p /usr/bin/images/movies
EXPOSE 80
ENTRYPOINT /go/bin/web-app --port 80