import { createContext, useState } from "react";

export interface Auth {
    accessToken?: string
    name?: string
}

const AuthContext = createContext<{
    auth?: Auth,
    setAuth?: any
}>({});

export const AuthProvider = ({ children }: any) => {
    const [auth, setAuth] = useState({});

    return (
        <AuthContext.Provider value={{ auth, setAuth }}>
            {children}
        </AuthContext.Provider>
    )
}

export default AuthContext;