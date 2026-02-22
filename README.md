# 🎓 Online Quiz Platform — Timed Assessment System

A full-stack quiz and examination platform built with **Go (Gin + GORM)** backend and a lightweight web interface for demonstration.

The system supports:

✔ Timed quiz sessions  
✔ Automatic grading  
✔ Concurrent attempts  
✔ Secure answer submission  

Designed to simulate real exam platforms like **Google Forms (Quiz Mode)** and **Kahoot**.

---

# 📑 Table of Contents

1. Architecture Overview  
2. Tech Stack  
3. Database Schema  
4. REST API Documentation  
5. Timer Implementation Strategy  
6. Anti-Cheating Mechanisms  
7. Concurrency Handling  
8. Auto-Grading Logic  
9. Example Quiz Flow  
10. API Request Examples  
11. Setup & Running Instructions  
12. Project Structure  
13. Key Features  
14. License  

---

# 🏗 Architecture Overview

The system follows a layered backend architecture with clear separation of responsibilities.

Client (Browser / Postman / Demo UI)
│
│ HTTP / JSON
▼
┌──────────────────────────────────────┐
│ Go Backend │
│ │
│ ┌──────────┐ ┌──────────┐ │
│ │ Handlers │──▶│ Services │ │
│ │ (HTTP) │ │ (Logic) │ │
│ └──────────┘ └────┬─────┘ │
│ │ │
│ ┌─────▼─────┐ │
│ │ Repository │ │
│ │ (Database) │ │
│ └─────┬─────┘ │
└───────────────────────┼─────────────┘
│ GORM ORM
┌──────▼──────┐
│ SQLite │
│ Database │
└──────────────┘

## Layer Responsibilities

| Layer | Responsibility |
|---|---|
| Handlers | HTTP request validation and response formatting |
| Services | Quiz logic, timer checks, scoring |
| Repository | Database operations |
| Models | Data structures |
| Config | Database connection & migration |

---

# 🔧 Tech Stack

## Backend

| Component | Technology | Purpose |
|---|---|---|
| Language | Go | Performance & concurrency |
| Framework | Gin | HTTP routing |
| ORM | GORM | Database access |
| Database | SQLite | Embedded storage |
| CORS | gin-contrib/cors | Cross origin support |

## Frontend (Demo UI)

| Technology | Purpose |
|---|---|
| HTML + CSS + JS | Lightweight interface |
| Fetch API | HTTP requests |

---

# 🗄 Database Schema

## Quiz
- id  
- title  
- time_limit_minutes  
- created_at  

## Question
- id  
- quiz_id  
- text  
- option_a  
- option_b  
- option_c  
- option_d  
- correct_option  

## Attempt
- id  
- quiz_id  
- user_name  
- start_time  
- end_time  
- score  

## Answer
- id  
- attempt_id  
- question_id  
- selected_option  

---

# 📡 REST API Documentation

Base URL:
http://localhost:8080

| Method | Endpoint | Description |
|---|---|---|
| POST | /quiz | Create quiz |
| POST | /question | Add question |
| POST | /start | Start attempt |
| POST | /submit | Submit answers |

---

# ⏱ Timer Implementation Strategy

Each quiz has a time limit enforced server-side.

## How it works

1. Attempt start time stored in database  
2. Time limit fetched from quiz  
3. On submission system calculates elapsed time 

elapsed = now - start_time

if elapsed > time_limit
reject submission
else
accept and score


## Why server-side timer?

Client timers can be manipulated.  
Server ensures authoritative time validation.

---

# 🔒 Anti-Cheating Mechanisms

| Method | Protection |
|---|---|
| Server-side timer | Prevents extended attempts |
| No correct answers in API | Students cannot fetch solutions |
| Attempt locking | Prevents multiple submissions |
| Time validation on submit | Prevents late submission |

---

# ⚡ Concurrency Handling

Multiple students can attempt simultaneously.

Handled using:

- Independent attempt records  
- Stateless HTTP handlers  
- Database isolation  
- Atomic scoring updates  

Each request processed independently.

---

# 🧮 Auto-Grading Logic

For each submitted answer:

1. Fetch correct option  
2. Compare with selected option  
3. Increment score if match  

Time Complexity:
O(n)

Where n = number of answers.

---

# 📊 Example Quiz Flow

1. Teacher creates quiz → adds questions  
2. Student starts attempt → timer begins  
3. Student submits answers → system:

✔ validates time  
✔ calculates score  
✔ stores attempt  
✔ returns result  

---

# 📮 API Request Examples

## Create Quiz
POST /quiz

```json
{
  "title": "Math Quiz",
  "time_limit_minutes": 10
}
POST /question
{
  "quiz_id": 1,
  "text": "2 + 2 = ?",
  "option_a": "3",
  "option_b": "4",
  "option_c": "5",
  "option_d": "6",
  "correct_option": "B"
}
Start Attempt
POST /start
{
  "quiz_id": 1
}
Submit Answers
POST /submit
{
  "attempt_id": 1,
  "answers": [
    {
      "question_id": 1,
      "selected": "B"
    }
  ]
}
🚀 Setup & Running Instructions
Backend
go mod tidy
go run main.go

Server runs at:

http://localhost:8080
Demo UI

Open in browser:

frontend-html/index.html
📁 Project Structure
quiz-api/
├── main.go
├── config/
├── models/
├── handlers/
├── routes/
├── docs/
├── frontend-html/
└── README.md
🎯 Key Features

✔ Timed quiz sessions
✔ Automatic grading
✔ Concurrent attempts
✔ Secure answer submission
✔ RESTful architecture
✔ Database persistence

📝 License

Academic project.