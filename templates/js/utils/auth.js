export const logout = async (token) => {
  try {
    const response = await fetch('/logout', {
      method: 'POST',
      credentials: 'include',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) throw new Error(response.statusText);

    sessionStorage.removeItem('token');
    window.location.href = '/login';
  } catch (err) {
    console.error(err);
  }
}