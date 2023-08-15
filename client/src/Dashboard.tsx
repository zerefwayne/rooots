import { Layout } from "antd";
import Sider from "antd/es/layout/Sider";
import { Content } from "antd/es/layout/layout";
import { useEffect, useState } from "react";
import useAxiosPrivate from "./hooks/useAxiosPrivate";

const siderStyle: React.CSSProperties = {
    lineHeight: '120px',
    color: '#fff',
    backgroundColor: '#000B15',
    padding: 0
};

const Dashboard = () => {
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const axiosPrivate = useAxiosPrivate();
    const [activities, setActivities] = useState([]);

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await axiosPrivate.get('/strava/activities', { withCredentials: true });
                setActivities(response.data);
            } catch (err) {
                console.error(err)
            } finally {
                setIsLoading(false);
            }
        };

        fetchUser();
    }, []);

    return (
        isLoading
            ? <p>Loading</p>
            : (
                <Layout style={{ height: "93.95vh" }}>
                    <Content />
                    <Sider style={siderStyle} width={300}>
                        <h3 style={{ margin: '1em 0', lineHeight: '24px', textAlign: 'center' }}>Activities</h3>
                        <ul>
                            {activities.length > 0 && activities.map((activity: any) => (<li style={{lineHeight: '24px'}}>{activity?.name}</li>))}
                        </ul>
                    </Sider>
                </Layout>
            )
    )

}

export default Dashboard;