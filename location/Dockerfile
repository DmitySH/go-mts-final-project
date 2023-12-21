FROM golang:1.21 as builder

WORKDIR /app

COPY pkg ../pkg

COPY location/go.mod .
COPY location/go.sum .

RUN go mod download

COPY location/cmd cmd
COPY location/api api
COPY location/pkg pkg
COPY location/env.dev env.dev
COPY location/internal internal

ENV CGO_ENABLED=0
RUN go build -o location ./cmd/location/main.go

FROM scratch

COPY location/configs configs
COPY location/third_party third_party

COPY --from=builder /app/location location

CMD ["./location"]
