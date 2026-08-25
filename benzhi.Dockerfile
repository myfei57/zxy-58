FROM golang:1.23
WORKDIR /app
ENV GOPROXY=off GOSUMDB=off CGO_ENABLED=0
COPY go.mod go.sum ./
COPY vendor ./vendor
COPY . .
RUN go build -mod=vendor -o /poolops ./cmd/poolops
ENV POOLOPS_ADDR=0.0.0.0:8080
EXPOSE 8080
CMD ["/poolops"]
