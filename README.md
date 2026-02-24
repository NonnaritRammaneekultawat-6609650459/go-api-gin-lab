# 🎓 Student API with Layered Architecture (Lab 4)

This project is an enhanced REST API for student management, built using the **Gin Framework** and **SQLite**. It follows the **Layered Architecture** pattern (Handler, Service, Repository, Model) to ensure clean code and maintainability.

## 👤 Author
- **Name:** Nonnarit Rammaneekultawat
- **Student ID:** 6609650459

---

## 🚀 How to Run

### 1. Prerequisites
- Go installed on your machine.
- SQLite3 (The database file `students.db` will be created automatically).

### 2. Setup & Run
```bash
# Clone the repository
git clone <your-repository-url>
cd go-api-gin-lab

# Install dependencies
go mod tidy

# Run the application
go run main.go

#The server will start at http://localhost:8080

## API Endpoints
Method,Endpoint,Description,Status Code (Success)
GET,/students,Get all students,200 OK
GET,/students/:id,Get student by ID,200 OK
POST,/students,Create a new student,201 Created
PUT,/students/:id,Update student information,200 OK
DELETE,/students/:id,Delete a student by ID,204 No Content

### TEST Examples

1. Create a Student (POST)
Endpoint: POST http://localhost:8080/students
Purpose: To add initial data before testing other functions.
Request JSON Body:
{
  "id": "6609650459",
  "name": "Nonnarit Rammaneekultawat",
  "major": "Computer Science",
  "gpa": 3.16
}
Expected Response: 201 Created

2. Update Student Information (PUT)
Endpoint: PUT http://localhost:8080/students/6609650459
Purpose: To test Update and Validation.
Request JSON Body:
{
  "id": "6609650459",
  "name": "Nonnarit (Updated)",
  "major": "Software Engineering",
  "gpa": 3.85
}
Expected Response: 200 OK

3. Test Input Validation (Bad Request)
Endpoint: PUT http://localhost:8080/students/6609650459
Purpose: To verify that the API rejects invalid data (GPA > 4.00)
Request JSON Body:
{
  "id": "6609650459",
  "name": "Nonnarit (Updated)",
  "major": "Software Engineering",
  "gpa": 5.00
}
Expected Response: 400 Bad Request
Response JSON:
{
  "error": "GPA must be between 0.00 and 4.00"
}

4. Delete Student (DELETE)
Endpoint: DELETE http://localhost:8080/students/6609650459
Purpose: To test Delete
Expected Response: 204 No Content

5. Test Not Found (404 Error)
Endpoint: GET http://localhost:8080/students/6609650459 (After deletion)
Purpose: To verify that the student was successfully removed and the API handles missing data.
Expected Response: 404 Not Found
Response JSON:
{
  "error": "Student not found"
}
