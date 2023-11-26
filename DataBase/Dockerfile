FROM postgres:14-alpine

ENV POSTGRES_PASSWORD=postgres 

COPY ./createscript.sql /docker-entrypoint-initdb.d/

