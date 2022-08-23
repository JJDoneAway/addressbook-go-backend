# Simple Enterprise Application 

> This branch hosts the GIN GONIC implementation.
>
> GIN is used to do all the routing and error handling.

It provides some REST CRUD services to get, store, update and delete a simple entity

## Run it
Just run 
```
git clone https://github.com/JJDoneAway/addressbook.git
cd addressbook
go run github.com/JJDoneAway/addressbook

or

go build github.com/JJDoneAway/addressbook
./addressbook
````
and open

http://localhost:8080/users

----
## Test it

### Get all
```
curl -L -v "http://localhost:8080/users/" | json_pp
```

### Get one
#### positive
```
curl -L -v "http://localhost:8080/users/17281579/" | json_pp
```
#### negative
```
curl -L -v "http://localhost:8080/users/pimmel/" | json_pp
```
```
curl -L -v "http://localhost:8080/users/000/" | json_pp
```

### Post one
#### positive
```
curl -L -v  -X POST "http://localhost:8080/users" -d '{"FirstName":"Johannes","LastName":"Höhne"}' | json_pp
```
#### negative
```
curl -L -v  -X POST "http://localhost:8080/users" -d '{"Pimmel":"Johannes","Pummel":"Höhne"}' | json_pp
```
```
curl -L -v  -X POST "http://localhost:8080/users" -d '{"ID":1234,"FirstName":"Johannes","LastName":"Höhne"}' | json_pp
```
```
curl -L -v  -X POST "http://localhost:8080/users/1234" -d '{"FirstName":"Johannes","LastName":"Höhne"}' | json_pp
```

### Put one
#### positive
```
curl -L -v -X PUT "http://localhost:8080/users/16888363" -d '{"ID": 16888363, "FirstName":"Johannes","LastName":"Höhne"}' | json_pp
```
#### negative
```
curl -L -v -X PUT "http://localhost:8080/users/16888363" -d '{"ID": 11111, "FirstName":"Johannes","LastName":"Höhne"}' | json_pp
```
```
curl -L -v -X PUT "http://localhost:8080/users/666" -d '{"ID": 666, "FirstName":"Johannes","LastName":"Höhne"}' | json_pp
```

### Delete one
#### positive
```
curl -L -v -X DELETE "http://localhost:8080/users/16888363/" | json_pp
```
#### negative
```
curl -L -v -X DELETE "http://localhost:8080/users/666/" | json_pp
```

### Delete all
#### positive
```
curl -L -v -X DELETE "http://localhost:8080/users/" | json_pp
```
####
nothing to test here




---
## Next Steps

- [x] Implement all CRUDs
- [ ] Add SWAGGER
- [ ] Add Examples out of file
- [ ] Add OAuth
- [ ] Add Real DB
- [ ] Add Prometheus metrics