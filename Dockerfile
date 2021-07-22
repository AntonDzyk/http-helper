FROM golang:1.16.4
WORKDIR /root/src/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o http-helper .

FROM scratch
COPY --from=0 /root/src/http-helper .
ENTRYPOINT ["/http-helper"]
