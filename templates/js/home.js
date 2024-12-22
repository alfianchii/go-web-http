import { logout } from "./utils/auth.js";
import { getToken } from "./utils/token.js";
import { validateJWT } from "./utils/token.js";

const elLogout = document.getElementById("logout");
let username;
const generateDate = (str) => {
  const date = new Date(str);
  const options = { year: 'numeric', month: 'short', day: '2-digit', hour: '2-digit', minute: '2-digit', second: '2-digit', hour12: false };
  const formattedDate = new Intl.DateTimeFormat('en-US', options).format(date);
  const [day, month, year, time] = formattedDate.split(/,?\s+/);
  return `${year}, ${month} ${day} at ${time}`;
};
const sendData = (ws, data) => ws.send(JSON.stringify(data));
const initDOMElements = () => ({
  inputEl: document.getElementById("chat"),
  buttonEl: document.getElementById("submit-message"),
  notifEl: document.getElementById("notification"),
  messageListEl: document.getElementById("messageList"),
});

const setupSocket = ({ inputEl, buttonEl, notifEl, messageListEl }) => {
  const ws = new WebSocket(`ws://${document.location.host}/chats`);

  if (!window.WebSocket) return notifEl.innerText = "WebSocket is not supported by your browser.";
  
  ws.onopen = () => {
    console.log("WebSocket connection established");

    buttonEl.addEventListener("click", () => {
      if (inputEl.value) {
        sendData(ws, { username, text: `${inputEl.value}`, send: true });
        resetInput(inputEl);
      }
    });

    inputEl.addEventListener("input", () => sendData(ws, { username, text: inputEl.value, typing: true }));

    inputEl.addEventListener("keypress", (event) => {
      if (event.key === "Enter" && inputEl.value) {
        sendData(ws, { username, text: `${inputEl.value}`, send: true });
        sendData(ws, { username, typing: false });
        resetInput(inputEl);
      }
    });
  }

  ws.onmessage = (event) => {
    const data = JSON.parse(event.data);
    displayChats(messageListEl, data)
    if (data?.typers) updateNotification(notifEl, data.typers);
    if (data?.send) displayMessage(messageListEl, data);
  };
}

const displayChats = (msgEl, data) => {
  if (Array.isArray(data)) 
    for (const chat of data) createTextElement("chats", msgEl, chat);
}

const updateNotification = (notifEl, typers) => {
  if (typers.length === 0) notifEl.innerText = "Online";
  if (typers.length === 1) notifEl.innerText = `${typers[0]} is typing...`;
  if (typers.length > 1) notifEl.innerText = `${typers.join(", ")} are typing...`;
}

const displayMessage = (msgEl, message) => createTextElement("message", msgEl, message);

const resetInput = (input) => {
  input.focus();
  input.value = "";
}

const createTextElement = (type, msgEl, data) => {
  const newEl = document.createElement("p");
  newEl.innerHTML = createText(data);

  if (type === "chats") msgEl.appendChild(newEl);
  if (type === "message") {
    if (msgEl.firstChild) msgEl.insertBefore(newEl, msgEl.firstChild);
    else msgEl.appendChild(newEl);
  }
}

const createText = (data) => `${generateDate(data.createdAt)} - <span style='background-color: yellow;'>${data.username}</span>: ${data.text}`

document.addEventListener("DOMContentLoaded", async () => {
  try {
    const { data } = await validateJWT(getToken());
    username = data.username;
  } catch (error) {
    window.location.href = "/login";
    console.error(error);
  }
  
  elLogout.addEventListener("click", async () => logout(getToken()));
  
  const elements = initDOMElements();
  setupSocket(elements);
});