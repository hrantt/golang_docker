## Description

A small go program that prints a greeting message with a randomly picked name out of 3 available options.

To run it go to src and type the following.
```sh
go run greetings.go
```
Running unit tests. 
```sh
go test
```



## Docker

Just a small container that downloads go-alpine, compiles and runs greetings program. 

```sh
docker build --tag go-greetings .
```

This will create the dillinger image and pull in the necessary dependencies.

Once done, run the Docker image.
```sh
docker run  go-greetings 
```


## License

MIT

