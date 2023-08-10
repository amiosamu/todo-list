# todo-list
Тестовое задание на позицию Junior Go разработчика в компанию ТОО Region LLC. 

Микросервис для работы с задачами. Данный микросервис позволяет создавать, получать, обновлять и удалять задачи, а также позваляет помечать задачу как "выполненной".


Используемые технологии

- mongo-driver/mongo (для хранилища данных)
- docker и docker-compose (для запуска сервиса)
- swagger (для API документации)
- go-chi/chi (веб фреймворк)
- golang/mock, testify (для тестирования)

## Запуск

- Запустить сервис можно с помощью команды `make compose-up`

Документацию находиться по адресу http://localhost:8080/swagger/index.html с портом 8080 по умолчанию


## Тестирование

Для запуска тестов необходимо выполнить команду `make test`, для запуска тестов с покрытием `make cover`


## Примеры запросов

- [Создание новой задачи](#create-task)
- [Обновление существующий задачи](#update-task)
- [Удаление задачи](#delete-task)
- [Выполнение задачи](#complete-task)
- [Получениe списка задач по статусу](#get-tasks-by-status)


### Создание новой задачи <a name="create-task"></a>


### Обновление существующий задачи <a name="update-task"></a>


### Удаление задачи <a name="delete-task"></a>

### Выполнение задачи <a name="complete-task"></a>

### Получение списка задач по статусу <a name="get-tasks-by-status"></a>