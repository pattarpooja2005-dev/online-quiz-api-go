# System Design

## Overview

This system is a REST API for an online quiz platform.

## Workflow

1. Teacher creates quiz
2. Teacher adds questions
3. Student starts quiz attempt
4. Student submits answers
5. System compares answers with correct options
6. Score calculated automatically
7. Attempt stored with start and end time

## Database Relationships

* One Quiz → Many Questions
* One Attempt → Many Answers

## Scoring Logic

Each correct answer = 1 point.
Total score = number of correct answers.
