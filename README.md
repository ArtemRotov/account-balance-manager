# Микросервис-менеджер баланса пользователей
Пример микровервиса предназначенного для работы с балансом пользователей. Сервис предоставляет HTTP API и принимает/отдает запросы/ответы в формате JSON. 

Используемые технологии:
- PostgreSQL (в качестве хранилища данных)
- Docker (для запуска сервиса) -- Work in process
- Swagger (для документации API)
- gorilla/mux (веб фреймворк)
- slog (логгер)
- golang-migrate/migrate (для миграций БД)
- tesify, mockery (для тестирования) -- Work in process

**Основные функции:**
1. Метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить.
2. Метод получения баланса пользователя. Принимает id пользователя.
3. Метод резервирования средств с основного баланса на отдельном счете. Принимает id пользователя, ИД услуги, ИД заказа, стоимость.
4. Метод признания выручки – списывает из резерва деньги. Принимает id пользователя, ИД услуги, ИД заказа, сумму.
5. Метод отмены резервирования - возвращает деньги на баланс. Принимает id пользователя, ИД услуги, ИД заказа, сумму.

## Использование
Документацию после завпуска сервиса можно посмотреть по адресу `http://localhost:8080/swagger/index.html`
с портом 8080 по умолчанию

## Примеры использования
- [Регистрация](#sign-up)
- [Аутентификация](#sign-in)
- [Пополнение счёта](#account-deposit)
- [Получение баланса](#account-getBalance)
- [Резервирование средств](#reservation-create)
- [Признание выручки](#reservation-revenue)
- [Отмена резервирования](#reservation-refund)

### Регистрация <a name="sign-up"></a>

```curl
curl -X POST 'http://localhost:8080/auth/sign-up' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"example@mail.org",
    "password":"123456"
}'
```
Ответ:
```json
{
  "id": 1
}
```


### Аутентификация <a name="sign-in"></a>

```curl
curl -X POST 'http://localhost:8080/auth/sign-in' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"example@mail.org",
    "password":"123456"
}'
```
Ответ:
```json
{
  "Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODgxMjMzNzksImlhdCI6MTY4ODExNjE3OSwiVXNlcklkIjoxfQ.tydIAyeaaZVP01mv5iEXGbkNbw4OhBZZn1mjwFe0TM8"
}
```


### Пополнение счета <a name="account-deposit"></a>

```curl
curl -X POST 'http://localhost:8080/api/v1/account/deposit' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODgxMjMzNzksImlhdCI6MTY4ODExNjE3OSwiVXNlcklkIjoxfQ.tydIAyeaaZVP01mv5iEXGbkNbw4OhBZZn1mjwFe0TM8' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id":1,
    "amount":1000
}'
```
Ответ:
```json
{
    "user_id": 1,
    "balance": 1000
}
```


### Получение баланса <a name="account-getBalance"></a>

```curl
curl -X GET 'http://localhost:8080/api/v1/account/' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODgxMjMzNzksImlhdCI6MTY4ODExNjE3OSwiVXNlcklkIjoxfQ.tydIAyeaaZVP01mv5iEXGbkNbw4OhBZZn1mjwFe0TM8' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id":1
}'
```
Ответ:
```json
{
    "user_id": 1,
    "balance": 1000
}
```


### Резервирование средств <a name="reservation-create"></a>

```curl
curl -X POST 'http://localhost:8080/api/v1/reservation/create' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODgxMjMzNzksImlhdCI6MTY4ODExNjE3OSwiVXNlcklkIjoxfQ.tydIAyeaaZVP01mv5iEXGbkNbw4OhBZZn1mjwFe0TM8' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_id": 1,
    "amount": 500,
    "order_id": 1,
    "service_id": 1
}'
```
Ответ:
```json
{
    "id": 1,
}
```

### Признание выручки <a name="reservation-revenue"></a>

```curl
curl -X POST 'http://localhost:8080/api/v1/reservation/revenue' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODgxMjMzNzksImlhdCI6MTY4ODExNjE3OSwiVXNlcklkIjoxfQ.tydIAyeaaZVP01mv5iEXGbkNbw4OhBZZn1mjwFe0TM8' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_id": 1,
    "amount": 500,
    "order_id": 1,
    "service_id": 1
}'
```
Ответ:
```json
{
    "msg": "OK"
}
```

### Отмена резервирования <a name="reservation-refund"></a>

```curl
curl -X POST 'http://localhost:8080/api/v1/reservation/refund' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODgxMjMzNzksImlhdCI6MTY4ODExNjE3OSwiVXNlcklkIjoxfQ.tydIAyeaaZVP01mv5iEXGbkNbw4OhBZZn1mjwFe0TM8' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_id": 1,
    "amount": 500,
    "order_id": 1,
    "service_id": 1
}'
```
Ответ:
```json
{
    "msg": "OK"
}
```
