import { useEffect, useState } from "react";
import useAuth from "./hooks/useAuth";
import useAxiosPrivate from "./hooks/useAxiosPrivate";
import { Alert, Button, Layout } from "antd";
import { useNavigate } from "react-router-dom";
import useLogout from "./hooks/useLogout";
import { Content } from "antd/es/layout/layout";

const User = () => {
    const [user, setUser] = useState();
    const axiosPrivate = useAxiosPrivate();
    const { auth } = useAuth();
    const logout = useLogout();
    const navigate = useNavigate();

    const handleLogout = async () => {
        await logout();
        navigate('/login');
    };

    useEffect(() => {
        let isMounted = true;
        const controller = new AbortController();

        const getUser = async () => {
            try {
                const response = await axiosPrivate.get(`/strava/user`, { signal: controller.signal });
                console.log(response.data);

                isMounted && setUser(response.data);
            } catch (err) {
                console.error(err);
            }
        }

        if (auth?.accessToken) {
            getUser();
        }

        return () => {
            isMounted = false;
            controller.abort();
        }
    }, []);

    return (
        <Layout>
            <Content>
                <Alert
                    description="Page under construction"
                    type="warning"
                    showIcon
                    closable
                />
                <h2>User</h2>
                <p>
                    {JSON.stringify(user)}
                </p>
                <p>
                    {user === null && <h1>No user available</h1>}
                </p>
                <Button onClick={handleLogout}>Logout</Button>
            </Content>
        </Layout>
    )
};

export default User;