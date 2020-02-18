# REST API

This is a sample REST API that aims to explore the Go's HTTP interaction (paths/routes, methods, error handling, testing, status codes, etc). 

Notice that this example code is very restricted to the API interaction. I'm not covering here many important aspects such as code organization, graceful shutdown, use of middlewares for code re-use, JWT auth, service layer dependency, etc. - things that I want to cover in a next opportunity.


```bash
$ mkdir rest && cd rest && touch main.go
$ go mod init github.com/soeirosantos/pathtogo/web/rest
$ go get -u github.com/gorilla/mux
$ go get -u github.com/google/uuid
```

## Hot Reload
https://medium.com/@hotlinedring/effective-development-environment-with-go-ab74bd308c0f
https://www.stevenrombauts.be/2018/02/automatically-recompile-and-restart-your-go-application/

```bash
go get github.com/githubnemo/CompileDaemon
```

```bash
$ ~/go/bin/CompileDaemon -command='go test -v'
```
