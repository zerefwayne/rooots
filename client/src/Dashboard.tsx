import { EyeTwoTone } from '@ant-design/icons';
import { Button, Layout, List } from "antd";
import Sider from "antd/es/layout/Sider";
import { Content } from "antd/es/layout/layout";
import { useEffect, useState } from "react";
import useAxiosPrivate from "./hooks/useAxiosPrivate";

const siderStyle: React.CSSProperties = {
    lineHeight: '120px',
    color: '#fff',
    backgroundColor: '#000B15',
    overflowY: 'scroll',
    padding: '1em',
    display: 'block'
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
        <Layout style={{ height: "95vh" }}>
            <Content style={{backgroundColor: '#111'}}>
                </Content>
            <Sider style={siderStyle} width={350}>
                {isLoading ? (<p>Loading activities</p>) : (<>
                    <h3 style={{ margin: '1em 0', lineHeight: '24px', textAlign: 'center' }}>Activities</h3>
                    <List
                        size="small"
                        dataSource={activities}
                        style={{ marginBottom: '1em' }}
                        bordered
                        renderItem={
                            (item: any) => (
                                <List.Item
                                    style={{ color: 'white', borderBottom: '0.5px solid #333', display: 'flex', justifyContent: 'space-between' }}>
                                    <p style={{ margin: 0 }}>{item?.name}</p>
                                    <Button size="small" shape="circle" icon={<EyeTwoTone />} />
                                </List.Item>
                            )
                        }
                    />
                </>
                )}
            </Sider>
        </Layout >
    )

}

export default Dashboard;