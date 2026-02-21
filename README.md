# 🎓 Online Quiz Platform API (Golang)

A production-ready REST API built in Go for managing quizzes, questions, attempts, and automatic scoring.

The system supports full quiz lifecycle:
create quiz → add questions → start attempt → submit answers → auto score.

---

## 🏗️ Architecture

This project follows a **layered modular architecture** with clearly separated responsibilities.

```
quiz-api/
├── main.go               # Entry point (server startup)
├── config/database.go    # Database connection & migration
├── models/               # Database schema definitions
├── handlers/             # API request handling logic
├── routes/               # Route registration
├── docs/                 # Design & AI prompts
└── README.md
```

### Dependency Flow

main.go → routes → handlers → models → database

Each layer handles a specific responsibility.

---

## 🚀 Setup Instructions

### Prerequisites

* Go installed
* Git installed

### Run the Project

```
go run main.go
```

Server runs on:

```
http://localhost:8080
```

Database tables are auto-created on startup.

---

## 📡 API Documentation

### Base URL

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

### Create Quiz

POST `/quiz`

```
{
  "title": "Science Quiz",
  "timeLimitMinutes": 15
}
```

---

### Add Question

POST `/question`

```
{
  "quizID": 1,
  "text": "Earth is a?",
  "optionA": "Star",
  "optionB": "Planet",
  "optionC": "Moon",
  "optionD": "Galaxy",
  "correct": "B"
}
```

---

### Start Attempt

POST `/start`

```
{
  "quizID": 1,
  "userName": "Pooja"
}
```

---

### Submit Answers

POST `/submit`

```
[
  {
    "attemptID": 1,
    "questionID": 1,
    "selected": "B"
  }
]
```

Response:

```
{
  "score": 1
}
```

---

## 🧮 Scoring Logic

Each submitted answer is compared with the stored correct option.

Algorithm:

1. Retrieve correct answer from database
2. Compare with selected answer
3. Increment score if matched
4. Store attempt end time
5. Return total score

Time complexity:
O(n) where n = number of answers.

---

## 🗄 Database Design

Tables:

* Quiz
* Question
* Attempt
* Answer

Relationships:

Quiz → Questions (one-to-many)
Attempt → Answers (one-to-many)

---

## 🔒 Data Consistency

* All answers stored per attempt
* Score computed only after submission
* Attempt start and end times tracked

This ensures complete attempt history.

---

## 🛠️ Tech Stack

| Technology | Purpose          |
| ---------- | ---------------- |
| Go         | Backend language |
| Gin        | HTTP routing     |
| GORM       | ORM              |
| SQLite     | Database         |

---

## 📄 Documentation

docs/design.md → system architecture
docs/ai_prompts.md → AI assistance used

---

## 🚀 Future Improvements

* User authentication
* Quiz timer enforcement
* Leaderboard
* Randomized questions
* Web frontend

---

## 📝 License

MIT
