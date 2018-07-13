FROM golang:alpine

RUN mkdir /app 
ADD ./main /app
WORKDIR /app

EXPOSE 8080

CMD /app/main --query=$QUERY --user=$DB_USER --database=$DB_NAME
