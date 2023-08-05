import axios from 'axios';

function App() {

  const handleLogin = () => {
    axios.get<string>("http://localhost:8081/auth/strava/login").then(({data: redirectUrl}) => {
      window.location.replace(redirectUrl);
    }).catch(err => {
      console.error(err);
    })
  };

  return (
    <div>
      <button onClick={handleLogin}>Login with Strava</button>
    </div>
  );
}

export default App;
