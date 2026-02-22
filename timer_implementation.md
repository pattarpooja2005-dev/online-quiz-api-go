1. Purpose

The quiz platform requires strict time control for each quiz attempt.
Every student must complete the quiz within a fixed duration defined by the teacher.

The timer system is designed to:

enforce time limits accurately

prevent late submissions

prevent timer manipulation

maintain fairness among users

support multiple simultaneous quiz attempts

2. Design Principle

The timer follows a server-authoritative design.

This means:

👉 The backend server controls time validation
👉 The frontend timer is only for display
👉 All enforcement happens on the server

This prevents cheating through browser manipulation.

3. Timer Storage in Database

Each quiz attempt stores timing information permanently in the database.

Attempt Table Fields
Field	Description
StartTime	Time when attempt begins
EndTime	Deadline for submission
IsSubmitted	Indicates completion status

These fields allow the system to track timing even if the user refreshes the page or disconnects.

4. Timer Initialization

When a student starts a quiz:

The system records the current server time.

The quiz duration is retrieved from quiz settings.

The submission deadline is calculated.

Formula
EndTime = StartTime + QuizDuration
Example

Quiz duration = 10 minutes

StartTime = 10:00 AM
EndTime   = 10:10 AM

Both values are stored in the Attempt record.

5. Timer Enforcement (Submission Validation)

Every submission request is validated using the stored deadline.

The system compares:

CurrentServerTime vs EndTime
If current time exceeds deadline

submission is rejected OR auto-closed

attempt marked completed

further submissions blocked

This guarantees strict time enforcement.

6. Frontend Timer Display

The frontend displays a countdown timer to help students track remaining time.

However:

this timer does not control submission

it is only informational

backend always performs final validation

Remaining time is calculated as:

RemainingTime = EndTime − CurrentServerTime
7. Handling Page Refresh or Reconnection

If the student refreshes the page:

attempt record already exists

start time remains unchanged

deadline remains unchanged

timer continues normally

This prevents restarting the quiz timer.

8. Concurrent Quiz Attempts

Each student attempt is stored separately.

Therefore:

multiple students can start the quiz at different times

each student has an independent deadline

no shared timer exists

The database ensures complete isolation of attempts.

9. Anti-Cheating Protection via Timer

The timer system prevents common cheating methods:

Cheating Method	Prevention
Refresh page	timer stored in database
Change device time	server time used
Submit after deadline	backend validation
Multiple submissions	attempt locked
Pause browser script	server authority
10. Submission Workflow
Start Attempt

Student begins quiz

Server records start time

Server calculates end time

Attempt stored in database

Submit Answers

Server retrieves attempt

Checks submission status

Validates deadline

If valid → evaluate answers

If expired → reject submission

11. Advantages of This Implementation

✔ Accurate time tracking
✔ Server-controlled validation
✔ Resistant to manipulation
✔ Persistent storage
✔ Supports concurrent users
✔ Real exam-like behavior

12. Possible Future Enhancements

automatic background submission when time expires

warning alerts before deadline

real-time WebSocket timer sync

distributed time synchronization

13. Conclusion

The timer is implemented using a database-backed, server-validated model.
All critical timing decisions are handled by the backend, ensuring fairness, accuracy, and resistance to manipulation.

This approach is suitable for secure online examination systems.