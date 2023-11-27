Что было сделано с предыдущего раза:

1. Небольшие изменения в обработке телефонов (pkg\phone.go);
2. Обработка только POST-запросов (controller\httpserver\httpserver.go);
3. Логгирование ошибок в файл и в консоль (pkg\eWrapper.go);
4. Внедрение структуры Response (models\dto\dto.go и controller\httpserver\httpserver.go);
5. Исправление возможности создания записей с одинаковыми номерами (gates\psg\queries.go). 