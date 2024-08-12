# Project: clean_taskManager

## End-point: Register
### Register User

This endpoint allows the client to register a new user by submitting the required user details.

#### Request Body

- `name` (string, required): The name of the user.
    
- `email` (string, required): The email address of the user.
    
- `password` (string, required): The password for the user account.
    
- `isadmin` (boolean, required): Indicates whether the user has admin privileges.
    

#### Response Body

The response will include the status of the registration process, along with any relevant error messages.
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
### Login API

This API endpoint is used to authenticate and login a user.

#### Request Body

- `email` (string, required): The email of the user.
    
- `password` (string, required): The password of the user.
    

#### Response

The response will contain the authentication token for the logged in user.
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

## End-point: Delete Task
This endpoint sends an HTTP DELETE request to the specified URL to delete a task with the given ID. The request payload includes the title and description of the task to be deleted.

### Response

The response of this request is a JSON schema.
### Method: DELETE
>```
>http://localhost:8080/api/tasks/66b9126031279e9a607cda56
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

## End-point: Get Tasks
This endpoint makes an HTTP GET request to retrieve a list of tasks from the server. The request does not require any parameters or a request body. The response will include a list of tasks, each with a title and description. The response may also include additional information such as task IDs, timestamps, or other metadata related to the tasks.
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

## End-point: Create Task
### Create a New Task

This endpoint allows you to create a new task.

#### Request Body

- `title` (string, required): The title of the task.
    
- `description` (string, required): The description of the task.
    

#### Response

The response is a JSON object with the following schema:

``` json
{
  "type": "object",
  "properties": {
    "taskId": {
      "type": "string"
    },
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
    "title" : "blah blah",
    "description" : "the description of your task"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{jwt_token}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update Task
This endpoint makes an HTTP PUT request to update a specific task identified by its ID. The request should be sent to [http://localhost:8080/api/tasks/66b9131631279e9a607cda58](http://localhost:8080/api/tasks/66b9131631279e9a607cda58) with a JSON payload in the raw request body. The payload should include the keys "title" and "description" to update the task with the provided values.

### Response

The response of this request is a JSON schema representing the structure of the response data. This schema can be used to validate and understand the format of the response returned by the API.
### Method: PUT
>```
>http://localhost:8080/api/tasks/66b9131631279e9a607cda58
>```
### Body (**raw**)

```json
{
    "title" : "blah update",
    "description" : "the description of your task"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{jwt_token}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get Task By Id 
This endpoint retrieves a specific task identified by the provided task ID. The request should be made using an HTTP GET method to the specified URL.

### Response

The response of this request is a JSON object representing the task. Below is the JSON schema for the response:

``` json
{
  "type": "object",
  "properties": {
    "taskId": {
      "type": "string"
    },
    "title": {
      "type": "string"
    },
    "description": {
      "type": "string"
    },
    "status": {
      "type": "string"
    },
    "createdAt": {
      "type": "string",
      "format": "date-time"
    },
    "updatedAt": {
      "type": "string",
      "format": "date-time"
    }
  }
}

 ```
### Method: GET
>```
>http://localhost:8080/api/tasks/66b9131631279e9a607cda58
>```
### Body (**raw**)

```json
{
    "title" : "blah update",
    "description" : "the description of your task"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{jwt_token}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get User By Id 
This endpoint makes an HTTP GET request to retrieve user information with the specified ID. The request should include the user ID in the URL path. The request body should contain the user's title and description.

The response will include the user's details, such as their ID, name, email, and other relevant information.
### Method: GET
>```
>http://localhost:8080/api/users/66b9108a31279e9a607cda55
>```
### Body (**raw**)

```json
{
    "title" : "blah update",
    "description" : "the description of your task"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{jwt_token}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get Users
This endpoint makes an HTTP GET request to retrieve a list of users from the specified API. The request does not include any query parameters or request body.

### Response

The response of this request is a JSON schema representing the structure of the data returned by the API. The schema will define the properties and their data types for the user data.
### Method: GET
>```
>http://localhost:8080/api/users
>```
### Body (**raw**)

```json
{
    "title" : "blah update",
    "description" : "the description of your task"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{jwt_token}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete User
### Delete User

This endpoint sends an HTTP DELETE request to delete a specific user with the given ID.

#### Request Body

- `title` (string, optional): The title of the user.
    
- `description` (string, optional): The description of the user.
    

#### Response

The response will indicate the success or failure of the deletion operation.
### Method: DELETE
>```
>http://localhost:8080/api/users/66b9108a31279e9a607cda55
>```
### Body (**raw**)

```json
{
    "title" : "blah update",
    "description" : "the description of your task"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{jwt_token}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update User
### Update User Details

This endpoint is used to update the details of a specific user.

#### Request Body

- `password` (string, required): The new password for the user.
    

#### Response

The response will contain the updated details of the user.
### Method: PUT
>```
>http://localhost:8080/api/users/66b9153631279e9a607cda59
>```
### Body (**raw**)

```json
{
    "password": "new99"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{jwt_token}}|string|

