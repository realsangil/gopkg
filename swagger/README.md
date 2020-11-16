<h1 align="center">Welcome to  👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-v0.0.1-blue.svg?cacheSeconds=2592000" />
  <a href="https://github.com/realsangil/gopkg/blob/main/LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

> swagger를 쉽게 배포할 수 있도록 정의한 패키지

## Install

```sh
go get github.com/realsangil/gopkg/swagger
```

## Usage

### Go Standard HTTP package
```go
func main() {
  if err := http.ListenAndServe(":1202", HTTPHandleFunc("https://petstore.swagger.io/v2/swagger.json")); err != nil {
		panic(err)
	}
}
```

### Echo framework
```go
func main() {
  e := echo.New()
	e.GET("/docs", EchoHandleFunc("/v2/swagger.json", WithTitle("Example")))
	if err := e.Server.ListenAndServe(); err != nil {
		panic(err)
	}
}
```

## Run tests

```sh
go test ./...
```

## Author

👤 **realsangil**

* Website: https://blog.realsangil.net
* Github: [@realsangil](https://github.com/realsangil)

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2020 [realsangil](https://github.com/realsangil).<br />
This project is [MIT](https://github.com/realsangil/gopkg/blob/main/LICENSE) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_