const registerFormEl = document.getElementById('register-form');
const usernameEl = document.getElementById('username');
const passwordEl = document.getElementById('password');
const registerEl = document.getElementById('register');
const notifEl = document.getElementById('notification');

const registerHandler = async () => {
  const body = new FormData();
  body.append('username', usernameEl.value);
  body.append('password', passwordEl.value);

  const response = await fetch('/user', {
    method: 'POST',
    body,
  });

  const res = await response.json();
  const message = res.message;
  notifEl.innerHTML = `${message}`;

  if (response.ok) {
    notifEl.innerHTML = `${message}. Let's <a href="/login">login</a> to your account`;
    notifEl.style.backgroundColor = "#22c55e";
  } else notifEl.style.backgroundColor = "#f44336";
}

const setupRegister = () => {
  registerEl.addEventListener('click', registerHandler);
  registerFormEl.addEventListener("keypress", (event) => {
    if (event.key === "Enter") registerHandler();
  });
}

document.addEventListener("DOMContentLoaded", async () => {
  setupRegister();
});