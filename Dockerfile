FROM golang:latest
COPY . /opt/source
WORKDIR /opt/source/cmd/marketplace
RUN go build
ENTRYPOINT [ "/opt/source/cmd/marketplace/marketplace", "server", "--extensions-dir", "/mnt/marketplace-extensions" ]
