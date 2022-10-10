# Maintenance Task

This project consists on an API that has two types of users: Managers and Technicians
    - Managers can create and delete users, and these users can also be managers.
    - Managers can see `every other` user's tasks, and are able to delete them.
    - Technicians can create, list and update their own tasks, and `all` managers are notified when a task is created or updated.

As the contents of a task may contain personal information, the task's contents are encrypted alongside the user password

# Endpoints
## Authentication
All endpoints require Basic Auth

## `/users`

### POST `/`: 
Allow a `Manager` to create a user. Body Parameters:
```
{
    "username": string,
    "password": string,
    "firstName": string,
    "lastName": string | null,
    "userRole": enum("MANAGER", "TECHNICIAN"
}
```

Example: 
```
curl --location --request POST 'localhost:3000/users' \
--header 'Authorization: Basic am9obl8xMjM6cGFzc3dvcmRfam9obl8xMjM=' \
--header 'Content-Type: application/json' \
--data-raw '{
"username": "username",
"password": "password",
"firstName": "User",
"lastName": "Name",
"userRole": "MANAGER"
}'
```

### DELETE `/{userID}`: 

Allow a `Manager` to delete a user. Path Parameters: 
```
userID: The user to delete
```

Example:
```
curl --location --request DELETE 'localhost:3000/users/3' \
--header 'Authorization: Basic bWFsdmdvbWVzOnBhc3M='
```

## `/tasks`

### POST `/`:

Allow a `Manager` or `Technician` to create a task. A technician can only see its own tasks, while managers can see every 
user's tasks. Body Parameters:

```
{
    "userId": string,
    "summary": string
}
```

Example: 
```
curl --location --request POST 'localhost:3000/tasks' \
--header 'Authorization: Basic Ym9iXzQ1NjpwYXNzd29yZF9ib2JfNDU2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": 2,
    "summary": "Summary 3"
}'
```

### DELETE `/{userID}`:

Allow a user to delete one of tasks belonging to itself. Path Parameters:
```
ID: The task to delete
```

Example:
```
curl --location --request DELETE 'localhost:3000/tasks/5' \
--header 'Authorization: Basic bWFsdmdvbWVzOnBhc3M='
```

### GET `/{userID}`:

Allow a user to list tasks belonging to the requested user. `Technicians` can only see tasks belonging to itself, while managers 
can see tasks from every other user. Path Parameters:
```
userID: The author of the task
```

Example:
```
curl --location --request GET 'localhost:3000/tasks/2' \
--header 'Authorization: Basic bWFsdmdvbWVzOnBhc3M=' 
```

### PUT `/`:

Allow a user to update one of the tasks belonging to it. Users cannot modify other user's tasks, not even `Managers`. Body 
Parameters:

```
{
    "id": int,
    "userId": int,
    "summary": string,
}
```
Example:
```
curl --location --request PUT 'localhost:3000/tasks' \
--header 'Authorization: Basic Ym9iXzQ1NjpwYXNzd29yZF9ib2JfNDU2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": 2,
    "id": 6,
    "summary": "Summary 3.5"
}'
```

# Technologies required to run the project

- make
- docker-compose
- docker

# Libraries used
- https://github.com/DATA-DOG/go-sqlmock: sql/driver mock used to mock database queries
- https://github.com/go-chi/chi: router used to build the REST API
- https://github.com/go-gorp/gorp: allows the mapping of database columns to struct fields
- https://github.com/go-sql-driver/mysql: driver used in the MariaDB connection
- https://github.com/golang/mock: tool that allows the mocking of interfaces to simulate function calls and return values
- https://github.com/nleof/goyesql: allows mapping of queries on a .sql file to a map
- https://github.com/streadway/amqp: allows rabbitmq integration
- https://github.com/stretchr/testify: allows use of assert in tests

# How to run
- `make up`: starts the application and its dependencies
- `make down`: removes the containers and volumes (Note: this will erase and database changes that you may have made)
- `make test`: run the project's unit tests


- Create users, list tags, and create tasks
- Verify that, when a task is created a job is sent through the RabbitMQ queue. You can check that through the logs, or at 
  http://localhost:15672/#/ (user: `guest`, password: `guest`)
