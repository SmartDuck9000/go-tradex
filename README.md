# Тестовое задание для стажёра Backend в команду Trade Marketing.

Нужно разработать микросервис для счетчиков статистики. Сервис должен уметь взаимодействовать с клиентом при помощи REST API или JSON RPC запросов. Также нужно реализовать валидацию входных данных.

# Результат
- REST API
- Использован go в связке с gin и gorm
- PostgreSQL
- Валидация входных данных
- Возможность выбора поля, по которому будет осуществляться сортировка

# Запуск микросервиса

```
docker-compose up
```

Файл с запросами на создание юзера и таблицы - ```/src/db/init.sql```

Конфигурационный файл - ```src/.env```

## Метод сохранения статистики.
Принимает на вход:
- **date** - дата события
- **views** - количество показов
- **clicks** - количество кликов
- **cost** - стоимость кликов (в рублях с точностью до копеек)

Поля **views**, **clicks** и **cost** - опциональные.
Статистика агрегируется по дате.

Пример вызова:
```shell
curl \
-H 'Content-Type: application/json' \
--data '{
    "date":"2020-03-21",
    "views": 10,
    "clicks": 100,
    "cost": 60.55
}' \
http://0.0.0.0:5000/api/stat/save
```

Возможные ответы:
- Код 200 - статистика сохранена
- Код 400 - данные не прошли валидацию 

## Метод показа статистики
Принимает на вход:
- **from** - дата начала периода (включительно)
- **to** - дата окончания периода (включительно)

Отвечает статистикой, отсортированной по дате. В ответе должны быть поля:
- **date** - дата события
- **views** - количество показов
- **clicks** - количество кликов
- **cost** - стоимость кликов
- **cpc** = cost/clicks (средняя стоимость клика)
- **cpm** = cost/views * 1000 (средняя стоимость 1000 показов)

Пример вызова:
```shell
curl -X GET http://0.0.0.0:5000/api/stat/get?from=2020-03-16&to=2020-03-19&order_by=date
```

Возможные ответы:
- Код 200 - статистика получена, в теле json со статистикой отсортированной по выбранному полю или null, если нет сохраненной статистики
- Код 400 - параметры запроса не прошли валидацию

Пример тела ответа:
```json
[
  {
    "date": "2020-03-16",
    "views": 12,
    "clicks": 60,
    "cost": 240,
    "cpc": 4,
    "cpm": 20000
  },
  {
    "date": "2020-03-17",
    "views": 12,
    "clicks": 60,
    "cost": 180,
    "cpc": 3,
    "cpm": 15000
  },
  {
    "date": "2020-03-18",
    "views": 12,
    "clicks": 60,
    "cost": 0,
    "cpc": 0,
    "cpm": 0
  },
  {
    "date": "2020-03-19",
    "views": 0,
    "clicks": 0,
    "cost": 0,
    "cpc": 0,
    "cpm": 0
  }
]
```

## Метод сброса статистики
Удаляет всю сохраненную статистику.

Пример вызова:
```shell
curl -X DELETE http://0.0.0.0:5000/api/stat/delete
```

Возможные ответы:
- Код 200 - статистика удалена
- Код 400 - ошибка при удалении