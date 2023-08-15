import { Alert, Layout } from "antd";
import Sider from "antd/es/layout/Sider";
import { Content } from "antd/es/layout/layout";

const siderStyle: React.CSSProperties = {
    textAlign: 'center',
    lineHeight: '120px',
    color: '#fff',
    backgroundColor: '#3ba0e9',
};

const Dashboard = () => {
    return (
        <Layout>
            <Content>
                <Alert
                    description="Page under construction"
                    type="warning"
                    showIcon
                    closable
                />
                <p>Hello</p>
            </Content>
            <Sider style={siderStyle}></Sider>
        </Layout>
    )
}

export default Dashboard;