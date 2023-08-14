import { useEffect } from "react";
import axios from "./api/axios";
import useAuth from "./hooks/useAuth";
import { useNavigate } from "react-router-dom";

interface ExchangeTokenResponse {
    accessToken: string
    name: string
    userId: string
}

const StravaLoading = () => {
    const params = new URLSearchParams(window.location.search);
    const code = params.get("code");
    const { setAuth } = useAuth();
    const navigate = useNavigate();

    useEffect(() => {
        const handleLoginWithStrava = async () => {
            try {
                const { data } = await axios.post<ExchangeTokenResponse>(
                    '/auth/strava/exchangeToken', { code })
                console.log(data);

                const accessToken = data?.accessToken;
                const name = data?.name;
                const userId = data?.userId;

                setAuth && setAuth({ accessToken, name, userId });
                navigate('/');
            } catch (err) {
                console.error(err)
            }
        }
        handleLoginWithStrava();
    })

    return (
        <div>
            <ul>
                <li>Code: {code}</li>
            </ul>
        </div>
    )

}

export default StravaLoading;