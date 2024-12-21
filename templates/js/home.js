import { logout } from "./utils/auth.js";
import { getToken } from "./utils/token.js";
import { validateJWT } from "./utils/token.js";

const elLogout = document.getElementById("logout");
let clientId;
const generateDate = (str) => str ? new Date(str).toISOString().replace('T', ' ').split('.')[0] : new Date().toISOString().replace('T', ' ').split('.')[0];
const sendData = (socket, data) => socket.send(JSON.stringify(data));
const initDOMElements = () => ({
  elInput: document.getElementById("fullname"),
  elButton: document.getElementById("submit-message"),
  elNotif: document.getElementById("notification"),
  elMessage: document.getElementById("messages"),
});

const setupSocket = ({ elInput, elButton, elNotif, elMessage }) => {
  const socket = new WebSocket(`ws://${document.location.host}/chats`);

  if (!window.WebSocket) return elNotif.innerText = "WebSocket is not supported by your browser.";
  
  socket.onopen = () => {
    console.log("WebSocket connection established");

    elButton.addEventListener("click", () => {
      if (elInput.value) {
        sendData(socket, { clientId, text: `${elInput.value}`, send: true });
        resetInput(elInput);
      }
    });

    elInput.addEventListener("input", () => sendData(socket, { clientId, text: elInput.value, typing: true }));

    elInput.addEventListener("keypress", (event) => {
      if (event.key === "Enter" && elInput.value) {
        sendData(socket, { clientId, text: `${elInput.value}`, send: true });
        sendData(socket, { clientId, typing: false });
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
  const newEl = document.createElement("p");
  newEl.innerText = text;

  if (type === "chats") msgEl.appendChild(newEl);
  if (type === "message") {
    if (msgEl.firstChild) msgEl.insertBefore(newEl, msgEl.firstChild);
    else msgEl.appendChild(newEl);
  }
}

document.addEventListener("DOMContentLoaded", async () => {
  try {
    const { data } = await validateJWT(getToken());
    clientId = data.username;
  } catch (error) {
    window.location.href = "/login";
    console.error(error);
  }
  
  elLogout.addEventListener("click", async () => logout(getToken()));
  
  const elements = initDOMElements();
  setupSocket(elements);
});