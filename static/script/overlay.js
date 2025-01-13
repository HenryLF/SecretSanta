const overlayResult = document.getElementById("overlay-result");

const offerToTemplate = document.getElementById("offerto-template")


function printResult(result) {
  submitButton.disabled = false;
  if (!result[0]) {
    alert("Too much constrains.");
    return;
  }
  overlayResult.style.display = "flex";
  result[0].forEach((g, k) => {
    let r = result[1][k];
    renderPair(g, r);
  });
  overlayResult.onclick = () => {
    overlayResult.innerHTML = "";
    overlayResult.style.display = "none";
    submitButton.disabled = false;
    overlayResult.onclick = undefined;
  };
}

function renderPair(g, r) {
  const Pair = document.createElement("div");
  Pair.className = "pair-container";
  
  const G = document.createElement("div");
  G.className = "gifter";

  G.appendChild(newName(g.Name));
  G.appendChild(newFName(g.Family));
  Pair.appendChild(G);

  Pair.appendChild(offerToTemplate.content.cloneNode(true))
  
  const R = document.createElement("div");
  R.className = "receiver";
  R.appendChild(newName(r.Name));
  R.appendChild(newFName(r.Family));
  Pair.appendChild(R);
  
  overlayResult.append(Pair);
}

const NameTemplate = document.createElement("p");
NameTemplate.className = "name";

function newName(txt) {
  let out = NameTemplate.cloneNode();
  out.innerText = txt;
  return out;
}
const FNameTemplate = document.createElement("p");
FNameTemplate.className = "name";

function newFName(txt) {
  let out = FNameTemplate.cloneNode();
  out.innerText = txt;
  return out;
}

const helpButton = document.getElementById("help");
helpButton.onclick = printHelp;

const overlayHelp = document.getElementById("overlay-help");
function printHelp() {
  overlayHelp.style.display = "flex";
  overlayHelp.onclick = () => {
    overlayHelp.onclick = undefined;
    overlayHelp.style.display = "none";
  };
}

const overlaySettings = document.getElementById("overlay-settings");
const quitSettings = document.getElementById("quit");

function openSettings() {
  overlaySettings.style.display = "flex";
  quitSettings.onclick = () => {
    overlaySettings.style.display = "none";
  };
}

const settingsButton = document.getElementById("settings");
settingsButton.onclick = openSettings;
