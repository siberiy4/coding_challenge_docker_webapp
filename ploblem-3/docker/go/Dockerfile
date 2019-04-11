FROM golang
WORKDIR /go/src/denki/go
ADD ./go .
RUN go get -u github.com/labstack/echo/... && go get -u  github.com/lib/pq && go get -u github.com/jinzhu/gorm
