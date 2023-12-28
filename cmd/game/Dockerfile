FROM golang:1.20.4 AS builder
COPY . workdir
WORKDIR workdir

RUN ls -l

RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN CGO_ENABLED=0 go build -o /bin/game ./cmd/game/main.go

FROM scratch
COPY --from=builder /bin/game .
CMD ["./game","-agones"]