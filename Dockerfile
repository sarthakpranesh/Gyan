FROM golang:1.16.15-alpine3.15

WORKDIR /gyan

COPY . .
RUN GOPROXY=https://proxy.golang.org/cached-only go install

RUN go build 
RUN chmod 777 -R ./Gyan

EXPOSE 9091

CMD ["./Gyan"]