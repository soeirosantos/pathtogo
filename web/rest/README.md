# REST API

```bash
$ mkdir rest-api && cd rest-api && touch main.go
$ go mod init github.com/soeirosantos/pathtogo/web/rest
$ go get -u github.com/gorilla/mux
```

```go
package main

import "fmt"

func main() {
	fmt.Println("foo")
}
```

```bash
$ go build
```
