FROM golang:1.22.1-alpine3.19

WORKDIR /gyan

COPY ./go.mod .
RUN GOPROXY=https://proxy.golang.org/cached-only go mod download

COPY . .
RUN go build 
RUN chmod 777 -R ./Gyan

EXPOSE 9090

CMD ["./Gyan"]