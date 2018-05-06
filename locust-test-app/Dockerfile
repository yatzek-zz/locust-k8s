#FROM gcr.io/distroless/base
#
#COPY example_server /
#
#EXPOSE 8080
#
#ENTRYPOINT ["./example_server"]

FROM golang:1.10.1-alpine3.7 as builder
COPY example_server.go .
RUN go build -o /app example_server.go

FROM alpine:3.7
EXPOSE 8080
CMD ["./app"]
COPY --from=builder /app .