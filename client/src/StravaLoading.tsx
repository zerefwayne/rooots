import axios from "axios";
import { useEffect } from "react";

function StravaLoading() {

    const params = new URLSearchParams(window.location.search);
    const code = params.get("code");

    useEffect(() => {
        axios.post('http://localhost:8081/auth/strava/exchangeToken', { code })
            .then(res => console.log(res))
            .catch(err => console.error(err));
    })

    return (
        <div>
            <ul>
                <li>Code: {code}</li>
            </ul>
        </div>
    )

}

export default StravaLoading