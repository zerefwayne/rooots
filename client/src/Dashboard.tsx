import { EyeTwoTone } from '@ant-design/icons';
import { Button, Layout, List } from "antd";
import Sider from "antd/es/layout/Sider";
import { Content } from "antd/es/layout/layout";
import { useEffect, useRef, useState } from "react";
import useAxiosPrivate from "./hooks/useAxiosPrivate";
import mapboxgl from "mapbox-gl";

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

    mapboxgl.accessToken =
        "pk.eyJ1IjoiemVyZWZ3YXluZSIsImEiOiJjbGtwcHZ6bTQwOWx5M3B1azB6bmhvN21kIn0.GeFYwb8ZhwGvi8l1ENtHnA";

    const mapContainer = useRef<any>(null);
    const map = useRef<any>(null);

    const [lat, setLat] = useState(46.2271);
    const [lng, setLng] = useState(6.5982);
    const [zoom, setZoom] = useState(7.5);

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

    useEffect(() => {
        if (map.current) return;

        map.current = new mapboxgl.Map({
            container: mapContainer.current,
            style: "mapbox://styles/mapbox/dark-v11",
            center: [lng, lat],
            zoom: zoom,
        });
    });

    useEffect(() => {
        if (!map.current) return;

        map.current.on("move", () => {
            setLng(map.current.getCenter().lng.toFixed(4));
            setLat(map.current.getCenter().lat.toFixed(4));
            setZoom(map.current.getZoom().toFixed(2));
        });
    });

    return (
        <Layout style={{ height: "95vh" }}>
            <Content style={{ backgroundColor: 'red' }}>
                <div ref={mapContainer} className="map-container" style={{ height: '95vh' }} />
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