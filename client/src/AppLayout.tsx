import { Outlet } from "react-router-dom"

const AppLayout = () => {
    return (
        <main className="App">
            <Outlet />
        </main>
    )
}

export default AppLayout;
