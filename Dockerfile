FROM golang:1.23-alpine as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN apk update
RUN apk add -U --no-cache ca-certificates && update-ca-certificates
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags="-w -s" -o cotify .
RUN ls -al

FROM scratch as go-runtime-container
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/cotify /cotify
EXPOSE 3000
ENTRYPOINT ["/cotify"]