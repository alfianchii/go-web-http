import { logout } from "./auth.js"

export const validateJWT = async (token = getToken()) => {
  try {
    const response = await fetch('/validate-jwt', {
      method: 'POST',
      credentials: 'include',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) throw new Error((await response.json()).message);
    
    return await response.json();
  } catch (err) {
    return Promise.reject(err);
  }
}

export const invalidateSession = async (error) => {
  sessionStorage.removeItem("token");
  await logout(getToken());
  window.location.href = "/login";
  console.error(error);
};

export const getToken = () => sessionStorage.getItem('token');