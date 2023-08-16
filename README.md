# todo-list
Тестовое задание на позицию Junior Go разработчика в компанию ТОО Region LLC. 

Микросервис для работы с задачами. Данный микросервис позволяет создавать, получать, обновлять и удалять задачи, а также позваляет помечать задачу как "выполненной".


Используемые технологии

- mongo-driver/mongo (для хранилища данных)
- docker и docker-compose (для запуска сервиса)
- swagger (для API документации)
- gin-gonic/gin (веб фреймворк)
- golang/mock, testify (для тестирования)

## Запуск

- Запустить сервис можно с помощью команды `sudo make compose-up`

Документацию находиться по адресу http://localhost:8080/swagger/index.html с портом 8080 по умолчанию


## Тестирование

Для запуска тестов необходимо выполнить команду `make test`


## Примеры запросов

- [Создание новой задачи](#create-task)
- [Обновление существующий задачи](#update-task)
- [Удаление задачи](#delete-task)
- [Выполнение задачи](#complete-task)
- [Получение списка задач по статусу](#get-tasks-by-status)


### Создание новой задачи <a name="create-task"></a>

Пример запроса

```
curl -X PUT "http://localhost:8080/api/todo-list/tasks" \
     -H "Content-Type: application/json" \
     -d '{
        "title": "Купить квартиру",
        "activeAt": "2023-08-10",
        "status" : "active"
     }'

```

Пример ответа

```
{
    "id": "64d74607f9c13a703bd73fcb",
    "code": 201
}
```


### Обновление существующий задачи <a name="update-task"></a>

Пример запроса

```
curl -X PUT "http://localhost:8080/api/todo-list/tasks/64d74607f9c13a703bd73fcb" \
     -H "Content-Type: application/json" \
     -d '{
        "title": "Купить квартиру в центре",
        "activeAt": "2023-08-10"
     }'

```

Примет ответа

```
{
    "title":"Купить квартиру в центре",
    "activeAt":"2023-08-10"
}
```

### Удаление задачи <a name="delete-task"></a>

Пример запроса 

```
curl -X DELETE "http://localhost:8080/api/todo-list/tasks/64d74607f9c13a703bd73fcb"
```

Пример ответа

```
{
    "status": "Successfully removed task"
}
```

### Выполнение задачи <a name="complete-task"></a>

Пример запроса

```
curl -X PUT "http://localhost:8080/api/todo-list/tasks/64d74607f9c13a703bd73fcb/done"
```

Пример ответа

```
{
    "message": "Task marked as done"
}
```

### Получение списка задач по статусу <a name="get-tasks-by-status"></a>

Пример запроса

```
curl -X GET "http://localhost:8080/api/todo-list/tasks?status=done"
```

Пример ответа

```
[
    {
        "id": "64d392d2b0c5710fbf9f74c1",
        "title": "Купить книгу - Над пропастью во ржи",
        "activeAt": "2023-08-10",
        "status": "done"
    },
    {
        "id": "64d46be11afc77c3d0e32850",
        "title": "Купить книгу - Высоконагруженные системы",
        "activeAt": "2023-08-10",
        "status": "done"
    },
    {
        "id": "64d74607f9c13a703bd73fcb",
        "title": "Купить квартиру в центре",
        "activeAt": "2023-08-10",
        "status": "done"
    }
]
```

Пример запроса 

```
curl -X GET http://localhost:8080/api/todo-list/tasks?status=active
```

Пример ответа 

```
[
    {
        "id": "64d46cb51afc77c3d0e32851",
        "title": "Купить машину",
        "activeAt": "2023-09-04",
        "status": "active"
    },
    {
        "id": "64da68b9932c443e11f89830",
        "title": "ВЫХОДНОЙ - Купить чехол",
        "activeAt": "2023-08-13",
        "status": "active"
    }
]
```
