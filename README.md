# rssagg

Маленькое приложение на Go по мотивам курса [freeCodeCamp.org: Go Programming –
Golang Course with Bonus Projects
(YouTube)](https://www.youtube.com/watch?v=un6ZyFkqFKo)

- Предоставляет REST API
- Обеспечивается аутентификация пользователей по API ключам
- Пользователи могут добавлять свои RSS и подписываться на RSS ленты
- Воркер в фоне загружает фид из интернета в БД

## Окружение

В качестве окружения нужна лишь СУБД PostgreSQL. В корне репозитория есть
[docker-compose.yml](./docker-compose.yml), можно поднять БД в Docker.

## Сборка и запуск приложения в Docker Compose

**Сборка**

```sh
docker compose build rssagg
```

**Запуск**

```sh
docker compose up -d rssagg
```
