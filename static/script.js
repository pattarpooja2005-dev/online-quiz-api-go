const API = "http://localhost:8080";
let attemptID = null;

// start quiz
async function startQuiz() {

  const quizID = document.getElementById("quizID").value;
  const name = document.getElementById("username").value;

  const res = await fetch(API + "/start", {
    method:"POST",
    headers:{ "Content-Type":"application/json"},
    body: JSON.stringify({
      quizID:Number(quizID),
      userName:name
    })
  });

  const data = await res.json();
  attemptID = data.ID;

  localStorage.setItem("attemptID", attemptID);
  localStorage.setItem("endTime", Date.now() + 600000); // 10 min

  window.location.href = "/quiz";
}

// timer countdown
if (document.getElementById("timer")) {

  const end = localStorage.getItem("endTime");

  setInterval(()=>{
    const remaining = Math.floor((end - Date.now())/1000);
    document.getElementById("timer").innerText =
      "Time Left: " + remaining + " sec";

    if(remaining <= 0){
      alert("Time up!");
      window.location.href="/";
    }
  },1000);
}

// submit answer
async function submitAnswer(){

  const res = await fetch(API + "/submit",{
    method:"POST",
    headers:{ "Content-Type":"application/json"},
    body: JSON.stringify([{
      attemptID:Number(localStorage.getItem("attemptID")),
      questionID:Number(questionID.value),
      selected:selected.value
    }])
  });

  const data = await res.json();
  window.location.href="/result?score="+data.score;
}