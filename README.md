# CRUD With GO!
CRUD with Go Language!

# How to use with Docker-Compose
- git clone this project
- open the cloned folder in the terminal and ```cd backend/```
- run this command ```sudo docker-compose up```
- check status, go to this url => ```localhost:8080```

# How to use with your local
- git clone this project
- open the cloned folder in the terminal and ```cd backend/```
- change db_host field value ```db``` to ```localhost``` in .env file
- run this command ```go build && go run main```
- check status, go to this url => ```localhost:8080```

# Postman Documentation
You can find the postman collection in the documentation folder.

# Test
You can test user model with this command ```go test -v ./models/```
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