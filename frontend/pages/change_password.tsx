import axios from 'axios';
import { useEffect, useRef, useState } from 'react';
import UserContext from '../components/Context';
import Message from '../components/message';


// Change password

type Data = {
    Email: string,
    Old_Password: string,
    New_Password: string;
};

const Change_password = () => {

    const Email = useRef<any>(null);
    const Old_Password = useRef<any>(null);
    const New_Password = useRef<any>(null);

    const [messagestate, setMessagestate] = useState<string>("");
    const [user, setUser] = useState<any>([]);
    const [isLoggedin, setisLoggedin] = useState<boolean>();
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        const Page = async () => {
            setLoading(true);
            const Data = UserContext();
            const check = Data ? true : false;
            setisLoggedin(check);
            if (check) {
                setUser(Data);
            }
            setLoading(false);
        };
        Page();
    }, []);

    const handleClick = () => {
        const myData: Data = {
            Email: Email.current?.value,
            Old_Password: Old_Password.current?.value,
            New_Password: New_Password.current?.value,
        };
        Post_Data(myData);
    };
    const Post_Data = async (data: Data) => {
        await axios('/api/profile/changepassword', {
            method: 'POST',
            headers: {
                "Content-Type": "application/json"
            },
            withCredentials: true,
            data: data,
        })
            .then(res => (Success_Post(res.data)))
            .catch(err => (Error_Post(err)));
    };

    const Success_Post = (res: any) => {
        localStorage.removeItem("access");
        setMessagestate("success");
        setTimeout(() => {
            setMessagestate("");
            window.location.href = "/login";
        }, 2000);
    };

    const Error_Post = (err: any) => {
        console.log(err);
        setMessagestate("error");
        setTimeout(() => {
            setMessagestate("");
        }, 2000);
    };

    return (
        <>
            {isLoggedin && !loading &&
                <div className="login-background">
                    <div className="login-form">
                        <label>Email: <input style={{ float: "right" }} type="email" id="Email" ref={Email}></input></label>
                        <label>Old Password: <input style={{ float: "right" }} type="password" id="Old_Password" ref={Old_Password} ></input></label>
                        <label>New Password: <input type="password" id="New_Password" ref={New_Password} ></input></label>&nbsp;

                        <button className="register-button" onClick={handleClick}>SAVE</button>
                    </div>
                </div>
            }
            {!isLoggedin && !loading &&
                <>
                    <h1 style={{ textAlign: "center", color: "red" }}>Error</h1>
                    <Message type="error" place="right_bottom" mytext="Authorization Error. Please Log in." />;
                </>
            }

            {messagestate === "success" &&
                <Message type="success" place="right_bottom" mytext="Successfully change password." />
            }
            {messagestate === "error" &&
                <Message type="error" place="right_bottom" mytext="Invalid Credentials, Please try again" />
            }
        </>
    );
};
export default Change_password;