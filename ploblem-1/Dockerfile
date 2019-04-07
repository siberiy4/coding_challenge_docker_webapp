FROM golang AS build-env
ADD . /app
WORKDIR /app
RUN go get -u github.com/labstack/echo/... && CGO_ENABLED=0 go build -o main main.go

FROM alpine
COPY --from=build-env /app/main /usr/local/bin/main
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/main"]