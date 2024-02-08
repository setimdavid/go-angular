FROM golang:1.13.1-alpine AS GO_BUILD
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.10
WORKDIR app
COPY --from=GO_BUILD /go/bin/server ./
RUN ls
CMD ./server
