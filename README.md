# Simple Enterprise Application with DB

> This branch hosts the GORM implementation.
>
> GORM is used to do the data base conversation in an most easy way.
>
> For detailed description please see main branch
>
> This branch is based on the gin branch. So it is already using GIN, Swagger and Prometheus


## GORM
Follow the instructions of https://gorm.io/

Open http:localhost:8080

## Run code
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

## Test it
----------

> The IDs of the users are generated and change from run to run. So we need to store one in an env and use it afterwards

### Get one ID for your examples
```
myID=$(curl localhost:8080/addresses | jq '.[0]' | jq '.ID') 
```


### Get all
```
curl -v "http://localhost:8080/addresses" | json_pp
```

### Get one
#### positive
```
curl -v "http://localhost:8080/addresses/$myID" | json_pp
```
#### negative
```
curl -v "http://localhost:8080/addresses/pimmel" | json_pp
```
```
curl -v "http://localhost:8080/addresses/000" | json_pp
```

### Post one
#### positive
```
curl -v  -X POST "http://localhost:8080/addresses" -d '{"first-name":"Johannes","last-name":"Höhne", "email":"Johannes@hoehne.de", "phone":"+49123456789"}' | json_pp
```
#### negative
```
curl -v  -X POST "http://localhost:8080/addresses" -d '{"first-name":"Johannes","last-name":"Höhne", "email":"NotARealMail", "phone":"123456789"}' | json_pp
```
```
curl -v  -X POST "http://localhost:8080/addresses" -d '{"id":1234567,"first-name":"Johannes","last-name":"Höhne", "email":"Johannes@hoehne.de", "phone":"+49123456789"}' | json_pp
```

### Put one
#### positive
```
curl -v -X PUT "http://localhost:8080/addresses/$myID" -d '{"id": '${myID}', "first-name":"Johannes","last-name":"Höhne", "email":"Johannes@hoehne.de", "phone":"+49123456789"}' | json_pp
```
#### negative
```
curl -v -X PUT "http://localhost:8080/addresses/$myID" -d '{"id": 666, "first-name":"Johannes","last-name":"Höhne", "email":"Johannes@hoehne.de", "phone":"+49123456789"}' | json_pp
```
```
curl -v -X PUT "http://localhost:8080/addresses/666" -d '{"id": 666, "first-name":"Johannes","last-name":"Höhne", "email":"Johannes@hoehne.de", "phone":"+49123456789"}' | json_pp
```

### Delete one
#### positive
```
curl -v -X DELETE "http://localhost:8080/addresses/$myID" | json_pp
```
#### negative
```
curl -v -X DELETE "http://localhost:8080/addresses/666" | json_pp
```

### Delete all
#### positive
```
curl -v -X DELETE "http://localhost:8080/addresses/"
```
#### negative
nothing to test here




