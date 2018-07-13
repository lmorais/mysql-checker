## Sidecar container to check Mysql service


### Dependencies:

	- Mysql server running
	- Docker


### Compile and build Docker image:

```
> CGO_ENABLED=0 go build -o main

> docker build . -t mysql-checker

```

### Run app:

```
> sudo docker run --rm -it -d -p 8080:8080 --net=host mysql-checker

```