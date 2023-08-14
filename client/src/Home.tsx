import { Button } from "antd";
import User from "./User";
import useLogout from "./hooks/useLogout";
import { useNavigate } from "react-router-dom";

const Home = () => {
    const logout = useLogout();
    const navigate = useNavigate();

    const handleLogout = async () => {
        console.log("LOGGING OUT!");
        await logout();
        navigate('/login');
    };

    return (
        <>
            <h1>
                <User />
            </h1>
            <div>
                <Button onClick={handleLogout}>Logout</Button>
            </div>
        </>
    )
}

export default Home;