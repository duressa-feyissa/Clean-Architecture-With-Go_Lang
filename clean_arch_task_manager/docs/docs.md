# Task Manager API Documentation

## Overview
This document outlines the API endpoints provided by the Task Manager system, allowing for the management of tasks within a project management context. The system supports operations such as creating, updating, retrieving, and deleting tasks. You can find postman docs [here](https://documenter.getpostman.com/view/30253109/2sA3rxruEy)

## API Endpoints

### Get All Tasks
- **Endpoint**: `GET /tasks`
- **Description**: Retrieves a list of all tasks.
- **Response**: An array of task objects.

### Get Task by ID
- **Endpoint**: `GET /tasks/:id`
- **Description**: Retrieves a task by its unique identifier.
- **Parameters**:
  - `id`: The unique identifier of the task.
- **Response**: A task object.

### Add a New Task
- **Endpoint**: `POST /tasks`
- **Description**: Adds a new task to the system.
- **Body**:
  - `title`: The title of the task.
  - `description`: The description of the task.
  - `status`: The current status of the task.
- **Response**: A message indicating success or failure.

### Update a Task
- **Endpoint**: `PUT /tasks/:id`
- **Description**: Updates an existing task.
- **Parameters**:
  - `id`: The unique identifier of the task to update.
- **Body**:
  - `title`: The new title of the task (optional).
  - `description`: The new description of the task (optional).
  - `status`: The new status of the task (optional).
- **Response**: A message indicating success or failure.

### Delete a Task
- **Endpoint**: `DELETE /tasks/:id`
- **Description**: Deletes a task from the system.
- **Parameters**:
  - `id`: The unique identifier of the task to delete.
- **Response**: A message indicating success or failure.

### Register User
- **Endpoint**: `POST /register`
- **Description**: Registers a new user.
- **Body**:
  - `username`: The username of the user.
  - `password`: The password of the user.
- **Response**: A message indicating success or failure.

### Login User
- **Endpoint**: `POST /login`
- **Description**: Logs in a user.
- **Body**:
  - `username`: The username of the user.
  - `password`: The password of the user.
- **Response**: A message indicating success or failure.

## Models
### Task
- **Fields**:
  - `ID`: Unique identifier for the task.
  - `Title`: Title of the task.
  - `Description`: Description of the task.
  - `DueDate`: Due date of the task.
  - `Status`: Current status of the task.
  - `UserID` : User ID associated with the task.

### User
- **Fields**:
  - `ID`: Unique identifier for the user.
  - `Username`: Username of the user.
  - `Password`: Password of the user.
  - `Role`: Role of the user.


## Error Handling
All endpoints return appropriate HTTP status codes along with error messages in the case of failure.

