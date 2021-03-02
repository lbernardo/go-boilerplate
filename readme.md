# go-boilerplate
Go Boilerplate with Docker, HTTP Service, AWS Lambda and API Gateway with serverless.yml 

### Install
```shell
~$ git clone https://github.com/lbernardo/go-boilerplate

~$ make init pkg=github.com/MYUSER/MYPKG
```

### Enable AWS API Gateway
```shell
make get-aws
```

### Build
```shell
make build
```

### Docker composer
```yaml
version: "3"
services:
  my-service:
  container_name: my-service
  build: ./docker
  volumes:
    - .:/var/app
  dns: 8.8.8.8
  ports:
    - "84:80"
  environment:
    HOST: 0.0.0.0
    PORT: 8080
```

Every time a **.go** file is updated, it will build or code and restart or service automatically