const dashboard = document.getElementById("dashboard");

const addFamilyButton = document.getElementById("add-family");
addFamilyButton.addEventListener("click", addFamily);

function addFamily() {
  const familyTemplate = document.createElement("div");
  familyTemplate.className = "family-container form";

  const familyName = document.createElement("input");
  familyName.type = "text";
  familyName.className = "family-name";
  familyName.defaultValue = `Family n°${dashboard.childElementCount}`;
  familyTemplate.appendChild(familyName);

  const addMemberButton = document.createElement("button");
  addMemberButton.className = "family-button add";
  addMemberButton.onclick = addFamilyMember;
  addMemberButton.innerText = "+";
  familyTemplate.append(addMemberButton);

  addFamilyMember({ target: addMemberButton });

  dashboard.insertBefore(familyTemplate, addFamilyButton);
}

const colorList = ["#00000000", "#00FF00", "#00FFFF", "#0000FF", "#FF0000"];
let MutExMap = new Map();
function addToMutExMAP(event) {
  let input = event.target;
  input.innerText++;
  input.innerText %= 5;
  input.parentNode.style.backgroundColor = colorList[parseInt(input.innerText)];
}

function addFamilyMember(event) {
  const familyContainer = event.target.parentNode;
  const memberContainer = document.createElement("div");
  memberContainer.classList = "member-container";

  const memberField = document.createElement("input");
  memberField.type = "text";
  memberField.className = "member";
  memberField.defaultValue = `#${familyContainer.childElementCount-1}`
  memberContainer.appendChild(memberField);

  const mutexButton = document.createElement("button");
  mutexButton.className = "member-button mutex";
  mutexButton.innerText = 0;
  mutexButton.onclick = addToMutExMAP;
  memberContainer.appendChild(mutexButton);

  const deleteButton = document.createElement("button");
  deleteButton.className = "member-button";
  deleteButton.innerText = "X";
  deleteButton.onclick = deleteFamilyMember;
  memberContainer.appendChild(deleteButton);

  familyContainer.insertBefore(memberContainer, event.target);
}

function deleteFamilyMember(event) {
  const member = event.target.parentNode;
  const family = member.parentNode;
  if (dashboard.childElementCount > 2) {
    family.removeChild(member);
    if (family.childElementCount == 2) {
      family.parentNode.removeChild(family);
    }
  }
}

window.onload = async () => {
  await connectWS()
  addFamily();
  addFamily();
  
};
