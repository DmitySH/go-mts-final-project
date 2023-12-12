FROM golang:1.21 as builder

WORKDIR /app

COPY pkg ../pkg

COPY driver/go.mod .
COPY driver/go.sum .

RUN go mod download

COPY driver/cmd cmd
COPY driver/api api
COPY driver/pkg pkg
COPY driver/internal internal

ENV CGO_ENABLED=0
RUN go build -o driver ./cmd/driver/main.go

FROM scratch

COPY driver/configs configs
COPY driver/third_party third_party

COPY --from=builder /app/driver driver

CMD ["./driver"]
