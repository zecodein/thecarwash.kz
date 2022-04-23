The Car Wash
============
### Описание:
TODO

Построить и запустить:
----------------------

    make run

Остановить:
-----------

    make stop

Запустить:
-----------

    make start

Удалить:
--------
    
    make delete


JSON для регистрации пользователя:
----------------------------------

    Запрос на http://localhost:8080/user/signup

    {
        "name": "Имя",
        "number": "+7...",
        "password": "пароль",
        "confirm_password": "пароль"
    }

JSON для входа пользователя:
----------------------------

    Запрос на http://localhost:8080/user/signin
    
    {
        "number": "+7...",
        "password": "пароль",
    }