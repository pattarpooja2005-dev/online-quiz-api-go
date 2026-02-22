🎓 Online Quiz Platform — Timed Assessment System

A full-stack quiz and examination platform built with Go (Gin + GORM) backend and a lightweight web interface for demonstration.

The system supports:

✔ timed quiz sessions
✔ automatic grading
✔ concurrent attempts
✔ secure answer submission

Designed to simulate real exam platforms like Google Forms (quiz mode) and Kahoot.

📑 Table of Contents

Architecture Overview

Tech Stack

Database Schema

REST API Documentation

System Workflow (End-to-End Execution) ⭐

Timer Implementation Strategy

Anti-Cheating Mechanisms

Concurrency Handling

Auto-Grading Logic

Example Quiz Flow

Setup & Running Instructions

Project Structure

🏗 Architecture Overview

The system follows a layered backend architecture with clear separation of responsibilities.

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
Layer Responsibilities
Layer	Responsibility
Handlers	HTTP request validation & response formatting
Services	Quiz logic, timer validation, scoring
Repository	Database operations
Models	Data structures
Config	DB connection & migration
🔧 Tech Stack
Backend
Component	Technology	Purpose
Language	Go	Performance & concurrency
Framework	Gin	HTTP routing
ORM	GORM	Database access
Database	SQLite	Embedded storage
CORS	gin-contrib/cors	Cross-origin support
Frontend (Demo UI)
Technology	Purpose
HTML + CSS + JS	Lightweight interface
Fetch API	HTTP requests
🗄 Database Schema
Tables
Quiz

id

title

time_limit_minutes

created_at

Question

id

quiz_id

text

option_a

option_b

option_c

option_d

correct_option

Attempt

id

quiz_id

user_name

start_time

end_time

score

Answer

id

attempt_id

question_id

selected_option

📡 REST API Documentation

Base URL:

http://localhost:8080
Method	Endpoint	Description
POST	/quiz	Create quiz
POST	/question	Add question
POST	/start	Start attempt
POST	/submit	Submit answers
🔄 System Workflow (End-to-End Execution)

This explains how the system works from server startup to final result.

1️⃣ Server Startup

When application starts:

✔ Database connection established
✔ Tables auto-migrated (Quiz, Question, Attempt, Answer)
✔ HTTP server starts on port 8080
✔ Templates and static files loaded

Access system:

http://localhost:8080
2️⃣ Start Quiz Page

User enters:

Quiz ID

User Name

Frontend sends:

POST /start
3️⃣ Attempt Creation

Server:

✔ validates quiz exists
✔ checks no active attempt
✔ records start time
✔ calculates deadline
✔ stores attempt

Timer formula:

EndTime = StartTime + QuizTimeLimit

User redirected to quiz page.

4️⃣ Quiz Attempt Page

Displays:

✔ countdown timer
✔ question
✔ answer selection

Timer shown:

RemainingTime = EndTime − CurrentServerTime

Frontend timer is informational only.
Backend enforces time.

5️⃣ Answer Submission

Frontend sends:

POST /submit

Includes:

attempt ID

answers

6️⃣ Server Validation

Server checks:

✔ attempt exists
✔ not already submitted
✔ within deadline

if CurrentTime > EndTime → reject
7️⃣ Automatic Scoring

Server:

✔ stores answers
✔ compares with correct options
✔ calculates score
✔ locks attempt

Score = Number of correct answers
8️⃣ Result Display

Server returns score → result page shown.

Attempt permanently closed.

9️⃣ Page Refresh Handling

If refreshed:

✔ timer unchanged
✔ attempt continues
✔ cannot restart

🔟 Concurrent Users

Multiple students supported because:

✔ separate attempt records
✔ independent timers
✔ database isolation

1️⃣1️⃣ Anti-Cheating Enforcement

Prevents:

✔ restart attempt
✔ late submission
✔ multiple submissions
✔ timer manipulation
✔ device clock change

1️⃣2️⃣ Attempt Lifecycle
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

✔ strict time enforcement
✔ server authoritative timer
✔ persistent tracking
✔ secure validation
✔ concurrent support

⏱ Timer Implementation Strategy

Each quiz has a time limit enforced server-side.

elapsed = now - start_time

if elapsed > time_limit
    reject submission
else
    accept

Server time is authoritative.

🔒 Anti-Cheating Mechanisms
Method	Protection
Server timer	prevents extended attempts
Hidden answers	prevents solution fetching
Attempt locking	prevents resubmission
Time validation	prevents late submission
⚡ Concurrency Handling

Handled using:

✔ independent attempts
✔ stateless handlers
✔ DB isolation
✔ atomic scoring

Each request processed independently.

🧮 Auto-Grading Logic

For each answer:

fetch correct option

compare

increment score

Time complexity:

O(n)
📊 Example Quiz Flow

Teacher → creates quiz → adds questions

Student → starts attempt → timer begins

Student submits → system:

✔ validates time
✔ calculates score
✔ stores result

🚀 Setup & Running
Backend
go mod tidy
go run main.go

Server:

http://localhost:8080
Demo UI

Open:

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

