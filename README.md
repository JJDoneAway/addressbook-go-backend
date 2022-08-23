# Simple Enterprise Application 

It provides some REST CRUD services to get, store, update and delete a simple entity

## Run it
Just run 
```
go run github.com/JJDoneAway/addressbook
````
and open

http://localhost:8080/users

----
## Test it

### Get all
```
curl -v "http://localhost:8080/users" | json_pp
```

### Get one
#### positive
```
curl -v "http://localhost:8080/users/17281579" | json_pp
```
#### negative
```
curl -v "http://localhost:8080/users/pimmel" | json_pp
```
```
curl -v "http://localhost:8080/users/000" | json_pp
```

### Post one
#### positive
```
curl -v  -X POST "http://localhost:8080/users" -d '{"FirstName":"Johannes","LastName":"Höhne"}' | json_pp
```
#### negative
```
curl -v  -X POST "http://localhost:8080/users" -d '{"Pimmel":"Johannes","Pummel":"Höhne"}' | json_pp
```
```
curl -v  -X POST "http://localhost:8080/users" -d '{"ID":1234,"FirstName":"Johannes","LastName":"Höhne"}' | json_pp
```
```
curl -v  -X POST "http://localhost:8080/users/1234" -d '{"FirstName":"Johannes","LastName":"Höhne"}' | json_pp
```

### Put one
#### positive
```
curl -v -X PUT "http://localhost:8080/users/16888363" -d '{"ID": 16888363, "FirstName":"Johannes","LastName":"Höhne"}' | json_pp
```
#### negative
```
curl -v -X PUT "http://localhost:8080/users/16888363" -d '{"ID": 11111, "FirstName":"Johannes","LastName":"Höhne"}' | json_pp
```
```
curl -v -X PUT "http://localhost:8080/users/666" -d '{"ID": 666, "FirstName":"Johannes","LastName":"Höhne"}' | json_pp
```

### Delete one
#### positive
```
curl -v -X DELETE "http://localhost:8080/users/16888363" | json_pp
```
#### negative
```
curl -v -X DELETE "http://localhost:8080/users/666" | json_pp
```

### Delete all
#### positive
```
curl -v -X DELETE "http://localhost:8080/users/" | json_pp
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