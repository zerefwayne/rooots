import { useEffect, useState } from "react";
import useRefreshToken from "./hooks/useRefreshToken";
import { Button } from "antd";
import useAxiosPrivate from "./hooks/useAxiosPrivate";

const User = () => {
    const [user, setUser] = useState();
    const refreshToken = useRefreshToken();
    const axiosPrivate = useAxiosPrivate();

    useEffect(() => {
        let isMounted = true;
        const controller = new AbortController();

        const getUser = async () => {
            try {
                const response = await axiosPrivate.get("/strava/user", { signal: controller.signal });
                console.log(response.data);

                isMounted && setUser(response.data);
            } catch (err) {
                console.error(err);
            }
        }

        getUser();

        return () => {
            isMounted = false;
            controller.abort();
        }
    }, []);

    return (
        <article>
            <h2>User</h2>
            <p>
                {JSON.stringify(user)}
            </p>
            <p>
                {user === null && <h1>No user available</h1>}
            </p>
            <Button type="primary" onClick={() => refreshToken()}>Refresh Token</Button>
        </article>
    )
};

export default User;