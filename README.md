# uboRequest

Библиотека для ограничения числа http-запросов за опеределенный период. Использует http://github.com/nikonor/uroboros.

Пример использования в каталоге example. В примере указан мой МОК, которые по данной ссылке выдает просто OK-200. В моем случае программа имеет следующий вывод

```
Try#0::ждем 200
        resp: status=200 OK
Try#1::ждем 200
        resp: status=200 OK
Try#2::ждем 200
        resp: status=200 OK
Try#3::ждем 200
        resp: status=200 OK
Try#4::ждем 200
        resp: status=200 OK
Try#5::ждем 200
        resp: status=200 OK
Try#6::ждем 200
        resp: status=200 OK
Try#7::ждем 200
        resp: status=200 OK
Try#8::ждем 200
        resp: status=200 OK
Try#9::ждем 200
        resp: status=200 OK
Try#10::ждем ошибку
        err=can't do request now
Try#11::ждем ошибку
        err=can't do request now
Try#12::ждем ошибку
        err=can't do request now
Try#13::ждем ошибку
        err=can't do request now
Try#14::ждем ошибку
        err=can't do request now

```