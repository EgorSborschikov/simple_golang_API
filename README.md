# Библиотечное API на Go

Данное веб-приложение на языке Go реализует простой RESTful API для управления библиотекой книг. Оно позволяет получать, создавать, обновлять и удалять книги.

## Библиотеки

###  bufio: 
Эта библиотека предоставляет буферизованный ввод-вывод. Она позволяет читать и записывать данные в более эффективном режиме, используя буфер, что уменьшает количество операций ввода-вывода. Например, с помощью bufio можно читать строки из файла или стандартного ввода, не выполняя каждую операцию чтения отдельно.

### encoding/json: 
Библиотека для работы с JSON (JavaScript Object Notation). Она позволяет кодировать (сериализовать) и декодировать (десериализовать) данные в формате JSON. Это особенно полезно для работы с API и обмена данными между клиентом и сервером.

### log: 
Библиотека для ведения логов. Она предоставляет функции для записи логов на стандартный вывод, в файлы или другие источники. Это полезно для отслеживания работы программы и отладки.

### math/rand: 
Библиотека для работы со случайными числами. Она предоставляет функции для генерации случайных чисел, как целых, так и с плавающей запятой. Это может быть полезно в различных приложениях, таких как игры или симуляции.

### net/http: 
Библиотека для работы с HTTP. Она позволяет создавать HTTP-серверы и клиентов, обрабатывать запросы и ответы, работать с URL и многим другим. Это основа для создания веб-приложений на Go.

### os: 
Библиотека для взаимодействия с операционной системой. Она предоставляет функции для работы с файловой системой, окружением, процессами и другими системными ресурсами.

### strconv: 
Библиотека для преобразования строк в другие типы данных и обратно. Она позволяет, например, преобразовывать строки в целые числа или числа с плавающей запятой и наоборот.

### github.com/gorilla/mux: 
Это пакет для маршрутизации HTTP-запросов в Go. Он позволяет определять маршруты и обрабатывать запросы по определённым URL-шаблонам. Это упрощает создание RESTful API и веб-приложений.


## Структуры данных

### Book

Структура, представляющая книгу, содержит следующие поля:

- **ID**: уникальный идентификатор книги.
- **Title**: название книги.
- **Author**: указатель на структуру `Author`, представляющую автора книги.

### Author

Структура, представляющая автора книги, с полями:

- **Firstname**: имя автора.
- **Lastname**: фамилия автора.

## Глобальные переменные

- **books**: Срез, который хранит список книг.

## Функции обработчики

### `getBooks`

- Устанавливает заголовок ответа на `application/json`.
- Выполняет HTTP-запрос к Open Library для получения списка книг (лимит 5).
- Декодирует ответ и заполняет срез `books` данными о книгах.
- Возвращает список книг в формате JSON.

### `getBook`

- Устанавливает заголовок ответа.
- Получает параметр `id` из URL и ищет книгу с соответствующим ID.
- Если книга найдена, возвращает её в формате JSON, иначе возвращает пустую книгу.

### `createBook`

- Устанавливает заголовок ответа.
- Читает данные о новой книге из консоли (название, имя и фамилию автора).
- Создает новую книгу и добавляет её в срез `books`.
- Возвращает созданную книгу в формате JSON.

### `updateBook`

- Устанавливает заголовок ответа.
- Получает параметр `id` из URL и ищет книгу с соответствующим ID.
- Если книга найдена, удаляет её из среза, декодирует новые данные из тела запроса и добавляет обновленную книгу обратно в срез.
- Возвращает обновленную книгу в формате JSON.

### `deleteBook`

- Устанавливает заголовок ответа.
- Получает параметр `id` из URL и ищет книгу с соответствующим ID.
- Если книга найдена, удаляет её из среза и возвращает обновленный список книг в формате JSON.

## Функция main

- Создает новый маршрутизатор с помощью `mux`.
- Инициализирует срез `books` с некоторыми книгами по умолчанию.
- Определяет маршруты для каждого из обработчиков (GET, POST, PUT, DELETE).
- Запускает HTTP-сервер на порту 8080 и обрабатывает ошибки при запуске.

## Запуск сервера

В конце кода сервер запускается, и в лог выводится сообщение о том, что он работает на порту 8080.

## Использование

Для использования API отправляйте HTTP-запросы на следующие маршруты:

- `GET /books` - получить список книг.
- `GET /books/{id}` - получить книгу по ID.
- `POST /books` - создать новую книгу.
- `PUT /books/{id}` - обновить книгу по ID.
- `DELETE /books/{id}` - удалить книгу по ID.