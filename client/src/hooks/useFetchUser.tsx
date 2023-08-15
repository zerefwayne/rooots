import { axiosPrivate } from "../api/axios";

const useFetchUser = () => {
    const fetchUser = async () => {
        try {
            const response = await axiosPrivate.get('/strava/user', {withCredentials: true});
            return response.data;
        } catch (err) {
            console.error(err)
            return null;
        }
    };

    return fetchUser;
};

export default useFetchUser;