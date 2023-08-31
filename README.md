# avito_assignment

# Тестовое задание для стажера Backend 

### Run:
Запус сервиса:
- `go run main.go`

## API
### 1. Создание сегмента:
Название метода:  `create_segment`
Входные параметры:
`name` - Название сегмента.
Выходные параметры:
`status` - Статус исполнения. `ok`/`ne ok`
Пример запроса:
`curl localhost:8082/create_segment --data '{"name": "avito200"}'`
Результат:
`{"status": "ok"}`

### 2. Удаление сегмента:
Название метода:  `delete_segment`
Входные параметры:
`name` - Название сегмента.
Выходные параметры:
`status` - Статус исполнения. `ok`/`ne ok`
Пример запроса:
`curl localhost:8082/delete_segment --data '{"name": "avito200"}'`
Результат:
`{"status": "ok"}`

### 3. Добавление пользователя в сегмент:
Название метода:  `add_user`
Входные параметры:
`user_id` - Id пользователя.
`name` - Название сегмента.
Выходные параметры:
`status` - Статус исполнения. `ok`/`ne ok`
Пример запроса:
`curl localhost:8082/add_user --data '{"user_id": 0, "name": "avito200"}'`
Результат:
`{"status": "ok"}`

### 4. Получение активных сегментов пользователя:
Название метода:  `get_user_segments`
Входные параметры:
`user_id` - Id пользователя.
Выходные параметры:
`segments` - Сегменты к которым принадлежит пользователь.
Пример запроса:
`curl localhost:8082/get_user_segments --data '{"user_id": 0}'`
Результат:
`{"segments": ["avito200", "avito30"]}`

## Использованные технологии
- Технология 1. для того-то тогото
- Технология 2. для того-то тогото
## Схема БД
1. Таблица name1:
- `id - PK`
- `name`
2. Таблица name2:
- `id - PK`
- `name`
