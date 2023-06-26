# Микросервис-менеджер баланса пользователей

**Задача:**

Необходимо реализовать микросервис для работы с балансом пользователей (зачисление средств, списание средств, перевод средств от пользователя к пользователю, а также метод получения баланса пользователя). Сервис должен предоставлять HTTP API и принимать/отдавать запросы/ответы в формате JSON. 

**Требования к сервису:**

1. Сервис должен предоставлять HTTP API с форматом JSON как при отправке запроса, так и при получении результата.
2. Реляционная СУБД: MySQL или PostgreSQL.
3. Использование docker и docker-compose для поднятия и развертывания dev-среды.
4. Разработка интерфейса в браузере НЕ ТРЕБУЕТСЯ. Взаимодействие с API предполагается посредством запросов из кода другого сервиса. Для тестирования можно использовать любой удобный инструмент. Например: в терминале через curl или Postman.

**Основное:**

1. Метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить.
2. Метод получения баланса пользователя. Принимает id пользователя.
3. Метод резервирования средств с основного баланса на отдельном счете. Принимает id пользователя, ИД услуги, ИД заказа, стоимость.
4. Метод признания выручки – списывает из резерва деньги. Принимает id пользователя, ИД услуги, ИД заказа, сумму.
5. Метод отмены резервирования - возвращает деньги на баланс. Принимает id пользователя, ИД услуги, ИД заказа, сумму.


**Будет плюсом:**

1. Покрытие кода тестами.
2. [Swagger](https://swagger.io/solutions/api-design/) файл для вашего API.
3. Реализовать сценарий разрезервирования денег, если услугу применить не удалось.


