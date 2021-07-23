# advita-comics-backend
Серверная часть advita-comics

## Доступные ручки:
```
GET /company - возвращает информацию по активной компании.
 Активная компания может быть только одна, иначе вернется ошибка

 ответ:
 {
    "terminationAmount": 30000,
    "collectedAmount": 4856,
    "dayRemains": 161
 }
 
 POST /donation - создает новое пожертвование
 
 ожидаемый запрос:
 {
    "areRegularPaymentsEnabled": true,
    "comicsId": 1,
    "donationAmount": 332,
    "isSubscribedToGetReport": true,
    "isSubscribedToTrackProgress": true,
    "userEmail": "testov@gmail.com",
    "userName": "Test",
    "personalisation": {
        "previewName": "eeeehow",
        "costumeColor": "blue",
        "characterGender": 0,
        "characterName": "superman"
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