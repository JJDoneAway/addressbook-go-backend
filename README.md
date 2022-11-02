# Simple Enterprise Application 

> This branch hosts the GORM implementation.
>
> GORM is used to do the data base conversation in an most easy way.
>
> For detailed description please see main branch
>
> This branch is based on the gin branch. So it is already using GIN, Swagger and Prometheus


# GORM
Follow the instructions of https://gorm.io/

Open http:localhost:8080

# Run code
1. Star local a PostgreSQL 
```
docker run --rm -p 5432:5432 -e POSTGRES_PASSWORD="1234" --name pg postgres:alpine
```

2. Check if your DB id running and the inserted stuff
```
pgcli postgresql://postgres:1234@localhost:5432/postgres
```

3. Start code
```
go run .
```


