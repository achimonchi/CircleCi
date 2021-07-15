FROM golang:1.15.11-alpine3.13 AS builder


RUN mkdir /app
ADD . /app

WORKDIR /app

# ENTRYPOINT [ "go", "test", "-v", "controller/**" ]
RUN go build -o main .
# RUN go test controller/**



# add alpine
FROM alpine

WORKDIR /app

COPY --from=builder /app/ /app/

CMD ["/app/main"]
EXPOSE 4000
