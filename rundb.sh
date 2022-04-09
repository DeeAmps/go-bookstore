#! /bin/bash

docker run -d \
    --name postgres \
    -p 5432:5432 \
    -e POSTGRES_PASSWORD=password \
    -e POSTGRES_USER=dee \
    -e POSTGRES_DB=bookstore \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v ~/Documents/data:/var/lib/postgresql/data \
    postgres