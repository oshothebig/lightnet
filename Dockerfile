FROM golang:1.16.6 as builder

WORKDIR /build
COPY . .
RUN make build


FROM scratch as exporter

COPY --from=builder /build/bin/* /
