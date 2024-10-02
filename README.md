## Docker-compose with Go

**_https://github.com/PacktPublishing/A-Developer-s-Essential-Guide-to-Docker-Compose_**
```
https://go.dev/doc/install
```

### Как использовать gin-gonic/gin
<br>Вместо установки с помощью go install, добавить этот пакет в зависимости проекта.
<br>Вот как это сделать:

**_Создайте или перейдите в свой проект:_**
<br>bash
```
mkdir myapp
cd myapp
```

**_Инициализируйте модуль Go:_**
<br>bash
```
go mod init myapp
```

Это создаст файл go.mod, который будет управлять зависимостями вашего проекта.

**_Добавьте пакет gin-gonic/gin в зависимости:_**
<br>bash
```
go get github.com/gin-gonic/gin@latest
```

Эта команда добавит gin в файл go.mod и скачает его в директорию вашего проекта.

**_Создайте файл main.go:_**
```
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // Запустит сервер на порту 8080
}
```

**_Запустите на исполнение:_**
```
go run main.go
```

**_Чтобы убедиться, что сервер работает корректно:_**
<br>выполните запрос на определенный маршрут, например, GET /ping.
<br>Это можно сделать через браузер, Postman, или используя команду curl:
<br>bash
```
curl http://localhost:8080/ping
```
Вы должны получить ответ:
<br>json
```
{"message":"pong"}
```
