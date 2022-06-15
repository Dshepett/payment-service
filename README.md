# Payment service
Простой API симулирующий работу платежной системы.
API имеет 7 роутов:
- **/payments/new** - создает новый платеж со статусом НОВЫЙ.
- **/payments/{id}/change** - изменяет статус платежа по его id, при его наличии и статусом НОВЫЙб на УСПЕХ или НЕУСПЕХ.
- **/payments/{id}/status** - возвращает статус платежа(НОВЫЙ, ОШИБКА, УСПЕХ, НЕУСПЕХ, ОТМЕНЕН) по id при его сузествовании.
- **/payments/{id}/deny** - устанавливает статус ОТМЕНЕН на платеж по id, при его существовании и статусе НОВЫЙ.
- **/payments/user/{id}** - возвращает все платежи пользоваетля по его id.
- **/payments/user/email/{email}** - возвращает все платежи польоваетля по его email.
- **/auth/login** - выполняет авторищацию, которая позволяет потом менять статус платежей(возвращает токе, который потом надо будет добавить в заголовок и перед ним написать Bearer).

Еще об API:
- API работает на порте 8080
- для хранилища была использована база данных PostgreSQL
- для удобной проверки работоспособности добавлен Swagger(**/swagger/index.html**)
- написаны тесты для сервисного слоя
- добавлен docker-compose файля для развертывания API на компьютере
- добавлен Makefile для упрощения написания часто-используемых команд
- реализова Graceful-shutdown
- добавлена авторизация для смены статуса платежа
