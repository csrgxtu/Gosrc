##An Simple JWT(Json Web Token) Tutorial In Golang.

### Usage
before testing jwt, u need 2 start simple http server  
```bash
go run server.go
```

first, auth and get the Token  
Request:
```bash
curl http://localhost:3000/auth
```
Response:
```json
{
  "code": 200,
  "msg":"Successful",
  "data":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IlRoaXMgaXMgbXkgc3VwZXIgZmFrZSBJRCIsImV4cCI6MTQ0NTYxODU1OH0.J_9Hrjkey-L-cGt5PbQU2UmKMO-26bUuJk-UVBJjiB3oQHf9T20uRxtnr0KKvd7_ps094r2GQqOUMrV6Qt6GWGYHPYevC_PukKXgobwvP4t8NGDGRj5dbeivD4vg2Rm8Od2H6TvBtPNJpMYsr65ozTejnfJj"
}
```

second, use the token from first step do a new request  
Request
```bash
curl -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IlRoaXMgaXMgbXkgc3VwZXIgZmFrZSBJRCIsImV4cCI6MTQ0NTYxODU1OH0.J_9Hrjkey-L-cGt5PbQU2UmKMO-26bUuJk-UVBJjiB3oQHf9T20uRxtnr0KKvd7_ps094r2GQqOUMrV6Qt6GWGYHPYevC_PukKXgobwvP4t8NGDGRj5dbeivD4vg2Rm8Od2H6TvBtPNJpMYsr65ozTejnfJj" http://localhost:3000/users
```
Response:
```json
{
  "code":200,
  "msg":"Successful",
  "data":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IlRoaXMgaXMgbXkgc3VwZXIgZmFrZSBJRCIsImV4cCI6MTQ0NTYxODU1OH0.J_9Hrjkey-L-cGt5PbQU2UmKMO-26bUuJk-UVBJjiB3oQHf9T20uRxtnr0KKvd7_ps094r2GQqOUMrV6Qt6GWGYHPYevC_PukKXgobwvP4t8NGDGRj5dbeivD4vg2Rm8Od2H6TvBtPNJpMYsr65ozTejnfJj"
}
```
otherwise
```json
{
  "code":401,
  "msg":"Token Unauthorized",
  "data":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IlRoaXMgaXMgbXkgc3VwZXIgZmFrZSBJRCIsImV4cCI6MTQ0NTYxODU1OH0.J_9Hrjkey-L-cGt5PbQU2UmKMO-26bUuJk-UVBJjiB3oQHf9T20uRxtnr0KKvd7_ps094r2GQqOUMrV6Qt6GWGYHPYevC_PukKXgobwvP4t8NGDGRj5dbeivD4vg2Rm8Od2H6TvBtPNJpMYsr65ozTejnfJj"
}
```
