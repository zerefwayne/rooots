import axios from 'axios';
import { useState, useEffect } from 'react';

function App() {

  const [user, setUser] = useState<string | null>(null);

  useEffect(() => {
    const authenticationToken = sessionStorage.getItem("authentication_token");
    setUser(authenticationToken);
  }, []);

  const handleLogin = () => {
    axios.get<string>("http://localhost:8081/auth/strava/login").then(({ data: redirectUrl }) => {
      window.location.replace(redirectUrl);
    }).catch(err => {
      console.error(err);
    })
  };

  const handleLogout = () => {
    sessionStorage.clear()
    window.location.reload()
  }

  return (
    <div>
      {user ?
        (
          <>
            <h1>Logged in!</h1>
            <button onClick={handleLogout}>Logout</button>
          </>
        ) :
        <button onClick={handleLogin}>Login with Strava</button>}
    </div>
  );
}

export default App;
