const socket = new WebSocket(`ws://${document.location.host}/chats`);
const clientId = `User-${Math.random().toString(36).substring(7)}`;
const generateDate = (str) => str ? new Date(str).toISOString().replace('T', ' ').split('.')[0] : new Date().toISOString().replace('T', ' ').split('.')[0];
const sendData = (data) => socket.send(JSON.stringify(data));
const initDOMElements = () => ({
  elInput: document.getElementById("fullname"),
  elButton: document.getElementById("submit-message"),
  elNotif: document.getElementById("notification"),
  elMessage: document.getElementById("messages"),
});

const setupSocket = ({ elInput, elButton, elNotif, elMessage }) => {
  if (!window.WebSocket) return elNotif.innerText = "WebSocket is not supported by your browser.";
  
  socket.onopen = () => {
    console.log("WebSocket connection established");

    elButton.addEventListener("click", () => {
      if (elInput.value) {
        sendData({ clientId, text: `${elInput.value}`, send: true });
        resetInput(elInput);
      }
    });

    elInput.addEventListener("input", () => sendData({ clientId, text: elInput.value, typing: true }));

    elInput.addEventListener("keypress", (event) => {
      if (event.key === "Enter" && elInput.value) {
        sendData({ clientId, text: `${elInput.value}`, send: true });
        sendData({ clientId, typing: false });
        resetInput(elInput);
      }
    });
  }

  socket.onmessage = (event) => {
    const data = JSON.parse(event.data);
    displayChats(data, elMessage)
    if (data?.typers) updateNotification(data.typers, elNotif);
    if (data?.send) displayMessage(data.text, elMessage, data.createdAt);
  };
}

const displayChats = (data, msgEl) => {
  if (Array.isArray(data)) 
    for (const message of data) createTextElement("chats", msgEl, `${generateDate(message.createdAt)}: ${message.text}`);
}

const updateNotification = (typers, elNotif) => {
  if (typers.length === 0) elNotif.innerText = "Online";
  if (typers.length === 1) elNotif.innerText = `${typers[0]} is typing...`;
  if (typers.length > 1) elNotif.innerText = `${typers.join(", ")} are typing...`;
}

const displayMessage = (text, msgEl, date) => createTextElement("message", msgEl, `${generateDate(date)}: ${text}`);

const resetInput = (input) => {
  input.focus();
  input.value = "";
}

const createTextElement = (type, msgEl, text) => {
  newEl = document.createElement("p");
  newEl.innerText = text;

  if (type === "chats") msgEl.appendChild(newEl);
  if (type === "message") {
    if (msgEl.firstChild) msgEl.insertBefore(newEl, msgEl.firstChild);
    else msgEl.appendChild(newEl);
  }
}

document.addEventListener("DOMContentLoaded", () => {
  const elements = initDOMElements();
  setupSocket(elements);
});