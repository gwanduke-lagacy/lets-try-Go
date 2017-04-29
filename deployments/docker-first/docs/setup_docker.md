# 도커 설정하기

- Author: letsget23@gmail.com (gwanduke)
- Created: 2017. 04. 30.

## `Dockerfile` 작성

> Dockerfile

```dockerfile
FROM golang:1.6

# Install beego and the bee dev tool
RUN go get github.com/astaxie/beego && go get github.com/beego/bee

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["bee", "run"]
```

## 이미지 설치 및 확인

> 명령

```terminal
docker build -t ma-image .

# ma-image, golang 이미지를 확인할 수 있을 것이다.
docker images
```

## `ma-image`의 컨테이너를 실행

```terminal
docker run -it --rm --name ma-instance -p 8080:8080 \
   -v $HOME/go/src/github.com/letsget23/go-playground/deployments/docker-first:/go/src/MathApp -w /go/src/MathApp ma-image
```

- `docker run`은 이미지로 부터 컨테이너를 실행하기 위해 사용한다.

- `-it` 플래그는 컨테이너를 쌍방향(interactive) 모드로 시작한다.

- `--rm` 플래그는 컨테이너가 종료된 후에 컨테이너를 clean out 한다.

- `--name ma-instance`는 컨테이너 이름을 지정한다.

- `-p 8080:8080` 플래그는 컨테이너가 8080 포트에서 접근 가능하도록 한다.

- `-v /app/MathApp:/go/src/MathApp`은 머신의 `/app/MathApp`을 컨테이너의 `/go/src/MathApp`에 매핑시킨다. 이는 development 파일들을 컨테이너 안/밖에서 이용가능하도록 한다.

- `ma-image` 는 컨테이너에서 사용할 이미지명이다.
