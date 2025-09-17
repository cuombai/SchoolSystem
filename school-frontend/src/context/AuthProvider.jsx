import React, {createContext, useState, useEffect} from "react";


export const AuthContext = createContext();

export const AuthProvider = ({children}) => {
    const [auth, setAuth]= useState({
        token : localStorage.getItem("token") || null,
        role: localStorage.getItem("role") || null,
        userID: localStorage.getItem("userID") || null,
    });


    //Persist auth state to localStorage

    useEffect(()=>{
        if (auth.token) {
            localStorage.setItem("token", auth.token);
            localStorage.setItem("role", auth.role);
            localStorage.setItem("userID", auth.userID);
        } else {
            localStorage.removeItem("token");
            localStorage.removeItem("role");
            localStorage.removeItem("userID");
        }
    }, [auth]);

    const login = ({ token, role, userID}) => {
        setAuth({ token, role, userID});
    };
    
    const logout = () => {
        setAuth({ token: null, role: null, userID : null});
    };

    return (
        <AuthContext.Provider value={{auth, login, logout}}>
            {children}
        </AuthContext.Provider>
    )
}