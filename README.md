# advita-comics-backend
Серверная часть advita-comics

## Запуск
```
git clone https://github.com/advita-comics/advita-comics-backend.git
make config
docker-compose up
```
Приложение будет доступно на порту :4040

## Доступные ручки:
```
GET /company - возвращает информацию по активной компании.
 Активная компания может быть только одна, иначе вернется ошибка

 ответ:
 {
    "terminationAmount": 30000,
    "collectedAmount": 4856,
    "dayRemains": 161,
    "donationCount": 10
 }
 
 POST /donation - создает новое пожертвование
 
 ожидаемый запрос:
{
  "character": {
    "gender": "0", 
    "name": "myname",
    "costumeColor": "blue",
    "hairColor": "red"
  },
  "donation": {
    "amount": 700,
    "directionId": 3,
    "userEmail": "bad@gmail.com",
    "areRegularPaymentsEnabled": true
  },
  "subscriptions": {
    "getReport": true,
    "trackProgress": true
  }
}

пример удачного ответа:
{
    "message": "Пожертвование создано"
}

пример ошибки:
{
    "error": "'comicsId' не передан"
}

```