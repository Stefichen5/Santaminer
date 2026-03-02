async function sendMineRequest(data) {
  try {
    let tableRef = document.getElementById("resultTable");
    let newRow = tableRef.insertRow();

    // Insert a cell in the row at index 0
    let answerKeyCell = newRow.insertCell(0);

    answerKeyCell.innerHTML="<div class=loader></div>";


    const res = await fetch("http://localhost:2512/mine",
      {method: "POST",
        headers: {'Content-Type':'text/plain'},
        body: JSON.stringify(data)});
    const resultData = await res.json();
    console.log(resultData);



    let answerMD5sumCell = newRow.insertCell(1);
    // Append a text node to the cell
    answerKeyCell.innerHTML=""
    answerKeyCell.textContent=resultData.AnswerKey;
    answerMD5sumCell.textContent=resultData.AnswerMD5sum;
  } catch (err) {
    console.log("Error: ",err);
  }
}

document.getElementById("minerForm").addEventListener("submit", async e=>{
  e.preventDefault();
  const form = e.target;
  const secretkey = document.getElementById("secretkey").value.trim();
  const requiredLeadingZeros = parseInt(document.getElementById("requiredZeros").value);
  const data = { "secretkey": secretkey, "requiredLeadingZeros": requiredLeadingZeros };
  console.log(data)

  sendMineRequest(data);

  let tableRef = document.getElementById("historyTable");
  let newRow = tableRef.insertRow();

  let secretKeyCell = newRow.insertCell(0);
  let requiredLeadingZerosCell = newRow.insertCell(1);

  // Append a text node to the cell
  secretKeyCell.appendChild(document.createTextNode(secretkey));
  requiredLeadingZerosCell.appendChild(document.createTextNode(requiredLeadingZeros));

  form.reset();
});
