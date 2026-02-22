# 🎓 Online Quiz Platform — Timed Assessment System

A full-stack quiz and examination platform built with **Go (Gin + GORM)** backend and a lightweight web interface for demonstration.
The system supports timed quiz sessions, automatic grading, concurrent attempts, and secure answer submission.

Designed to simulate real exam platforms like **Google Forms (quiz mode)** and **Kahoot**.

---

# 📑 Table of Contents

* Architecture Overview
* Tech Stack
* Database Schema
* REST API Documentation
* Timer Implementation Strategy
* Anti-Cheating Mechanisms
* Concurrency Handling
* Auto-Grading Logic
* Example Quiz Flow
* Setup & Running Instructions
* Project Structure

---

# 🏗 Architecture Overview

The system follows a layered backend architecture with clear separation of responsibilities.

```
Client (Browser / Postman / Demo UI)
                │
                │ HTTP / JSON
                ▼
┌──────────────────────────────────────┐
│              Go Backend              │
│                                      │
│  ┌──────────┐   ┌──────────┐        │
│  │ Handlers │──▶│ Services │        │
│  │  (HTTP)  │   │ (Logic)  │        │
│  └──────────┘   └────┬─────┘        │
│                       │              │
│                 ┌─────▼─────┐       │
│                 │ Repository │       │
│                 │ (Database) │       │
│                 └─────┬─────┘       │
└───────────────────────┼─────────────┘
                        │ GORM ORM
                 ┌──────▼──────┐
                 │   SQLite     │
                 │   Database   │
                 └──────────────┘
```

### Layer Responsibilities

| Layer      | Responsibility                                  |
| ---------- | ----------------------------------------------- |
| Handlers   | HTTP request validation and response formatting |
| Services   | Quiz logic, timer checks, scoring               |
| Repository | Database operations                             |
| Models     | Data structures                                 |
| Config     | DB connection & migration                       |

---

# 🔧 Tech Stack

### Backend

| Component | Technology       | Purpose                   |
| --------- | ---------------- | ------------------------- |
| Language  | Go               | Performance & concurrency |
| Framework | Gin              | HTTP routing              |
| ORM       | GORM             | Database access           |
| Database  | SQLite           | Embedded storage          |
| CORS      | gin-contrib/cors | Cross origin support      |

### Frontend (Demo UI)

| Technology      | Purpose               |
| --------------- | --------------------- |
| HTML + CSS + JS | Lightweight interface |
| Fetch API       | HTTP requests         |

---

# 🗄 Database Schema

## Tables

### Quiz

```
id
title
time_limit_minutes
created_at
```

### Question

```
id
quiz_id
text
option_a
option_b
option_c
option_d
correct_option
```

### Attempt

```
id
quiz_id
user_name
start_time
end_time
score
```

### Answer

```
id
attempt_id
question_id
selected_option
```

---

# 📡 REST API Documentation

Base URL:

```
http://localhost:8080
```

| Method | Endpoint  | Description    |
| ------ | --------- | -------------- |
| POST   | /quiz     | Create quiz    |
| POST   | /question | Add question   |
| POST   | /start    | Start attempt  |
| POST   | /submit   | Submit answers |

---

# ⏱ Timer Implementation Strategy

Each quiz has a time limit.
Timer is enforced server-side.

### How it works

1. Attempt start time stored in DB
2. Time limit fetched from quiz
3. On submission:

```
elapsed = now - start_time

if elapsed > time_limit
    reject submission
else
    accept and score
```

### Why server-side timer?

Client timers can be manipulated.
Server ensures authoritative time validation.

---

# 🔒 Anti-Cheating Mechanisms

| Method                    | Protection                      |
| ------------------------- | ------------------------------- |
| Server-side timer         | Prevents extended attempts      |
| No correct answers in API | Students cannot fetch solutions |
| Attempt locking           | Prevents multiple submissions   |
| Time validation on submit | Prevents late submission        |

---

# ⚡ Concurrency Handling

Multiple students can attempt simultaneously.

Handled using:

* Independent attempt records
* Stateless HTTP handlers
* Database isolation
* Atomic scoring update

