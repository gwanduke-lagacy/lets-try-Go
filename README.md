# Go Playground

## 개요
> Repo
- Author: letsget23@gmail.com (gwanduke)
- Language: Go
- Sub Language: HTML, CSS, JavaScript (jQuery)

> Saying...
~~Go-lang 으로 작성한 날씨 조회 시스템입니다.~~

처음 튜토리얼이 날씨조회 시스템이었는데,

현재는 플레이그라운드!

웹소켓 + API 통신으로 무언가 할것입니다~

## Usage
```
    $ go get letsget23/go-playground
    
    - 주제별 분류되는 폴더는 복수형
    - 실행가능한 루트 폴더는 단수
    
    # 프로젝트별 실행
    $ cd projects/weather-man
    $ go run *.go
    
    # 주제별 예제 실행
    $ cd websockets/echo
    $ go run *.go
```

## 무엇이 있나요?
- CSV Read/Write
- Log (Stdout, File, ...)
- HTTP 요청 (API call)
- HTTP Route, Simple static file server
- channel, go routine
- Websocket

- 위 내용을 이용한 Examples

## References
- 기초
    - [예제로 배우는 GO 프로그래밍](http://golang.site/)

- API, HTTP
    - [How i start Go](http://howistart.org/posts/go/1/index.html)
    - [HTTP Responses Snippets for Go](http://www.alexedwards.net/blog/golang-response-snippets)
    - [Making a RESTful JSON API in Go](https://thenewstack.io/make-a-restful-json-api-go/)

- Websocket
    - [Playing with websockets in Go](https://www.jonathan-petitcolas.com/2015/01/27/playing-with-websockets-in-go.html)
    - [websocket ping-pong](http://arlimus.github.io/articles/gin.and.gorilla/)

- Project 구성
    - [mingrammer/make-restful-api-with-go](https://speakerdeck.com/mingrammer/make-restful-api-with-go)