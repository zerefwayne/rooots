
import { Layout } from "antd";
import { Outlet } from "react-router-dom";
import Navbar from "./Navbar";

const { Content } = Layout;

const AppNavLayout = () => {
    return (
        <Layout>
            <Navbar />
            <Content>
                <Outlet />
            </Content>
        </Layout>
    )
}

export default AppNavLayout;
