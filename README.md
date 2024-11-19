## Requirements

Задача 1. Счетчик кликов.
Есть набор баннеров (от 10 до 100). У каждого есть ИД и название (id, name)

Нужно сделать сервис, который будет считать клики и собирать их в поминутную статистику (timestamp, bannerID, count)


Нужно сделать АПИ с двумя методами:

1. /counter/<bannerID> (GET)

Должен посчитать +1 клик по баннеру с заданным ИД



2. /stats/<bannerID> (POST)

Должен выдать статистику показов по баннеру за указанный промежуток времени (tsFrom, tsTo)



Язык: golang

СУБД: mongo или psql

Сложность:

- junior = кол-во запросов /counter 10-50 в секунду

- middle+ = кол-во запросов /counter 100-500 в секунду


## Launch

```
## Run the solution in docker-compose environment
make compose-run

## Other commands description
make help
```
