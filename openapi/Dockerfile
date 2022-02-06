FROM jamescooke/openapi-validator
# Для CI необходимо переопределение entrypoint для того,
# чтобы его можно было использовать как базовый образ выполнения stage'ей.
# Иначе будет запускаться бинарник - lint-openapi.
ENTRYPOINT sh
CMD lint-openapi
