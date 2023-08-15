import { Header } from "antd/es/layout/layout";

import { Avatar, Button } from "antd";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import useAxiosPrivate from "./hooks/useAxiosPrivate";

const Navbar = () => {
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [userData, setUserData] = useState<any>(null);

    const navigate = useNavigate();
    const axiosPrivate = useAxiosPrivate();

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await axiosPrivate.get('/strava/user', { withCredentials: true });
                return setUserData(response.data);
            } catch (err) {
                console.error(err)
                return null;
            } finally {
                setIsLoading(false);
            }
        };

        fetchUser();
    }, []);

    return (
        <Header style={{
            position: 'sticky',
            top: 0,
            zIndex: 1,
            height: '3rem',
            width: '100%',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-between'
        }}>
            {isLoading ? "Loading" : (
                <>
                    <Button
                        type="primary"
                        size="small"
                        onClick={() => navigate('/dashboard')}
                    >Dashboard</Button>

                    <Avatar style={{ "cursor": "pointer", marginLeft: '1em' }} src={<img src={userData?.profileMediumUrl} alt="avatar" onClick={() => navigate('/profile')} />} />
                </>
            )}
        </Header>
    )
};

export default Navbar;