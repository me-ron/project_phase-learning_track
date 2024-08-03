# Project: task_manager

## End-point: http://localhost:8080/api/tasks
This endpoint allows the client to create a new task by sending a POST request to the specified URL.

### Request Body

- The request should include a JSON object in the body with the following parameters:
    
    - `title` (string): The title of the task.
        
    - `description` (string): The description of the task.
        

### Response

The response to the request is a JSON array with the following schema:

``` json
[
    {
        "id": "",
        "title": "",
        "description": "",
        "due_date": ""
    }
]

 ```

The response contains an array of task objects, where each task object includes the following properties:

- `id` (string): The unique identifier of the task.
    
- `title` (string): The title of the task.
    
- `description` (string): The description of the task.
    
- `due_date` (string): The due date of the task.
### Method: POST
>```
>http://localhost:8080/api/tasks
>```
### Body (**raw**)

```json
{
    "title" : "the title of your task",
    "description" : "the description of your task"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: http://localhost:8080/api/tasks
This endpoint makes an HTTP GET request to retrieve a list of tasks from the server. The request does not require any parameters. The response is in JSON format and returns an array of task objects, each containing the following properties:

- "id" (string): The unique identifier of the task.
    
- "title" (string): The title of the task.
    
- "description" (string): The description of the task.
    
- "due_date" (string): The due date of the task.
    

Here is the JSON schema for the response:

``` json
{
  "type": "array",
  "items": {
    "type": "object",
    "properties": {
      "id": {
        "type": "string"
      },
      "title": {
        "type": "string"
      },
      "description": {
        "type": "string"
      },
      "due_date": {
        "type": "string"
      }
    }
  }
}

 ```
### Method: GET
>```
>http://localhost:8080/api/tasks
>```
### Body (**raw**)

```json
{
    "title" : "the title of your task",
    "description" : "the description of your task"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: http://localhost:8080/api/tasks
This endpoint retrieves a specific task with the ID "0" from the server.

### Request

The request does not require any parameters in the URL.

### Response

The response will have a status code of 200 and a content type of "application/json". The response body will be a JSON object with the following schema:

``` json
{
    "id": "",
    "title": "",
    "description": "",
    "due_date": ""
}

 ```

The "id" represents the unique identifier of the task, "title" represents the title of the task, "description" represents the description of the task, and "due_date" represents the due date of the task.
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

## End-point: http://localhost:8080/api/tasks
### Update Task

This endpoint is used to update a specific task by its ID.

#### Request

- Method: PUT
    
- URL: `http://localhost:8080/api/tasks/0`
    
- Body:
    
    - title (text, optional): The title of the task.
        
    - description (text, optional): The description of the task.
        

#### Response

The response is in JSON format and follows the schema below:

``` json
{
  "type": "object",
  "properties": {
    "id": {
      "type": "string"
    },
    "title": {
      "type": "string"
    },
    "description": {
      "type": "string"
    },
    "due_date": {
      "type": "string"
    }
  }
}

 ```
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

## End-point: http://localhost:8080/api/tasks
This endpoint sends an HTTP DELETE request to delete a task with the specified ID. The request should be sent to [http://localhost:8080/api/tasks/0](http://localhost:8080/api/tasks/0).

### Response

The response to this request is in the JSON format with a status code of 200. The response schema is as follows:

``` json
{
    "type": "object",
    "properties": {
        "messages": {
            "type": "string"
        }
    }
}

 ```
### Method: DELETE
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
_________________________________________________
Powered By: [postman-to-markdown](https://github.com/bautistaj/postman-to-markdown/)
