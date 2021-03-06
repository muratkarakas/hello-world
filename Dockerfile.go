FROM golang:1.14
 
WORKDIR /go/src/myapp
COPY main.go .
 
# build
RUN CGO_ENABLED=0 go build -o main .
 
FROM alpine:latest 
RUN apk --no-cache add ca-certificates
 
WORKDIR /root/
 
COPY --from=0 /go/src/myapp/main .
 
CMD ["./main"]
EXPOSE 6060