# CRUD With GO!
Crud with Go Language!

# How to use
- git clone this project
- open the cloned folder on terminal
- run this command ```sudo docker-compose up```
- check status, go to this url => ```localhost:8080```


# Documentation
---

#### Add a user

##### Sample Request
```
curl -X PUT \
 -d '{"name": "Test", "email": "test@example.com", "password": "securepasswd"}' \
 -H 'Content-Type: application/json' \
  http://localhost:8080/users
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success | {"id": 1, "name": "Test", "email": "test@example.com"} |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 403 | If user already exists | {"error": "User with that email already exists"} |
| 500 | When something unexpected happens | {"error": "server error"} |

---

#### Edit a user's attributes

##### Sample Request
```
curl -X PATCH \
 -d '{"name": "No name", "password": "strongpasswd"}' \
 -H 'Content-Type: application/json' \
  http://localhost:8080/users/1
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success | {"id": 1, "name": "No name", "email": "test@example.com"} |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 404 | If user not found | {"error": "User with that id does not exist"} |
| 500 | When something unexpected happens | {"error": "server error"} |

---

#### Delete a user

##### Sample Request
```
curl -X DELETE \
  http://localhost:8080/users/1
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success |  |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 404 | If user not found | {"error": "User with that id does not exist"} |
| 500 | When something unexpected happens | {"error": "server error"} |

---

#### Find a user with ID

##### Sample Request
```
curl -X GET \
  http://localhost:8080/users/1
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success | {"id": 1, "name": "No name", "email": "test@example.com"} |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 404 | If user not found | {"error": "User with that id does not exist"} |
| 500 | When something unexpected happens | {"error": "server error"} |


---

#### Get All Users

##### Sample Request
```
curl -X GET \
  http://localhost:8080/users
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success | [{"id": 1, "name": "No name", "email": "test@example.com"}, {"id": 2, "name": "Test 2", "email": "test2@example.com"}] |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 404 | If user not found | {"error": "User with that id does not exist"} |
| 500 | When something unexpected happens | {"error": "server error"} |


---