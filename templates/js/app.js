const socket = new WebSocket("ws://localhost:3333/ws");
const tabId = `Tab-${Math.random().toString(36).substring(7)}`;
const generateDate = () => new Date().toISOString().replace('T', ' ').split('.')[0];
const sendData = (data) => socket.send(JSON.stringify(data));
const initDOMElements = () => ({
  input: document.getElementById("fullname"),
  button: document.getElementById("submit-message"),
  notification: document.getElementById("notification"),
  messages: document.getElementById("messages"),
});

const setupSocket = ({ input, button, notification, messages }) => {
  socket.onopen = () => {
    console.log("WebSocket connection established");

    button.addEventListener("click", () => {
      if (input.value) {
        sendData({ tabId, text: `${generateDate()} - ${input.value}`, send: true });
        input.focus();
        input.value = "";
      }
    });

    input.addEventListener("input", () => sendData({ tabId, text: input.value, typing: true }));

    input.addEventListener("keypress", (event) => {
      if (event.key === "Enter" && input.value) {
        sendData({ tabId, text: `${generateDate()} - ${input.value}`, send: true });
        sendData({ tabId, typing: false });
        input.focus();
        input.value = "";
      }
    });
  }

  socket.onmessage = (event) => {
    const data = JSON.parse(event.data);
    if (data.typers) updateNotification(data.typers, notification);
    if (data.send) displayMessage(data.text, messages);
  };
}

const updateNotification = (typers, notification) => {
  if (typers.length === 0) notification.innerText = "Online";
  if (typers.length === 1) notification.innerText = `${typers[0]} is typing...`;
  if (typers.length > 1) notification.innerText = `${typers.join(", ")} are typing...`;
}

const displayMessage = (text, messages) => {
  const messageElement = document.createElement("p");
  messageElement.innerText = text;

  if (messages.firstChild) messages.insertBefore(messageElement, messages.firstChild);
  else messages.appendChild(messageElement);
}

document.addEventListener("DOMContentLoaded", () => {
  const elements = initDOMElements();
  setupSocket(elements);
});