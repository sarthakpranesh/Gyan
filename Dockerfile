FROM golang:1.16.15-alpine3.15

WORKDIR /gyan

COPY ./go.mod .
RUN GOPROXY=https://proxy.golang.org/cached-only go install

COPY . .
RUN go build 
RUN chmod 777 -R ./Gyan

EXPOSE 9090

CMD ["./Gyan"]