# DockerGoProject

It is a 3-tier based architecture with dependency injection.

**Author**
  - *Andres Felipe Alfonso Ortiz*

**Technologies**
  - *Golang*: programming language.
  - *Mysql*: data persistence.
  - *Gin*: framework for rest applications.
  - *Mokery*: automatic mocks for unit tests.
  - *Dig*: automatic dependency injection.
  - *Docker*: application's contenerization

**Run unit tests**
  - execute tests
  ```
    export CONFIG_DIR=${workspaceRoot}/DockerGoProject/pkg/config && export SCOPE=local && go test -v ./... -covermode=atomic -coverprofile=coverage.out -coverpkg=./... -count=1
  ```
  - Look result in html
  ```
    go tool cover -html=coverage.out
  ```
**Gin**
  - Documentation
    - https://gin-gonic.com/docs/quickstart/

**Mokery**
  - Documentacion
    - https://vektra.github.io/mockery/installation/#homebrew
  - Instalacion mac
    ```
      brew install mockery
    ```
  - Crear mocks
    ```
      mockery --all --disable-version-string
    ```
**Dig**
  - Documentation
    - https://ruslan.rocks/posts/golang-dig
    - https://www.golanglearn.com/golang-tutorials/golang-dig-a-better-way-to-manage-dependency/

**Mysql**
  - Added DB with user configuration test_R and test_W.
  - The Gpool is removed since it causes it to be slower and GO already manages the automatic pool of connections.
    - https://koho.dev/understanding-go-and-databases-at-scale-connection-pooling-f301e56fa73

**Start Aplication**
  - Requeriments
  ```
    docker volume create cow_{scope}_vol
    docker network create cow_{scope}_network
  ```
  - Execute the next command for start the application.
  ```
    docker-compose up -d
  ```
**Config project**
  - For unit test
  ```
    "go.testEnvVars": {
          "CONFIG_DIR": "${workspaceRoot}/pkg/config",
          "SCOPE":"local"
      },
  ```
  - Environment vs-code
  ```
    "SCOPE": "local",
    "PORT": "8080",
    "CONFIG_DIR": "${workspaceRoot}/pkg/config",
    "GIN_MODE":"release",
    //config db
    "USER_DB":"",
    "PASSWORD_DB":"",
    "PORT_DB":"",
    "HOST_DB":"",
    "SCHEMA_DB":""
  ```

**Utils**
- https://www.youtube.com/watch?v=Ms5RKs8TNU4&t=1504s