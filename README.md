# Планирование досуговых мероприятий (Сервер)

API позволяющее получать рекомендации, фильтровать и искать необходимые пользователю мероприятия.

## Первый запуск проекта

Проект развертывается в Docker контейнерах — основное приложение и БД PostgreSQL. Для запуска/остановки проекта используйте команды make.

### Сборка контейнера и запуск
```
make app-setup-and-up
```
### Запуск проекта
```
make app-up
```

### Подключиться в контейнер приложения
```
make app-bash
```

### Подключиться в контейнер БД
```
make db-bash
```

### Статус миграции goose
```
make db-migrate-status
```

### Выполнить миграцию goose
```
make db-migrate-up
```

### Откат миграции goose
```
make db-migrate-down
```

## API Запросы

### Изменить шаблон пользователя

Запрос:
```http request
http://localhost:4000/register [POST]
```

```json
{
  "gender": "female",
  "age": "23",
  "hobbies": [
    {
      "label": "Спорт",
      "value": "Спорт"
    },
    {
      "label": "Концерты",
      "value": "Концерты"
    }
  ]
}
```

Ответ:

```json
{
  "CreatedAt": "2021-06-13T06:02:24.633288Z",
  "UpdatedAt": "2021-06-14T04:45:27.921220095Z",
  "DeletedAt": null,
  "ID": 4,
  "Username": "",
  "Email": "",
  "Gender": "female",
  "Age": "23",
  "Events": null,
  "Preferences": [
    {
      "CreatedAt": "2021-06-13T23:18:32.643233Z",
      "UpdatedAt": "2021-06-13T23:18:32.643233Z",
      "DeletedAt": null,
      "ID": 9,
      "Title": "Спорт",
      "Users": null
    },
    {
      "CreatedAt": "2021-06-13T05:11:48.506574Z",
      "UpdatedAt": "2021-06-13T05:11:48.506574Z",
      "DeletedAt": null,
      "ID": 3,
      "Title": "Концерты",
      "Users": null
    }
  ],
  "Likes": null
}
```

### Добавить мероприятие в избранное

Запрос:
```http request
http://localhost:4000/wishlist [POST]
```

```json
{
  "id": 98016257,
  "wishlist": true
}
```

Где ID - id мероприятия, wishlist - флаг лайка.

Ответ:

```json
[
  {
    "id": 98016257,
    "lang": "ru",
    "title": "«Прометей прикованный» в культурном центре «Новослободский»",
    "text": "<p>Молодые актеры театра-студии &laquo;Три ступени&raquo; покажут спектакль &laquo;Прометей прикованный&raquo; по пьесе древнегреческого драматурга Эсхила. Дарья Кабанова, режиссер постановки, предлагает зрителям взглянуть на поступки Прометея с непривычной точки зрения.</p>",
    "date_from": "2021-06-29 14:00:00",
    "date_to": "2021-06-29 21:30:00",
    "date_from_timestamp": 1624964400,
    "date_to_timestamp": 1624991400,
    "date": "2021-06-29 14:00:00",
    "date_timestamp": 1623417263,
    "created_at": "2021-06-11 16:14:23",
    "created_at_timestamp": 1623417263,
    "updated_at": "2021-06-14 02:02:04",
    "updated_at_timestamp": 1623625324,
    "published_at": "2021-06-11 17:04:00",
    "status": "public",
    "search": 1,
    "ya_rss": 0,
    "free": 0,
    "is_powered": 0,
    "label": null,
    "oiv_id": 12585090,
    "restriction": {
      "age": 12
    },
    "has_image": 1,
    "active_from": null,
    "active_to": null,
    "active_from_timestamp": null,
    "active_to_timestamp": null,
    "icon_id": null,
    "kind": {
      "id": "afisha",
      "title": "События афиши",
      "type": 48
    },
    "is_oiv_publication": 0,
    "ebs_id": 7126,
    "ebs_type": "event",
    "ebs_title": "Прометей прикованный ",
    "ebs_agent_uid": "museum201",
    "image": {
      "id": 3172801281,
      "title": null,
      "copyright": null,
      "original": {
        "id": 3172801281,
        "title": null,
        "src": "/upload/newsfeed/events/imagesWRpb1x(3).jpg",
        "width": 1766,
        "height": 968,
        "type": "original"
      },
      "small": {
        "id": 3172998281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(4).jpg",
        "width": 480,
        "height": 263,
        "type": "small"
      },
      "middle": {
        "id": 3172999281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(6).jpg",
        "width": 800,
        "height": 439,
        "type": "middle"
      },
      "big": {
        "id": 3173000281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(1).jpg",
        "width": 1600,
        "height": 877,
        "type": "big"
      },
      "download": {
        "id": 3173001281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(10).jpg",
        "width": 1766,
        "height": 968,
        "type": "download"
      },
      "thumb": {
        "id": 3173002281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(9).jpg",
        "width": 400,
        "height": 219,
        "type": "thumb"
      },
      "show": {
        "id": 3173003281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(12).jpg",
        "width": 1200,
        "height": 658,
        "type": "show"
      },
      "1x1_m": {
        "id": 3173004281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(7).jpg",
        "width": 800,
        "height": 800,
        "type": "1x1_m"
      },
      "1x1_s": {
        "id": 3173005281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(16).jpg",
        "width": 480,
        "height": 480,
        "type": "1x1_s"
      },
      "2x1_b": {
        "id": 3173006281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(3).jpg",
        "width": 1600,
        "height": 800,
        "type": "2x1_b"
      },
      "2x1_m": {
        "id": 3173007281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(2).jpg",
        "width": 800,
        "height": 400,
        "type": "2x1_m"
      },
      "2x1_s": {
        "id": 3173008281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(15).jpg",
        "width": 480,
        "height": 240,
        "type": "2x1_s"
      },
      "3x1_b": {
        "id": 3173009281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(8).jpg",
        "width": 1599,
        "height": 533,
        "type": "3x1_b"
      },
      "3x1_m": {
        "id": 3173010281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(14).jpg",
        "width": 798,
        "height": 266,
        "type": "3x1_m"
      },
      "3x1_s": {
        "id": 3173011281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(13).jpg",
        "width": 480,
        "height": 160,
        "type": "3x1_s"
      },
      "4x1_b": {
        "id": 3173012281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(11).jpg",
        "width": 1600,
        "height": 400,
        "type": "4x1_b"
      },
      "4x1_m": {
        "id": 3173013281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3)(5).jpg",
        "width": 800,
        "height": 200,
        "type": "4x1_m"
      },
      "4x1_s": {
        "id": 3173014281,
        "title": null,
        "src": "/upload/newsfeed/afisha/imagesWRpb1x(3).jpg",
        "width": 480,
        "height": 120,
        "type": "4x1_s"
      }
    },
    "organizations": [
      12585090
    ],
    "spheres": [
      {
        "id": 80299,
        "title": "Спектакли",
        "special": 0,
        "activated": 1,
        "priority": 0
      }
    ]
  }
]
```

