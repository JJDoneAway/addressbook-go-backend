# Simple Enterprise Application 

It provides some REST CRUD services to get, store, update and delete a simple entity

> 
> I used branches to try out different implementation styles. 
>
> so please check the branches
>


## Run it
---------
### Check out the project 
```
git clone https://github.com/JJDoneAway/addressbook.git
cd addressbook
```

### Switch to the branch of your interest
```
git switch vanilla
```

or

```
git switch gin
```

### Run the GO code
```
go run github.com/JJDoneAway/addressbook
```

or

```
go build github.com/JJDoneAway/addressbook
./addressbook
```

### Open the application in your browser

http://localhost:8080/addresses

## Swagger
----------
### Install swag
https://github.com/swaggo/swag/blob/master/example/basic/api/api.go

### Add description
Add to all exposed REST endpoints the corresponding swagger description

e.g.: https://github.com/swaggo/swag/blob/master/example/basic/api/api.go

### Create swagger.json stuff
run `swag init` in your root 

### Show
open http://localhost:8080


## Test it
----------

> The IDs of the users are generated and change from run to run. So we need to store one in an env and use it afterwards

### Get one ID for your examples
```
myID=$(curl localhost:8080/addresses | jq '.[0]' | jq '.id') 
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




---
## Next Steps

- [x] Implement all CRUDs in vanilla GO
- [x] Implement all CRUDs in GIN
- [x] Add SWAGGER (it is a nightmare)
- [x] Add Prometheus metrics
- [x] Add Examples out of file (Using embed files, to have it in the executable)
- [ ] Add OAuth
- [ ] Add Real DB