Each request processed independently.

---

# 🧮 Auto-Grading Logic

For each submitted answer:

```
fetch correct option
compare with selected
increment score if match
```

Time complexity: O(n)

Where n = number of answers.

---

# 📊 Example Quiz Flow

Teacher creates quiz → adds questions.

Student starts attempt → timer begins.

Student submits answers → system:

1. Validates time
2. Calculates score
3. Stores attempt
4. Returns result

---

# 🚀 Setup & Running

## Backend

```
go mod tidy
go run main.go
```

Server:

```
http://localhost:8080
```

## Demo UI

Open:

```
frontend-html/index.html
```

---

# 📁 Project Structure

```
quiz-api/
├── main.go
├── config/
├── models/
├── handlers/
├── routes/
├── docs/
├── frontend-html/
└── README.md
```

---
# Complete System Workflow
This section explains how the quiz system operates from the user interface to backend processing.

1️⃣ Server Startup

When the application starts:

Database connection is established

Tables are auto-migrated (Quiz, Question, Attempt, Answer)

HTTP server starts on port 8080

HTML templates and static files are loaded

Users can access the system at:

http://localhost:8080
2️⃣ Start Quiz Page (Home Page)

When the user opens the application, the start page is displayed.

The user must enter:

Quiz ID

User Name

When the Start Quiz button is clicked:

Frontend sends request:

POST /start

Request body:

{
  "quizID": 1,
  "userName": "Student Name"
}
3️⃣ Attempt Creation (Backend Processing)

The server performs the following steps:

Validates quiz exists

Checks if user already has active attempt

Records start time

Calculates end time using quiz duration

Stores attempt in database

Timer calculation:

EndTime = StartTime + QuizTimeLimit

Attempt record is saved with:

start time

deadline

submission status

User is redirected to quiz page.

4️⃣ Quiz Attempt Page

The quiz page displays:

countdown timer

question input

answer selection

The timer shown in UI is based on remaining time:

RemainingTime = EndTime − CurrentServerTime

Important:

The frontend timer is only for display.
Actual enforcement is done on the server.

5️⃣ Answer Submission

When user submits answers:

Frontend sends request:

POST /submit

Request contains:

attempt ID

question ID

selected option

6️⃣ Server Validation (Critical Step)

Before evaluating answers, the server performs security checks:

✔ Attempt exists
✔ Attempt not already submitted
✔ Current time is before deadline

Time validation:

if CurrentTime > EndTime → submission rejected

If time expired:

attempt is closed

submission blocked

This ensures strict timer enforcement.

7️⃣ Automatic Scoring

If submission is valid:

Answers are stored

Each answer is compared with correct option

Score is calculated

Attempt marked as submitted

Score formula:

Score = Number of correct answers
8️⃣ Result Display

Server returns final score.

Frontend redirects to result page showing:

total score

completion status

Attempt is permanently closed.

9️⃣ Handling Page Refresh

If user refreshes during quiz:

attempt remains active

start time unchanged

deadline unchanged

timer continues

User cannot restart quiz to gain extra time.

🔟 Concurrent Users

Multiple students can take quiz simultaneously because:

each attempt stored separately

each has independent timer

database isolates records

No shared timer exists.

1️⃣1️⃣ Anti-Cheating Mechanisms

The system prevents:

✔ restarting attempt
✔ submitting after deadline
✔ multiple submissions
✔ timer manipulation
✔ changing device clock

All timing decisions are server controlled.

1️⃣2️⃣ Attempt Lifecycle Summary
Start Quiz
   ↓
Attempt Created
   ↓
Timer Running
   ↓
Submit Answers
   ↓
Server Time Validation
   ↓
Score Calculation
   ↓
Attempt Locked
✅ System Guarantees

strict time enforcement

server authoritative timer

persistent attempt tracking

secure submission validation

concurrent user support

# 🎯 Key Features

✔ Timed quiz sessions
✔ Automatic grading
✔ Concurrent attempts
✔ Secure answer submission
✔ RESTful architecture
✔ Database persistence

---

# 📝 License

Academic project.
