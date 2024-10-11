# api_entities_management

API to manage entities into a database (students in this case)

Routes:
- GET /students - List all students
- POST /students - Create students - OK
- GET /students/:id - Get info from a specific student
- PUT /students/:id - Update student
- DELETE /students/:id - Delete student
- GET /students?active=<true/false> - List all active/non-active students

Struct Student:
- Name 
- CPF
- Email
- Age 
- Active


