# Avito_Start_2023
тестовое задание для Авито https://github.com/avito-tech/backend-trainee-assignment-2023

## Команды для запуска 
```bash
git clone https://github.com/Sib-Coder/avitoStart
cd avitoStart
sudo docker-compose build --no-cache  
sudo docker-compose up   
```

![dockerstart](https://github.com/Sib-Coder/avitoStart/blob/main/images/startdoker.png)


## Документация по API находится https://documenter.getpostman.com/view/24934668/2s9Y5bNfhd

## Примеры работы API
### Создаём Юзера

![createuser](https://github.com/Sib-Coder/avitoStart/blob/main/images/create_user.png)
### Проверяем создание
![users](https://github.com/Sib-Coder/avitoStart/blob/main/images/exec_users.png)
### Создаём сегмент
![createslug](https://github.com/Sib-Coder/avitoStart/blob/main/images/create_slug.png)

### Добавляем сегмент Юзеру
![addsluguser](https://github.com/Sib-Coder/avitoStart/blob/main/images/addsluguser.png)

### Проверяем 
![ecexsluguser1](https://github.com/Sib-Coder/avitoStart/blob/main/images/ecexsluguser1.png)

### Добавляем и удаляем различные сегменты (которые создали заранее)
![addanddelsluguser](https://github.com/Sib-Coder/avitoStart/blob/main/images/addanddelsluguser.png)
### Проверяем
![ecexuserslug2](https://github.com/Sib-Coder/avitoStart/blob/main/images/ecexuserslug2.png)
### Удаляем Сегмент
![deleteslug](https://github.com/Sib-Coder/avitoStart/blob/main/images/deleteslug.png)
### Удаляем пользователя
![deleteuser](https://github.com/Sib-Coder/avitoStart/blob/main/images/deleteuser.png)

## Решение доп задания 2 TTL
### Не уверен что успею реализовать , но решение напишу.
В моей реализации можно добавить в таблицу slugtraker дополнительное поле date_validity в котором можно задать дефолтное значение и менять которое  на (сегодняшнюю дату + время валидности).</br>
При получении данных из базы данных мы должны будем проверять это поле если оно у нас соответствует дефолтному значению то просто выдаём данные , а если дата то :
*  Проверяем на валидность даты , если дата позже сегодняшеней ,то удаляем сегмент у пользователя , если нет то просто выдаём.