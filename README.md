# avito_assignment

# Тестовое задание для стажера Backend 

### Run:
Запус сервиса:
- `go run main.go`

## API
### 1. Создание сегмента:
Название метода:  `/segment/create`

Входные параметры:

`name` - Название сегмента.

Выходные параметры:
`status` - Статус исполнения. `ok`/`ne ok`
`name` - Возвращаемое значение созданного сегмента
Пример запроса:
`curl localhost:8082/create_segment --data '{"name": "avito200"}'`
Результат:
`{
    "status": "ok",
    "name": "avito200"
}`

![image](https://github.com/nllllibeth/avito_assignment/assets/73400470/31896a59-51e2-4540-8444-29a65c0d0e4f)


### 2. Удаление сегмента:
Название метода:  `segment/delete`
Входные параметры:
`name` - Название сегмента.
Выходные параметры:
`status` - Статус исполнения. `ok`/`ne ok`
Пример запроса:
`curl localhost:8082/delete_segment --data '{"name": "avito200"}'`
Результат:
`{"status": "ok"}`

### 3. Добавление пользователя в сегмент:
Название метода:  `users/add_segments`
Входные параметры:
`segmentsToAdd` - Список сегментов к добавлению.
`segmentsToDelete` - Список сегментов к удалению.
`user_id` - ID пользователя.
Выходные параметры:
`status` - Статус исполнения. `ok`/`ne ok`
Пример запроса:
`curl localhost:8082/add_user --data '{"user_id": 0, "name": "avito200"}'`
Результат:
`{"status": "ok"}`

### 4. Получение активных сегментов пользователя:
Название метода:  `users/get_segments`
Входные параметры:
`user_id` - Id пользователя.
Выходные параметры:
`segments` - Сегменты к которым принадлежит пользователь.
Пример запроса:
`curl localhost:8082/get_user_segments --data '{"user_id": 0}'`
Результат:
`{"segments": ["avito200", "avito30"]}`

## Использованные технологии
- "github.com/go-chi/chi"
- "github.com/lib/pq"
- "log/slog"
- "net/http"
  
## Схема БД
1. Таблица users:
- `id - PK`
- `name`
2. Таблица segments:
- `id - PK`
- `name`
3. Таблица users_segments:
- `id - PK`
- `user_id`
- `segment_id`
