import { Button } from "antd";
import axios from "./api/axios";

const LOGIN_URL = '/auth/strava/login';

const handleLogin = async () => {
    try {
        const { data: stravaRedirectUrl } = await axios.get<string>(LOGIN_URL);
        console.log("Redirecting to: ", stravaRedirectUrl);
        window.location.replace(stravaRedirectUrl);
    } catch (err) {
        console.error(err);
    }
}

const Login = () => {
    return (
        <div>
            <Button type="primary" onClick={handleLogin}>Login with Strava</Button>
        </div>
    )
}

export default Login;