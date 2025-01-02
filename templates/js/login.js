import { validateJWT } from "./utils/token.js";

const loginFormEl = document.getElementById('login-form');
const usernameEl = document.getElementById('username');
const passwordEl = document.getElementById('password');
const loginEl = document.getElementById('login');
const notifEl = document.getElementById('notification');

const loginHandler = async () => {
  const body = new FormData();
  body.append('username', usernameEl.value);
  body.append('password', passwordEl.value);
  
  const response = await fetch('/login', {
    method: 'POST',
    body,
  });

  const res = await response.json();
  const token = res?.data?.token;
  const message = res.message;
  notifEl.innerHTML = `${message}`;

  if (response.ok) notifEl.style.backgroundColor = "#22c55e";
  else notifEl.style.backgroundColor = "#f44336";

  if (token) {
    try {
      await validateJWT(token);
      window.location.href = '/';
    } catch (error) {
      console.error(error);
    }
  }
}

const setupLogin = () => {
  loginEl.addEventListener('click', loginHandler);
  loginFormEl.addEventListener("keypress", (event) => {
    if (event.key === "Enter") loginHandler();
  });
}

document.addEventListener("DOMContentLoaded", async () => {
  setupLogin();
});