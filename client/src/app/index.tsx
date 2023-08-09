
import axios from 'axios';
import { useEffect, useState } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './index.css';
import { Home } from './pages/Home';

import { Button, Layout } from 'antd';
import { StravaLoading } from './pages/StravaCallback';

const { Content, Header } = Layout;

export function App() {
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
    <BrowserRouter>
      <Layout>
        <Header style={{ display: 'flex', alignItems: 'center' }}>
          {user
            ? <Button type='primary' danger onClick={handleLogout}>Logout</Button>
            : <Button type='primary' onClick={handleLogin}>Login with Strava</Button>}
        </Header>
        <Content>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/auth/strava/callback" element={<StravaLoading />} />
            <Route path="*" element={<h1>Not Found</h1>} />
          </Routes>
        </Content>
      </Layout>
    </BrowserRouter>
  );
}