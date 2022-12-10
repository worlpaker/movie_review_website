import axios from "axios";
import { useEffect, useState } from "react";
import UserContext from "../components/Context";

// User Logout

const Logout = () => {
    const [user, setUser] = useState<any>([]);
    const [isLoggedin, setisLoggedin] = useState<boolean>();
    const [success, setSuccess] = useState<boolean>(false);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        setLoading(true);
        const Data = UserContext();
        const check = Data ? true : false;
        setisLoggedin(check);
        if (check) {
            setUser(Data);
        }
        setLoading(false);
        const AuthCheck = async () => {
            await axios('/api/profile/logout/', {
                method: 'POST',
                headers: {
                    "Content-Type": "application/json"
                },
                withCredentials: true,
            })
                .then(res => (Delete_data(res.status)))
                .catch(err => (console.log(err)));
        };

        (check && AuthCheck(), setSuccess(true));

    }, []);

    const Delete_data = (res: any) => {
        console.log(res);
        localStorage.removeItem("access");
    };
    return (
        <>
            {success && !loading &&
                <>
                    {window.location.href = "/"}
                </>
            }
        </>
    );



};

export default Logout;