# Project: task_manager2.0

## End-point: Register
# Register User

This endpoint is used to register a new user.

## Request Body

- `name` (string, required): The name of the user.
    
- `email` (string, required): The email address of the user.
    
- `password` (string, required): The password for the user.
    
- `isadmin` (boolean, required): Indicates whether the user is an admin or not.
    

## Response

The response for this request is a JSON object with the following schema:

``` json
{
  "userId": "string",
  "name": "string",
  "email": "string",
  "isAdmin": "boolean"
}

 ```
### Method: POST
>```
>http://localhost:8080/api/register
>```
### Body (**raw**)

```json
{
    "name":"mer",
    "email": "mr@gmail.com",
    "password":"mer123",
    "isadmin": true
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Login
### API POST /api/login

This endpoint is used to authenticate a user and obtain a token for accessing protected resources.

#### Request Body

- `email` (string, required): The email address of the user.
    
- `password` (string, required): The password of the user.
    

#### Response

- Status: 200
    
- Content-Type: application/json
    

``` json
{
    "token": "********",
    "user": {
        "id": "********",
        "email": "********",
        "isadmin": true
    }
}

 ```

The response contains a token for authentication and user information including user ID, email, and admin status.
### Method: POST
>```
>http://localhost:8080/api/login
>```
### Body (**raw**)

```json
{

    "email": "mr@gmail.com",
    "password":"mer123"

}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Create Task
### API Request Description

This POST request is used to create a new task by providing the title and description in the request body.

### Request Body

- `title` (string, required): The title of the task.
    
- `description` (string, required): The description of the task.
    

### Response

The response will have a status code of 201 and a JSON content type. The response body will include a message indicating the success or any additional information related to the task creation.

### JSON Schema for Response

``` json
{
  "type": "object",
  "properties": {
    "message": {
      "type": "string"
    }
  }
}

 ```
### Method: POST
>```
>http://localhost:8080/api/tasks
>```
### Body (**raw**)

```json
{
    "title" : "blah",
    "description" : "mmm"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{jwt_token}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get tasks under user
This endpoint makes an HTTP GET request to retrieve a list of tasks. The request does not require any specific parameters. The response will be in JSON format and will include an array of task objects, each containing an ID, title, description, due date, and user information including ID, name, email, and admin status.
### Method: GET
>```
>http://localhost:8080/api/tasks
>```
### Body (**raw**)

```json
{
    "title" : "updated",
    "description" : "the description of your task"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{jwt_token}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get task by id under user
### GET /api/tasks/{taskId}

This endpoint retrieves a specific task identified by its unique ID.

#### Request

No request body is required for this endpoint.

- Path Parameters
    
    - `taskId` (string, required): The unique ID of the task.
        

#### Response

The response will be a JSON object with the following schema:

``` json
{
  "task_id": "string",
  "title": "string",
  "description": "string",
  "status": "string",
  "created_at": "string",
  "updated_at": "string"
}

 ```
### Method: GET
>```
>http://localhost:8080/api/tasks/0
>```
### Body (**raw**)

```json
{
    "title" : "the title of your task",
    "description" : "the description of your task"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update task under user
## Update Task

This endpoint is used to update a specific task identified by its ID.

### Request Body

- `title` (string, required): The updated title of the task.
    
- `description` (string, required): The updated description of the task.
    

### Response

The response will include the updated details of the task after the successful update operation.
### Method: PUT
>```
>http://localhost:8080/api/tasks/0
>```
### Body (**raw**)

```json
{
    "title" : "the title of your task.",
    "description" : "the description of your task."
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete task under user
### DELETE /api/tasks/{taskId}

This endpoint is used to delete a specific task identified by its unique ID.

#### Request Body

- title (text, optional): The title of the task.
    
- description (text, optional): The description of the task.
    

#### Response

- Status: 401
    
- Content-Type: application/json
    
- msg (string): A message indicating the reason for the unauthorized status.
### Method: DELETE
>```
>http://localhost:8080/api/tasks/66b48726c3abf841cde4098c
>```
### Body (**raw**)

```json
{
    "title" : "please update",
    "description" : "grtdr"
}
```