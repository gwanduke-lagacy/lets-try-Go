# 도커를 이용한 배포

## 개요
docker를 이용해 go 어플리케이션을 배포하는 방법을 실습합니다. 참고자료는 아래에 명시했으며, semaphore의 자료를 참고했지만 해당서비스를 사용해 배포하지 않습니다. AWS로 배포를 목표로 합니다.

## References
### [How To Deploy a Go Web Application with Docker](https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker)
- 목표
    - Docker에 대한 기본적인 이해
    - Go application을 개발하는데 Docker가 어떤 도움을 줄 수 있는지 안다
    - Go application을 production하기 위해 Docker container를 어떻게 만드는지 배운다
    - ~~Semaphore를 이용해 Docker container를 지속적으로 배포하는 방법을 안다~~ 
        - -> Semaphore를 사용하지 않기 때문에 이 부분은 다른방법으로 해결합니다.