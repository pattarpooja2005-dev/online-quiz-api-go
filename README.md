# Online Quiz Platform API (Golang)

## 📌 Overview

This project is a REST API for an online quiz system built using **Go, Gin, and GORM**.
It allows teachers to create quizzes and questions, and students to attempt quizzes and receive automatic scores.

---

## ✨ Features

* Create quizzes
* Add questions to quizzes
* Start quiz attempts
* Submit answers
* Automatic score calculation
* Attempt tracking with start and end time

---

## 🛠 Tech Stack

* Go (Golang)
* Gin Web Framework
* GORM ORM
* SQLite Database

---

## 🗄 Database Design

### Tables

* Quiz
* Question
* Attempt
* Answer

### Relationships

* One Quiz → Many Questions
* One Attempt → Many Answers

---

## 🔗 API Endpoints

### Create Quiz

POST `/quiz`

Example request:

```json
{
  "title": "Science Quiz",
  "timeLimitMinutes": 15
}
```

---

### Add Question

POST `/question`

---

### Start Attempt

POST `/start`

---

### Submit Answers

POST `/submit`

---

## 🧮 Scoring Logic

Each submitted answer is compared with the correct option.
Score = number of correct answers.

---

## ▶ How to Run the Project

1. Install Go
2. Clone repository
3. Run server:

```
go run main.go
```

4. Test APIs using Postman

---

## 📁 Project Structure

```
config/
models/
handlers/
routes/
docs/
main.go
```

---

## 🚀 Future Improvements

* Authentication system
* Quiz timer enforcement
* Leaderboard
* Randomized questions
* Web frontend integration
