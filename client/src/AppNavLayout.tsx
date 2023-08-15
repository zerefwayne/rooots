
import { Layout } from "antd";
import { Outlet } from "react-router-dom";
import Navbar from "./Navbar";

const AppNavLayout = () => {
    return (
        <Layout>
            <Navbar />
            <Outlet />
        </Layout>
    )
}

export default AppNavLayout;