### Добавить мероприятие в избранное

Запрос:
```http request
http://localhost:4000/wishlist [GET]
```

```json
Без параметров
```

Ответ:

```json
Аналогичный ответ...
```

### Получить рекомендации для пользователя

Запрос:
```http request
http://localhost:4000/recommendations [GET]
```

```json
Без параметров
```

Ответ:

```json
{
  "items": [
    {
      "id": 97623257,
      "lang": "ru",
      "title": "«Ожерелье Финноугории» в Московском культурном фольклорном центре под руководством Людмилы Рюминой",
      "text": "<p>В Московском культурном фольклорном центре под руководством Людмилы Рюминой состоится концертная программа <a title=\"\" href=\"https://www.rumina.ru/performances/ozherele-finnougorii/\" target=\"_blank\" rel=\"noopener\">&laquo;Ожерелье Финноугории&raquo;</a>. &nbsp;Мероприятие посвящено 100-летию Республики Коми.</p>\n<p>На сцене выступит Государственный ансамбль песни и танца Республики Коми имени Виктора Морозова &laquo;Асъя кыа&raquo;. Он исполнит яркие сценические номера мордовской, удмуртской, марийской культур&nbsp; и культуры коми.</p>\n<p>Концерт пройдет в рамках федерального проекта &laquo;Мы &mdash; Россия&raquo;.</p>",
      "date_from": "2021-06-21 19:00:00",
      "date_to": "2021-06-21 20:30:00",
      "date_from_timestamp": 1624291200,
      "date_to_timestamp": 1624296600,
      "date": "2021-06-21 19:00:00",
      "date_timestamp": 1623336829,
      "created_at": "2021-06-10 17:53:49",
      "created_at_timestamp": 1623336829,
      "updated_at": "2021-06-11 10:52:31",
      "updated_at_timestamp": 1623397951,
      "published_at": "2021-06-11 10:53:00",
      "status": "public",
      "search": 0,
      "ya_rss": 0,
      "free": 0,
      "is_powered": 0,
      "label": "",
      "oiv_id": 12585090,
      "restriction": {
        "age": 6
      },
      "has_image": 1,
      "active_from": null,
      "active_to": null,
      "active_from_timestamp": null,
      "active_to_timestamp": null,
      "icon_id": null,
      "kind": {
        "id": "afisha",
        "title": "События афиши",
        "type": 48
      },
      "is_oiv_publication": 0,
      "ebs_id": 0,
      "ebs_type": "",
      "ebs_title": "",
      "ebs_agent_uid": "",
      ...
      

```

### Отправить сообщение в чат мероприятия

Запрос:

```http request
http://localhost:4000/comment [POST]
```

```json
{
  "EventID": 6,
  "Message": "Привет!",
  "MessageID": 8
}
```

Ответ:

```json
{
  "CreatedAt": "2021-06-13T21:59:58.471266Z",
  "UpdatedAt": "2021-06-13T22:07:57.866826366Z",
  "DeletedAt": null,
  "ID": 1,
  "EventID": 6,
  "Messages": [
    {
      "CreatedAt": "2021-06-13T22:07:57.867176859Z",
      "UpdatedAt": "2021-06-13T22:07:57.867176859Z",
      "DeletedAt": null,
      "ID": 5,
      "UserID": 4,
      "ChatID": 1,
      "Message": "Привет!"
    }
  ]
}
```


