# Standart Repository Pattern with Dependency Injection
Can be used as template

## Folder Design

~~~~
├── cmd
│   └── main
├── pkg
│   └── controller
│   └── database
│   └── dto
│   └── entity
│   └── repository
│   └── server
│   └── service
│   └── util
│ go.mod
~~~~

### Sample Response
```go
{
    "success": true,
    "data": [
        {
            "id": "8be3f486-d3f7-4f61-9859-7c85cf425d43",
            "task": "do ...",
            "todo_date": "2022-12-31T10:30:00Z",
            "create_date": "2022-12-31T00:01:02.980624397+03:00"
        },
        {
            "id": "832d375d-980b-45be-8f85-87cc482ff412",
            "task": "do smth",
            "todo_date": "2022-12-31T10:45:00Z",
            "create_date": "2022-12-31T00:01:25.073548071+03:00"
        }
    ],
    "status_code": 200
}
```

### Sample Error Response
```go
{
    "message": "record not found",
    "status_code": 400
}
```

