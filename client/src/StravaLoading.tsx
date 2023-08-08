import axios from "axios";
import { useEffect, useState } from "react";

interface ExchangeTokenResponse {
    authenticationToken: string
}

function StravaLoading() {
    const params = new URLSearchParams(window.location.search);
    const code = params.get("code");
    const [authenticationToken, setAuthenticationToken] = useState('');

    useEffect(() => {
        const handleLoginWithStrava = async () => {
            try {
                // The API will be hit twice when running locally
                // To avoid there are 2 alternatives: 1. Remove StrictMode from index.tsx 2. Run the production build
                const {data} = await axios.post<ExchangeTokenResponse>('http://localhost:8081/auth/strava/exchangeToken', { code })
                console.log(data);
                setAuthenticationToken(data.authenticationToken || '');
                if (!data.authenticationToken) return
                sessionStorage.setItem("authentication_token", data.authenticationToken)
            } catch (err) {
                console.error(err)
            }
        }  
        handleLoginWithStrava();
    }, [code])

    return (
        <div>
            <ul>
                <li>Code: {code}</li>
                <li>Authentication Token: {authenticationToken}</li>
            </ul>
        </div>
    )

}

export default StravaLoading