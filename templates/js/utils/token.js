import { logout, getCookie } from "./auth.js"
import { JWT_TOKEN } from "./const.js"

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
  await logout(getToken());
  window.location.href = "/login";
  console.error(error);
};

export const getToken = () => getCookie(JWT_TOKEN);