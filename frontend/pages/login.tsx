import { useRef, useState } from 'react';
import axios from 'axios';
import Message from '../components/message';

// User Login

type Data = {
    Email: string,
    Password: string,
};

const Login = () => {
    const Email = useRef<any>(null);
    const Password = useRef<any>(null);
    const [messagestate, setMessagestate] = useState<string>("");

    const Post_Login = async (logindata: any) => {
        await axios('/api/login/', {
            method: 'POST',
            withCredentials: true,
            data: logindata,
        })
            .then(res => (Success_Login(res.data)))
            .catch(err => (Error_Login(err)));
    };


    const handleClick = () => {
        const myData: Data = {
            Email: Email.current?.value,
            Password: Password.current?.value,
        };
        Post_Login(myData);
    };

    const Success_Login = (res: any) => {
        localStorage.setItem("access", res);
        setMessagestate("success");
        setTimeout(() => {
            setMessagestate("");
            window.location.href = "/";
        }, 2000);

    };
    const Error_Login = (err: any) => {
        console.log(err);
        setMessagestate("error");
        setTimeout(() => {
            setMessagestate("");
        }, 2000);
    };
    return (
        <>
            <div className="login-background">
                <div className="login-form">
                    <label>Email: <input style={{ float: "right" }} type="email" id="Email" ref={Email}></input></label>
                    <label>Password: <input className='input-content' type="password" id="Password" ref={Password} ></input></label>
                    <button className="register-button" onClick={handleClick}>Login</button>
                </div>
            </div>

            {messagestate === "success" &&
                <Message type="success" place="right_bottom" mytext="Successfully Logged in." />
            }
            {messagestate === "error" &&
                <Message type="error" place="right_bottom" mytext="Invalid Credentials, Please try again" />
            }

        </>
    );
};
export default Login;