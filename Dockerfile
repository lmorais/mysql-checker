FROM golang:alpine

RUN mkdir /app 
ADD ./main /app/ 
WORKDIR /app

EXPOSE 8080

CMD ["/app/main"]
