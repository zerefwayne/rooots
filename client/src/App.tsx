import { Route, Routes } from 'react-router-dom';
import AppLayout from "./AppLayout";
import AppNavLayout from './AppNavLayout';
import Home from './Home';
import Login from './Login';
import StravaLoading from './StravaCallback';
import PersistentLogin from './components/PersistLogin';
import RequireAuth from './components/RequireAuth';

const App = () => {
    return (
        <Routes >
            <Route path="/" element={<AppLayout />} >
                <Route path="login" element={<Login />} />
                <Route path="auth/strava/callback" element={<StravaLoading />} />

                <Route element={<PersistentLogin />}>
                    <Route element={<RequireAuth />}>
                            <Route path="" element={<AppNavLayout />}>
                                <Route path="" element={<Home />} />
                            </Route>
                        </Route>
                    </Route>
            </Route>
        </Routes>
    )
}

export default App;