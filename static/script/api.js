const submitButton = document.getElementById("submit");
submitButton.onclick = submitHandle;

const isUnstrict = document.getElementById("is-unstrict")

function collectQuery() {
  let Person = new Array();
  let families = document.querySelectorAll(".family-container.form");

  for (let family of families) {
    let familyName = family.querySelector(".family-name").value;
    let members = family.querySelectorAll(".member-container");
    for (let member of members) {
      let memberField = member.getElementsByClassName("member")[0];
      let str = memberField.value;
      Person.push({
        Name: str,
        Family: familyName,
        Mutex: parseInt(member.getElementsByClassName("mutex")[0].innerText),
      });
    }
  }
  return {
    Gifters : Person,
    Unstrict : isUnstrict.checked
  };
}

async function submitHandle() {
  submitButton.disabled = true;
  query = collectQuery();
  if (!ws) {
    connectWS();
  }
  ws.send(JSON.stringify(query))
}

let ws;
function connectWS() {
  return new Promise((r) => {
    ws = new WebSocket("ws://" + location.host + "/ws", "draft");
    ws.onopen = () => {
      console.log("Connection openned.");
      r();
    };
    ws.onclose = () => {
      setTimeout(connectWS, 1000);
    };
    ws.onmessage = (m)=>{
      console.log(JSON.parse(m.data))
      printResult(JSON.parse(m.data))
    }
  });
}
