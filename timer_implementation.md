🎯 1. Purpose

The quiz platform enforces strict time control for every quiz attempt.
Each student must complete the quiz within a fixed duration defined by the teacher.

The timer system is designed to:

✔ Enforce time limits accurately
✔ Prevent late submissions
✔ Prevent timer manipulation
✔ Maintain fairness among users
✔ Support multiple simultaneous quiz attempts

🧠 2. Design Principle — Server Authoritative Timer

The timer follows a server-authoritative architecture.

This means:

✔ The backend server controls all time validation
✔ The frontend timer is only for visual display
✔ All enforcement happens on the server

🚫 Users cannot cheat by changing browser time or pausing scripts.

🗄 3. Timer Storage in Database

Each quiz attempt stores timing data permanently.

Attempt Table — Timer Fields
Field	Description
StartTime	Time when the attempt begins
EndTime	Submission deadline
IsSubmitted	Indicates completion status

This ensures:

✔ timer persists across refresh
✔ reconnection safe
✔ accurate tracking

▶ 4. Timer Initialization

When a student starts a quiz:

1️⃣ Server records current time
2️⃣ Quiz duration is fetched
3️⃣ Submission deadline is calculated

Formula
EndTime = StartTime + QuizDuration
Example

Quiz duration = 10 minutes

StartTime = 10:00 AM
EndTime   = 10:10 AM

Both values are stored in the Attempt record.

🛡 5. Timer Enforcement (Submission Validation)

Every submission is validated using stored deadline.

Validation Logic
CurrentServerTime vs EndTime
If time exceeded:

❌ submission rejected
❌ attempt auto-closed
❌ further submissions blocked

This guarantees strict enforcement.

🖥 6. Frontend Timer Display

The frontend shows a countdown timer for user awareness.

However:

✔ does NOT control submission
✔ only informational
✔ backend performs final validation

Remaining Time Calculation
RemainingTime = EndTime − CurrentServerTime
🔄 7. Handling Page Refresh or Reconnection

If a student refreshes the page:

✔ attempt already stored
✔ start time unchanged
✔ deadline unchanged
✔ timer continues normally

🚫 Timer cannot be reset.

👥 8. Concurrent Quiz Attempts

Each attempt is stored independently.

Therefore:

✔ multiple students can start at different times
✔ each student has separate deadline
✔ no shared timer exists

Database ensures full isolation.

🔐 9. Anti-Cheating Protection via Timer
Cheating Method	Prevention
Page refresh	Timer stored in database
Changing device time	Server time used
Submitting after deadline	Backend validation
Multiple submissions	Attempt locked
Pausing browser script	Server authority
🔄 10. Submission Workflow
Start Attempt

1️⃣ Student begins quiz
2️⃣ Server records start time
3️⃣ Server calculates deadline
4️⃣ Attempt stored in database

Submit Answers

1️⃣ Server retrieves attempt
2️⃣ Checks submission status
3️⃣ Validates deadline
4️⃣ If valid → evaluate answers
5️⃣ If expired → reject submission

✅ 11. Advantages of This Implementation

✔ Accurate time tracking
✔ Server-controlled validation
✔ Resistant to manipulation
✔ Persistent storage
✔ Supports concurrent users
✔ Real exam-like behavior

🚀 12. Possible Future Enhancements

⭐ Automatic submission when time expires
⭐ Warning alerts before deadline
⭐ Real-time WebSocket timer sync
⭐ Distributed time synchronization

🏁 13. Conclusion

The timer is implemented using a database-backed, server-validated model.

All critical timing decisions are handled by the backend, ensuring:

✔ fairness
✔ accuracy
✔ security
✔ resistance to manipulation

This approach is suitable for secure online examination systems.