# 도커를 이용한 배포

- Author: letsget23@gmail.com (gwanduke)
- Created: 2017. 04. 30.

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

- 개발시 도커의 이점
    - 모든 팀 멤버가 표준 개발환경을 가지게 됨
    - 중앙에서 의존성을 업데이트하고 같은 컨테이너를 어떤 곳에서든 이용한다.
    - development와 production의 환경이 같다.
    - production에서만 나타날 수 있는 잠재적 문제를 해결

- 왜 도커를 Go 웹 어플리케이션에서 이용하는가?
    - 웹 어플리케이션은 전형적으로 템플릿과 설정 파일들을 가지고 있다. 도커는 이 파일들을 binary와 함께 동기화되도록 돕는다.
    - 도커는 development와 production에서 동일한 설정들을 가지도록 한다. 어플리케이션이 devleopment에서는 잘 작동하지만, production에서는 제대로 동작하지 않는 경우가 있다. 도커는 이런 문제들로 부터 자유롭게 해준다.
    - 머신, OS 그리고 설치된 소프트웨어는 다양할 수 있고 이는 큰팀에서 심각하다. 도커는 개발 셋업을 일관되게 해준다. 이는 팀이 더욱 생산적이고 개발중에 일어나는 마찰과 피할 수 있는 이슈들을 줄여준다.

- Development를 위해 도커 설정하기
    - devlopement를 위해 `Dockerfile`을 사용한다.
        - 파일들은 container의 안/밖에서 둘다 접근가능하다.
        - 개발중에 도커 컨테이너에서 live reload를 위해 `beego`의 `bee`를 이용한다.
        - 도커는 어플리케이션을 `8080`포트에 노출한다.
        - 머신에서 어플리케이션은 `/app/MathApp`에 위치한다.
        - 도커 컨테이너에서 어플리케이션은 `/go/src/MathApp`에 위치한다.
        - development를 위해 생성할 도커 이미지의 이름은 `ma-image`이다.
        - development 중에 실행할 도커 컨테이너 이름은 `ma-instance`이다.

- Production에서 도커 사용하기

    - 우선은 Semaphore를 이용해 도커 컨테이너의 Go 어플리케이션 배포를 수행한다.
        - git 저장소에 푸쉬된 후에 자동으로 빌드한다.
        - 자동으로 테스트를 수행한다.
        - 빌드가 성공적으로 수행되고 테스트가 패스하면, 도커 이미지를 생성한다.
        - 도커이미지를 Docker Hub로 푸쉬한다.
        - 서버를 최신 도커 이미지로 업데이트한다.

    - [Semaphore](https://semaphoreci.com/)로 자동으로 빌드와 테스트 수행하기
        - 의존성을 가져온다
        - 프로젝트를 빌드한다.
        - 테스트를 수행한다.
        - 만약 빌드와 테스트 둘중 하나라도 실패하면 프로세스는 정지되고 어떤것도 배포되지 않을 것이다.




