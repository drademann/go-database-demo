# Go Database Demo

A demo accessing a (local) MongoDB and Postgres instance.

## Links

Mongo Quickstart:
https://www.mongodb.com/blog/post/quick-start-golang-mongodb-starting-and-setup

GORM:
https://gorm.io/docs/


## Databases

### Mongo

```shell
$ docker run -it --name mongo -p 27017:27017 -d mongo
```

### PostgreSQL

```shell
$ docker run -it --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
```
