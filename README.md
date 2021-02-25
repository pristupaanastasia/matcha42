# matcha42

### `POST /api/user/login`

Запрос, возвращающий токен сессии для переданный входных данных пользователя.
Параметры:

* login: логин пользователя

* password: пароль пользователя

Ответ:

* 200 в случае успеха и токен пользовательской сессии

* 400 с описанием ошибки

### `POST /api/user/logout`

Запрос на завершение сессии
Параметры:

* login: логин пользователя

* password: пароль пользователя

Ответ:

* 200 в случае успеха 

* 400 с описанием ошибки

### `POST /api/user/registr`

Сохранение переданных входных данных пользователя.

Параметры:

* login: логин пользователя

* mail: электронная почта пользователя
  
* password: пароль пользователя

* first_name: имя пользователя

* last_name: фамилия пользователя

Необязательно

* gender: гендер пользователя

* age: возраст

* image: фотографии

* description: описание

* sex: сексуальные предпочтения

* tags: теги

* gps: локация пользователя

Ответ:

* 200 в случае успеха, id нового пользователя

* 400 с описанием ошибки

### `POST /api/user/reset_pass`

Восстановление пароля через ссылку

Параметры:

* reset_id: уникальный номер восстановления

* new_pass: новый пароль

* old_pass: старый пароль

### `POST /api/user/verification`

Подтверждение учетной записи

Параметры:

* ver_id: уникальный номер восстановления

* login: логин пользователя

### `POST /api/user/save_profile`

Редактирует информацию текущего профиля.

Параметры:

* login: логин пользователя

* password: пароль пользователя

* first_name: имя пользователя

* last_name: фамилия пользователя

* gender: гендер пользователя

* age: возраст

* image: фотографии

* description: описание

* sex: сексуальные предпочтения

* tags: теги
  
* gps: локация пользователя

Ответ:

* 200 в случае успеха

* 400 с описанием ошибки

### `GET /api/user/set_online`

Помечает пользователя как онлайн на 5 минут

Ответ : 

* true

### `GET /api/user/set_offline`

Помечает пользователя как оффлайн

Ответ :

* true

### `GET /api/research`

Запрос на показ списка рекомендованных пользователей.

* user_id: айди пользователя делающего запрос

* token: токен сессии пользователя делающего запрос

* sort_by: код выбора сортировки

Ответ:

* Возвращает 200-ОК если запрос отправлен удачно и список пользователей в формате json

* 400 с описанием ошибки

* 401 в случае если пользователь не авторизован


### `GET /api/profile/`

Запрос на показ страницы пользователя.

* user_id: айди пользователя делающего запрос

* token: токен сессии пользователя делающего запрос

* target_user_id: айди пользователя на страницу которого заходим.

Ответ:

* Возвращает 200-ОК если запрос отправлен удачно и профиль пользователя в формате json

* 400 с описанием ошибки

* 401 в случае если пользователь не авторизован

### `GET /api/user/like/`

Добавляет указанный профиль в список понравившихся.

* user_id: айди пользователя делающего запрос

* token: токен сессии пользователя делающего запрос

* target_user_id: айди пользователя которого мы хотим лайкнуть

Ответ:

* Возвращает 200-ОК если запрос отправлен удачно

* 400 с описанием ошибки

* 401 в случае если пользователь не авторизован

### `GET /api/user/connect`

Авторизованный запрос на добавление "В друзья" другого пользователя.

Параметры:

* user_id: id пользователя

* token: токен пользовательской сессии

* target_id: логин пользователя которого мы хотим добавить в друзья

Ответ:

* 200 в случае успеха операции

* 400 с описание ошибки

* 401 в случае если пользователь не авторизован

### `GET /api/user/disconnect/`

Закрывает доступ к чату.

* user_id: айди пользователя делающего запрос

* token: токен сессии пользователя делающего запрос

* target_user_id: айди пользователя которого мы хотим отключить

Ответ:

* Возвращает 200-ОК если запрос отправлен удачно

* 400 с описанием ошибки

* 401 в случае если пользователь не авторизован

### `GET /api/user/block/`

Запрос на блокирование пользователя.

* user_id: айди пользователя делающего запрос

* token: токен сессии пользователя делающего запрос

* target_user_id: айди пользователя которого мы хотим заблокировать

Ответ:

* Возвращает 200-ОК если запрос отправлен удачно

* 400 с описанием ошибки

* 401 в случае если пользователь не авторизован

### `GET /api/user/unblock/`

Запрос на разблокирование пользователя.

* user_id: айди пользователя делающего запрос

* token: токен сессии пользователя делающего запрос

* target_user_id: айди пользователя которого мы хотим разблокировать

Ответ:

* Возвращает 200-ОК если запрос отправлен удачно

* 400 с описанием ошибки

* 401 в случае если пользователь не авторизован

### `GET /api/user/history/`

Возвращает список юзеров, посетивших данный профиль.

* user_id: айди пользователя делающего запрос

* token: токен сессии пользователя делающего запрос


Ответ:

* Возвращает 200-ОК если запрос отправлен удачно и список пользователей

* 400 с описанием ошибки

* 401 в случае если пользователь не авторизован

### `GET /api/report/`

Отправка жалобы на аккаунт

* user_id: айди пользователя делающего запрос

* token: токен сессии пользователя делающего запрос

* target_user_id: айди пользователя аккаунта на чье имя жалоба

* report_id: id жалобы

Ответ:

* Возвращает 200-ОК если запрос отправлен удачно

* 400 с описанием ошибки

* 401 в случае если пользователь не авторизован

### ` /api/message/send/`

Отправляет сообщение

* user_id: айди пользователя делающего запрос

* token: токен сессии пользователя делающего запрос

* target_user_id: айди пользователя кому отправляется

* random_id: число int32 - уникальный индификатор для предотвращения повторного сообщения

* message: текст сообщения

Ответ:

* Возвращает 200-ОК если запрос отправлен удачно

* 400 с описанием ошибки

* 401 в случае если пользователь не авторизован

* 900 - Нельзя отправлять сообщение пользователю из черного списка

### ` /api/notifications/`

Отправляет уведомление пользователю

* user_id: айди пользователя кому отправить уведомление

* random_id: число int32 - уникальный индификатор для предотвращения повторного сообщения


Ответ:

* Возвращает 200-ОК если запрос отправлен удачно

* 400 с описанием ошибки

* 401 в случае если пользователь не авторизован


### Проблемы

* как сделать чат.
* как сделать оповещения.

