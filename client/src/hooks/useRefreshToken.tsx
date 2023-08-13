import axios from "../api/axios";
import { Auth } from "../context/AuthProvider";
import useAuth from "./useAuth";

interface RefreshSuccessResponse {
    accessToken: string
    expiresAt: Date
}

const useRefreshToken = () => {
    const { setAuth } = useAuth();

    const refresh = async () => {
        try {
            const response = await axios.get<RefreshSuccessResponse>('/auth/strava/refreshToken', {
                withCredentials: true
            });

            setAuth && setAuth((auth: Auth) => {
                if (response.data.accessToken) {
                    console.log("PREVIOUS STATE", JSON.stringify(auth));
                    console.log(response.data.accessToken);
                    return { ...auth, accessToken: response.data.accessToken }
                } else {
                    console.log("Nothing to change in setAuth");
                    return auth
                }
            });

            return response.data.accessToken;
        } catch (err) {
            console.error(err)
            return null;
        }
    };

    return refresh;
};

export default useRefreshToken;