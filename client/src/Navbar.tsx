import { Header } from "antd/es/layout/layout";

import { HomeOutlined } from '@ant-design/icons';
import { Avatar, Button, Menu, MenuProps } from "antd";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import useFetchUser from "./hooks/useFetchUser";
import MenuItem from "antd/es/menu/MenuItem";

const Navbar = () => {
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [userData, setUserData] = useState<any>(null);

    const navigate = useNavigate();
    const fetchUser = useFetchUser();

    useEffect(() => {
        const getUser = async () => {
            try {
                const response = await fetchUser();
                setUserData(response);
            } catch (err) {
                console.error(err);
            } finally {
                setIsLoading(false);
            }
        };

        getUser();
    }, []);

    return (
        <Header style={{
            position: 'sticky',
            top: 0,
            zIndex: 1,
            width: '100%',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-between'
        }}>
            {isLoading ? "Loading" : (
                <>
                    <Button
                        type="primary"
                        danger={true}
                        onClick={() => navigate('/dashboard')}
                    >Dashboard</Button>

                    <Avatar style={{ "cursor": "pointer", marginLeft: '1em' }} src={<img src={userData?.profileMediumUrl} alt="avatar" onClick={() => navigate('/profile')} />} />

                </>
            )}
        </Header>
    )
};

export default Navbar;