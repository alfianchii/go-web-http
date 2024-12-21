export const validateJWT = async (token) => {
  try {
    const jwtToken = token || getToken();
    
    const response = await fetch('/validate-jwt', {
      method: 'POST',
      credentials: 'include',
      headers: {
        Authorization: `Bearer ${jwtToken}`,
      },
    });

    if (!response.ok) throw new Error(response.statusText);
    
    return await response.json();
  } catch (err) {
    return Promise.reject(err);
  }
}

export const getToken = () => sessionStorage.getItem('token');