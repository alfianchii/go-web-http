import { validateJWT } from "./utils/token.js";

const username = document.getElementById('username');
const password = document.getElementById('password');
const login = document.getElementById('login');
const messages = document.getElementById('messages');

const loginHandler = async () => {
  const body = new FormData();
  body.append('username', username.value);
  body.append('password', password.value);
  
  const response = await fetch('/login', {
    method: 'POST',
    body,
  });

  const result = await response.json();
  const token = result.data.token;
  const message = result.message;
  messages.innerHTML = message;

  try {
    await validateJWT(token);
    sessionStorage.setItem('token', token);
    window.location.href = '/';
  } catch (error) {
    console.error(error);
  }
}

const setupLogin = () => {
  login.addEventListener('click', loginHandler);
}

document.addEventListener("DOMContentLoaded", async () => {
  setupLogin();
});