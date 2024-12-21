import { logout } from "./utils/auth.js";
import { getToken } from "./utils/token.js";
import { validateJWT } from "./utils/token.js";

const elLogout = document.getElementById("logout");
let clientId;
const generateDate = (str) => str ? new Date(str).toISOString().replace('T', ' ').split('.')[0] : new Date().toISOString().replace('T', ' ').split('.')[0];
const sendData = (socket, data) => socket.send(JSON.stringify(data));
const initDOMElements = () => ({
  inputEl: document.getElementById("chat"),
  buttonEl: document.getElementById("submit-message"),
  notifEl: document.getElementById("notification"),
  messageListEl: document.getElementById("messageList"),
});

const setupSocket = ({ inputEl, buttonEl, notifEl, messageListEl }) => {
  const socket = new WebSocket(`ws://${document.location.host}/chats`);

  if (!window.WebSocket) return notifEl.innerText = "WebSocket is not supported by your browser.";
  
  socket.onopen = () => {
    console.log("WebSocket connection established");

    buttonEl.addEventListener("click", () => {
      if (inputEl.value) {
        sendData(socket, { clientId, text: `${inputEl.value}`, send: true });
        resetInput(inputEl);
      }
    });

    inputEl.addEventListener("input", () => sendData(socket, { clientId, text: inputEl.value, typing: true }));

    inputEl.addEventListener("keypress", (event) => {
      if (event.key === "Enter" && inputEl.value) {
        sendData(socket, { clientId, text: `${inputEl.value}`, send: true });
        sendData(socket, { clientId, typing: false });
        resetInput(inputEl);
      }
    });
  }

  socket.onmessage = (event) => {
    const data = JSON.parse(event.data);
    displayChats(data, messageListEl)
    if (data?.typers) updateNotification(data.typers, notifEl);
    if (data?.send) displayMessage(data.text, messageListEl, data.createdAt);
  };
}

const displayChats = (data, msgEl) => {
  if (Array.isArray(data)) 
    for (const message of data) createTextElement("chats", msgEl, `${generateDate(message.createdAt)}: ${message.text}`);
}

const updateNotification = (typers, notifEl) => {
  if (typers.length === 0) notifEl.innerText = "Online";
  if (typers.length === 1) notifEl.innerText = `${typers[0]} is typing...`;
  if (typers.length > 1) notifEl.innerText = `${typers.join(", ")} are typing...`;
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