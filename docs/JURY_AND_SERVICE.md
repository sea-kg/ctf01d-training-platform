# Как собрать два архива для журейной системы и для vuln образа

## Исходные данные: Архивы / репозитории с уязвимыми сервисами

Структура:

- Директория `service` - надо будет закинуть в vuln образ (Внутри есть README.md - можно взять как дескрипшн сервиса или типа того)
- Директория `checker` - Надо будет переименовать в `checker_example_service1_py` и как есть скопировать в образ журейки в `data_game`
- Директория `writeup` - от автора чтобы было что разбирать
- Директория `exploits` - тоже от автора что бы можно было тестировать уязвимости
- Файл `.ctf01d-service.yml` - конфигурационный файл для `ctf01d-training-platform` там есть id сервиса + часть конфига для журейной системы (с версионированием)
- Файл `LICENSE` - лицензия (вроде как не обязательно но пусть будет)
- Файл `README.md` - Описание общее для разработчиков сервиса (можно игнорировать)

Репозитории с примерами:

1. https://github.com/sea-kg/ctf01d-service-example1-py
2. https://github.com/sea-kg/ctf01d-service-example2-php


vuln-image:
- Директорию `service` копируем с переименовыванием в `%id-of-service%`, где id-of-service это id из файла `.ctf01d-service.yml` в секции `checker-config-*`/`id`

jury-image:
- Директорию `checker` копируем с переименовыванием в `data_game/checker_%id-of-service%`, где id-of-service это id из файла `.ctf01d-service.yml` в секции `checker-config-*`/`id`


## Образ с журейной системой


`docker-compose.yml` (в нем может поменяться только версия `v0.5.2` нуу и может container name):
```
version: '3'

services:
  ctf01d-jury:
    build: .
    container_name: ctf01d_jury_game1
    volumes:
      - "./data_game:/usr/share/ctf01d"
    environment:
      CTF01D_WORKDIR: "/usr/share/ctf01d"
    ports:
      - "8080:8080"
    restart: always
    networks:
      - ctf01d_net

networks:
  ctf01d_net:
    driver: bridge
```

`Dockerfile` - надо будет из файла `.ctf01d-service.yml` из секций `install-checker-requirements-v0.5.2` добавить команды.
```
FROM sea5kg/ctf01d:v0.5.2

# checker_ctf01d-service-example1-py

# nothing

# checker_ctf01d-service-example2-php
# copied from https://github.com/sea-kg/ctf01d-service-example2-php/blob/main/.ctf01d-service.yml

RUN apt-get -y update
RUN apt install -y python3 python3-pip python3-pip python3-requests
```

`data_game/config.yml`
```
## Combined config for ctf01d
# use 2 spaces for tab

game:
  id: "game" # uniq gameid must be regexp [a-z0-9]+
  name: "Game1" # visible game name in scoreboard
  start: "2023-11-12 16:00:00" # start time of game (UTC)
  end: "2030-11-12 22:00:00" # end time of game (UTC)
  coffee_break_start: "2023-11-12 20:00:00" # start time of game coffee break (UTC), but it will be ignored if period more (or less) then start and end
  coffee_break_end: "2023-11-12 21:00:00" # end time of game coffee break (UTC), but it will be ignored if period more (or less) then start and end
  flag_timelive_in_min: 1 # you can change flag time live (in minutes)
  basic_costs_stolen_flag_in_points: 1 # basic costs stolen (attack) flag in points for adaptive scoreboard
  cost_defence_flag_in_points: 1.0 # cost defences flag in points

scoreboard:
  port: 8080 # http port for scoreboard
  htmlfolder: "./html" # web page for scoreboard see index-template.html
  random: no # If yes - will be random values in scoreboard

checkers:
  - id: "example_service1_py" # work directory will be checker_example_service1_py # copied from https://github.com/sea-kg/ctf01d-service-example1-py/blob/main/.ctf01d-service.yml
    service_name: "Service1 Py"
    enabled: yes
    script_path: "./checker.py"
    script_wait_in_sec: 5 # max time for running script
    time_sleep_between_run_scripts_in_sec: 15
  - id: "example_service2_php" # work directory will be checker_example_service2_php # copied from https://github.com/sea-kg/ctf01d-service-example2-php/blob/main/.ctf01d-service.yml
    service_name: "Service2 PHP"
    enabled: yes
    script_path: "./checker.py"
    script_wait_in_sec: 5 # max time for running script
    time_sleep_between_run_scripts_in_sec: 15

teams:
  - id: "t01" # must be uniq
    name: "Team #1"
    active: yes
    logo: "./html/images/teams/team01.png"
    ip_address: "127.0.1.1" # address to vulnserver
  - id: "t02" # must be uniq
    name: "Team #2"
    active: yes
    logo: "./html/images/teams/team02.png"
    ip_address: "127.0.2.1" # address to vulnserver
```


Еще надо скопировать файлы из директории `checker` (для каждого сервиса) в `data_game/checker_%id-of-service%`, где id-of-service это id из файла `.ctf01d-service.yml` в секции `checker-config-*`/`id`

то есть:

- в директорию `data_game/checker_example_service1_py` скопировать файлы из https://github.com/sea-kg/ctf01d-service-example1-py/tree/main/checker
- в директорию `data_game/checker_example_service2_php` скопировать файлы из https://github.com/sea-kg/ctf01d-service-example2-php/tree/main/checker
