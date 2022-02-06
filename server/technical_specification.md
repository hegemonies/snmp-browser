# snmp-browser server

- [snmp-browser server](#snmp-browser-server)
  - [Требование](#требование)
    - [Первая версия](#первая-версия)
    - [Будущие версии](#будущие-версии)
  - [Описание требований](#описание-требований)
    - [Поддержка *SNMP GET и WALK* методов.](#поддержка-snmp-get-и-walk-методов)
    - [Поддержка работы в режиме *cli*](#поддержка-работы-в-режиме-cli)
      - [Команда _get_](#команда-get)
      - [Команда _walk_](#команда-walk)
      - [Команда _help_](#команда-help)
    - [Поддержка работы в режиме *http* сервера с возможностью общения через *HTTP* и *WebSocket*.](#поддержка-работы-в-режиме-http-сервера-с-возможностью-общения-через-http-и-websocket)

## Требование

### Первая версия

1. Поддержка *SNMP GET и WALK* методов.
2. Поддержка работы в режиме *cli*.
   1. Поддержка опций *port*, *communities*, *version*, *retries*, *timeout*.
   2. Поддержка команд *get*, *walk* и *help*.
3. Поддержка работы в режиме *http* сервера с возможностью общения через *HTTP* и *WebSocket*.

### Будущие версии

1. Поддержка *SNMP V3*.
2. Поддержка *MIB* парсера.
3. Поддержка *SNMP* трапов.
4. Поддержка  *SNMP BULKGET, BULKWALK, PING, GETNEXT, STATUS, SET, TRAP* методов.

## Описание требований

### Поддержка *SNMP GET и WALK* методов. 

Используя библиотеку _gosnmp_ реализуется апи для поддержки методов SNMP GET и WALK.

Методы требуют параметров:
* ip/hostname
* snmp port
* snmp communities
* snmp version
* snmp oids
* timeout
* number of retries

### Поддержка работы в режиме *cli*

Приложение snmp-browser должно выполнять команды из консоли, опциями задавать параметры 
snmp port, communities, version, timeout, number of retries, а командами название метода - get
или walk, и ip/hostname с snmp iods.

Значения по умолчанию для опций:
* snmp port: 161
* snmp communities: [public]
* snmp version: Version 2c
* timeout: 5000 ms
* number of retries: 3

#### Команда _get_

Выполняет SNMP GET запрос на указанный ip/host с указанным SNMP oids для получения
одного значения.

#### Команда _walk_

Выполняет SNMP WALK запрос на указанный ip/host для получения
множества значений по каждому списку SNMP oids.

#### Команда _help_

Выводит всю справочную информацию по опциям и командам snmp-browser.

###  Поддержка работы в режиме *http* сервера с возможностью общения через *HTTP* и *WebSocket*.

TODO...
