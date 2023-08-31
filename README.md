
## Запуск

	docker-compose up --build

## Общее
#### В случае некорректного ввода все методы возвращают ответ (400)
```json
{
  "type": "error",
  "message": "bad input"
}
```

## Работа с сегментами

> http://localhost:8080/segment

### Добавление сегмента (POST)
> http://localhost:8080/segment/add
#### Пример запроса
	curl -X POST http://localhost:8080/segment/add -d '{"slug":"ONE"}'
#### Ответ в случае успеха (201)
```json
{
  "message":"success",
  "type":"success"
}
```
#### Если сегмент с переданным slug существует в БД, ответ (409)
```json
{
  "type": "error",
  "message": "this segment is exists"
}
```
### Удаление сегмента (POST)
> http://localhost:8080/segment/delete
#### Пример запроса
	curl -X POST http://localhost:8080/segment/delete -d '{"slug":"ONE"}'
#### Ответ в случае успеха (201)
```json
{
  "message":"success",
  "type":"success"
}
```

#### Если сегмента с переданным slug нет в БД, ответ (404)
```json
{
  "type": "error",
  "message": "there is no segment with this slug"
}
```

## Работа с пользователями и их сегментами
> http://localhost:8080/users_segments

### Изменение пользовательских сегментов (добавление/удаление) (POST)
> http://localhost:8080/users_segments/change
#### Пример тела запроса
```json
{
  "slugs_to_add": [
    "AVITO_DISCOUNT_30"
  ],
  "slugs_to_delete": [
    "AVITO_DISCOUNT_50"
  ],
  "user_id": 1000
}
```
#### Пример запроса
	curl -X POST http://localhost:8080/users_segments/add -d '{"slugs_to_add":["ONE", "THREE"], "slugs_to_delete": ["TWO"], "user_id":10100}'

> Сегменты, которых нет в БД, игнорируются, если пользователя нет в БД, он создается

#### Ответ в случае успеха (201)
```json
{
  "message":"success",
  "type":"success"
}
```
#### Ответ, если существующие в БД сегменты в массиве для добавления встретились в списке для удаления (409)
```json
{
  "type": "error",
  "message": "segments to add are found in segments to delete"
}
```
### Получение пользовательских сегментов (POST)
	http://localhost:8080/users_segments/get
#### Пример запроса
	curl -X POST http://localhost:8080/users_segments/get -d '{"user_id":1001}'
#### Ответ в случае успеха (200)
```json
{
  "type":"success",
  "user_id":1001,
  "slugs": ["TWO"]
}
```
#### Ответ, если у пользователя с переданным user_id нет сегментов (204)
```json
{
  "type": "info",
  "message": "this user has no segments"
}
```
#### Ответ, если пользователя с переданным user_id нет в БД (404)
```json
{
  "type": "error",
  "message": "this user is not exists"
}
```



