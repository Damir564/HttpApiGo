# Решение тестового задания

<a>https://github.com/avito-tech/backend-trainee-assignment-2023</a>

## Запуск сервиса

Для запуска сервиса необходимо выполнить следующие шаги:

1. Установить Docker и Docker Compose, если они еще не установлены на вашем компьютере.
2. Склонировать репозиторий:

bash
git clone https://github.com/Damir564/HttpApiGo.git


3. Перейти в директорию проекта.


4. Установить необходимые значение в .env файле.


5. Запустить сервис с помощью Docker Compose:

bash
docker compose up --build


6. Проверить, что сервис успешно запустился:

bash
docker-compose ps


Вы должны увидеть следующий вывод:


    NAME                IMAGE           COMMAND                  SERVICE    CREATED             STATUS              PORTS
    ---------------------------------------------------------------------------------------------------------------------------------------
    httpapigo-db-1      postgres        "docker-entrypoint.s…"   db         14 seconds ago      Up 13 seconds       0.0.0.0:5432->5432/tcp       
    httpapigo-web-1     httpapigo-web   "air main.go -b 0.0.…"   web        14 seconds ago      Up 13 seconds       0.0.0.0:8080->8080/tcp  


Сервис успешно запущен и доступен по адресу http://localhost:8080.

## Swagger

Swagger-спецификация всего API доступна по адресу http://localhost:8080/swagger/index.html.

## Список выполненных заданий

**Основное задание (минимум):**

1. <mark style="background-color: grey">Метод создания сегмента. Принимает slug (название) сегмента.</mark> 
        
        URL: /segment

        Метод: POST

        Тело запроса:
        {
            "slug": "AAA",
        }

2. <mark style="background-color: grey">Метод удаления сегмента. Принимает slug (название) сегмента.</mark>

        URL: /segment

        Метод: DELETE

        Тело запроса:
        {
            "slug": "BBB",
        }

3. <mark style="background-color: grey">Метод добавления пользователя в сегмент. Принимает список slug (названий) сегментов которые нужно добавить пользователю, список slug (названий) сегментов которые нужно удалить у пользователя, id пользователя.</mark>

        URL: /bind

        Метод: POST

        Тело запроса:
        {
            "segmentsAdd": ["BBB"],
            "segmentsRemove": ["AAA"],
            "user_id": 3
        }

4. <mark style="background-color: grey">Метод получения активных сегментов пользователя. Принимает на вход id пользователя.</mark>

        URL: /bind

        Метод: GET

        Тело запроса:
        {
            "user_id": 1001
        }

## Опциональные задания

*Доп. задание 1:*

Иногда пользователи приходят в поддержку и спрашивают почему у них пропал/появился какой-то новый функционал. Нужно иметь возможность посмотреть когда точно пользователь попал в конкретный сегмент. 

Задача: реализовать сохранение истории попадания/выбывания пользователя из сегмента с возможностью получения отчета по пользователю за определенный период. На вход: год-месяц. На выходе ссылка на CSV файл.

    URL: /history

    Метод: GET

    Тело запроса:
    {
        "year": 2023,
        "month": 8
    }

*Доп. задание 3:*

Мы хотим добавлять пользователя в сегмент не в ручную, а автоматически. В сегмент будет попадать заданный процент пользователей.

Задача: в методе создания сегмента, добавить опцию указания процента пользователей, которые будут попадать в сегмент автоматически. В методе получения сегментов пользователя, добавленный сегмент должен отдаваться у заданного процента пользователей.

Пример: создали сегмент AVITO_VOICE_MESSAGES и указали что 10% пользователей будут попадать в него автоматически. Пользователь 1000 попал в этот сегмент автоматически. При запросе сегментов пользователя 1000, сегмент AVITO_VOICE_MESSAGES должен отдаваться всегда.

    URL: /segment

    Метод: POST

    Тело запроса:
    {
        "slug": "BBB",
        "auto_percentage": 80
    }

## Список заданий в разработке

## Опциональные задания

*Доп. задание 2:*

Бывают ситуации когда нам нужно добавить пользователя в эксперимент на ограниченный срок. Например выдать скидку всего на 2 дня. 

Задача: реализовать возможность задавать TTL (время автоматического удаления пользователя из сегмента)

Пример: Хотим чтобы пользователь попал в сегмент на 2 дня - для этого в метод добавления сегментов пользователю передаём время удаления пользователя из сегмента отдельным полем