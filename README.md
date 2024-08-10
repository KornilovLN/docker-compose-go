## Docker-compose with Go

**_https://github.com/PacktPublishing/A-Developer-s-Essential-Guide-to-Docker-Compose_**
```
https://go.dev/doc/install
```

Как использовать gin-gonic/gin

Вместо установки с помощью go install, нужно просто добавить этот пакет в зависимости проекта.
Вот как это сделать:

    Создайте или перейдите в свой проект:

    bash
```
mkdir myapp
cd myapp
```

Инициализируйте модуль Go:

bash
```
go mod init myapp
```

Это создаст файл go.mod, который будет управлять зависимостями вашего проекта.

Добавьте пакет gin-gonic/gin в зависимости:

bash
```
go get github.com/gin-gonic/gin@latest
```

Эта команда добавит gin в файл go.mod и скачает его в директорию вашего проекта.


Создайте файл main.go:
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

```
go run main.go
```


Чтобы убедиться, что сервер работает корректно, выполните запрос на определенный маршрут, например, GET /ping. Это можно сделать через браузер, Postman, или используя команду curl:

bash

curl http://localhost:8080/ping

Вы должны получить ответ:

json

{"message":"pong"}
