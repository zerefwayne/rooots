import { Header } from "antd/es/layout/layout";

import { Button } from "antd";
import useLogout from "./hooks/useLogout";
import { useNavigate } from "react-router-dom";

const Navbar = () => {
    const logout = useLogout();
    const navigate = useNavigate();

    const handleLogout = async () => {
        await logout();
        navigate('/login');
    };

    return (
        <Header>
            <Button onClick={handleLogout}>Logout</Button>
        </Header>
    )
};

export default Navbar;