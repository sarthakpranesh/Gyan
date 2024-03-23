FROM golang:1.22.1-alpine3.19

WORKDIR /gyan

COPY . .

RUN go install
RUN go build 
RUN chmod 777 ./Gyan

EXPOSE 9090

CMD ["./Gyan"]