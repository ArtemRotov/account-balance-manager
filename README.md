# Микросервис-менеджер баланса пользователей
Пример микровервиса предназначенного для работы с балансом пользователей. Сервис предоставляет HTTP API и принимает/отдает запросы/ответы в формате JSON. 

Используемые технологии:
- PostgreSQL (в качестве хранилища данных)
- Docker (для запуска сервиса)
- Swagger (для документации API)
- gorilla/mux (веб фреймворк)
- golang-migrate/migrate (для миграций БД)

**Основные функции:**
1. Метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить.
2. Метод получения баланса пользователя. Принимает id пользователя.
3. Метод резервирования средств с основного баланса на отдельном счете. Принимает id пользователя, ИД услуги, ИД заказа, стоимость.
4. Метод признания выручки – списывает из резерва деньги. Принимает id пользователя, ИД услуги, ИД заказа, сумму.
5. Метод отмены резервирования - возвращает деньги на баланс. Принимает id пользователя, ИД услуги, ИД заказа, сумму.


## Примеры использования
- [Регистрация](#sign-up)
- [Аутентификация](#sign-in)

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
