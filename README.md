# Student Management API (Study Project)

This is a simple RESTful API built for study purposes. It manages a collection of **students** stored in a database and allows basic CRUD operations using standard HTTP methods.

## 📚 Purpose

The goal of this project is to practice and understand the fundamentals of designing and implementing a REST API.

## 📌 Features

- List all students
- Create new students
- View a single student by ID
- Update student information
- Delete a student
- Filter students by active status

## 🧱 Entity: Student

Each student has the following structure:

```json
{
  "name": "string",
  "cpf": "string",
  "email": "string",
  "age": number,
  "active": boolean
}
```

## 🚀 API Endpoints

| Method | Endpoint                    | Description                        |
|--------|-----------------------------|------------------------------------|
| GET    | `/students`                 | List all students                  |
| POST   | `/students`                 | Create a new student               |
| GET    | `/students/:id`             | Get a student by ID                |
| PUT    | `/students/:id`             | Update a student's information     |
| DELETE | `/students/:id`             | Delete a student                   |
| GET    | `/students?active=true`     | List active students               |
| GET    | `/students?active=false`    | List inactive students             |

## 🛠️ Technologies Used

- **RESTful API architecture** – Follows standard HTTP methods for CRUD operations
- **Go (Golang)** – Programming language used to build the API
- **Echo** – High performance, minimalist web framework for Go
- **GORM** – ORM library for Go, used for interacting with the SQLite database
- **SQLite** – Lightweight relational database used for local data persistence

## 📎 Note

This is not a production-ready project. It was developed solely for educational purposes and hands-on experience with RESTful API design.

