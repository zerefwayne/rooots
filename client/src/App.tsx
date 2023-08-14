import { Route, Routes } from 'react-router-dom';
import AppLayout from "./AppLayout";
import Login from './Login';
import StravaLoading from './StravaCallback';
import RequireAuth from './components/RequireAuth';
import Home from './Home';
import PersistentLogin from './components/PersistLogin';

const App = () => {
    return (
        <Routes >
            <Route path="/" element={<AppLayout />} >
                <Route path="login" element={<Login />} />
                <Route path="auth/strava/callback" element={<StravaLoading />} />

                <Route element={<PersistentLogin />}>
                    <Route element={<RequireAuth />}>
                        <Route path="" element={<Home />} />
                    </Route>
                </Route>
            </Route>
        </Routes>
    )
}

export default App;