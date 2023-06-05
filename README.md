# wb_lo
##### Проект предлагает сохранение и выдачу информации о заказах через веб-запросы с использованием postgres, nats-streaming и redis.

## Как установить?

1. **Установка и запуск**
```bash
docker compose up --build
```

2. **Восстановление базы данных (нужен установленный [go](https://go.dev/dl/))**
```bash
bash dump.sh -r base.sql
```

3. **Проверка работы**
```bash
curl http://localhost:3000/get-all-orders/
```

## Возможности

#### 1. Создание заказа через `form url encoded` с ключом `order` и `value` 

<details>
  <summary>Показать JSON</summary>

```json
{
  "order_uid": "39",
  "track_number": "track",
  "entry": "BIILLLO",
  "delivery": {
    "ID": 64,
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "ID": 44,
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "ID": 27,
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
```

</details>

URL: http://localhost:3000/create-order/

#### 2. Просмотр заказа по `id`
URL: http://localhost:3000/get-order/2

#### 3. Просмотр всех `id` заказов

URL: http://localhost:3000/get-all-orders

## Иерархия

- **backup**
    - **db** (Бекапы базы данных)
- **cmd**
  - **apiserver** (Entry point)
  - **other** (Dump утилита)
- **docker**
  - ... (Dockerfiles)
- **internal**
  - **apiserver** (Конфигурирование проекта)
  - **cache** (Кеш)
  - **messaging** (Брокер сообщений)
  - **model** (Описание моделей)
  - **storage** (Базы данных)

## Dump

#### Создание дампа и восстановление базы данных, с помощью команды bash dump.sh -key.<br> Ключи: 
- `-c` // создать дамп
- `-r filename.sql` // восстановить бд из файла
- `-d` // удалить данные в таблицах

#### Пример запросов:
- `bash dump.sh -c`
- `bash dump.sh -r dump_date_2023-05-28_23-17-28.sql`
- `bash dump.sh -d`

### Что люди говорят об этом проекте:
> 1 2 3 жопа. Евгений Алехин<br>

> Дорого я заплатил, дорого. Константин Сперанский<br>