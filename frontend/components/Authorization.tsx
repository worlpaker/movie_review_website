import axios from 'axios';
import { useEffect } from 'react';


const Authorization = ({ children }: any) => {

    const Error_Refresh = (err: any) => {
        console.log(err);
        localStorage.removeItem("access");
        window.location.href = "/";
    };

    const Success_Refresh = (res: any) => {
        localStorage.setItem("access", res);
    };

    useEffect(() => {
        const isLoggedin = localStorage.getItem("access");
        const AuthCheck = async () => {
            await axios('/api/profile/refresh/', {
                method: 'POST',
                withCredentials: true,
            })
                .then(res => (Success_Refresh(res.data)))
                .catch(err => (Error_Refresh(err)));
        };
        if (!isLoggedin) {
            console.log("Not logged in");
            return;
        }
        AuthCheck();
    }, []);
    return (
        <>
            {children}
        </>
    );

};

export default Authorization;