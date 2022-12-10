import { useRef, useState } from 'react';
import Message from '../components/message';
import axios from 'axios';
import Router from "next/router";

// Signup Page

type Data = {
    First_name: string,
    Last_name: string,
    Email: string,
    Password: string,
    Birth_Date: number,
    Gender: string;
};

const Register = () => {
    const [messagestate, setMessagestate] = useState<string>("");

    const First_name = useRef<any>(null);
    const Last_name = useRef<any>(null);
    const Email = useRef<any>(null);
    const Password = useRef<any>(null);
    const Birth_Date = useRef<any>(null);
    const Gender = useRef<any>(null);

    const formData = new FormData();
    const [myimage, setMyimage] = useState<any>();

    const fileChange = (e: any) => {
        setMyimage(e.target.files[0]);
    };

    const handleClick = () => {
        const myData: Data = {
            First_name: First_name.current?.value,
            Last_name: Last_name.current?.value,
            Email: Email.current?.value,
            Password: Password.current?.value,
            Birth_Date: Birth_Date.current?.value,
            Gender: Gender.current?.value
        };
        Post_Register(myData);
    };

    const Post_Register = async (data: any) => {
        await axios
            .post('/api/register', data)
            .then(res => (Form_Success(res.data)))
            .catch(err => (Handle_Errror(err)));
    };

    const Form_Success = (data: any) => {
        formData.append("myFile", myimage);
        formData.append("user_profile", data);
        Post_Photo(formData);
    };
    const Post_Photo = async (photo_data: any) => {
        await axios
            .post('/api/uploadphoto', photo_data)
            .then(res => (Register_Success(res.data)))
            .catch(err => (Handle_Errror(err)));
    };



    const Register_Success = (data: any) => {
        console.log(data);
        setMessagestate("success");
        setTimeout(() => {
            setMessagestate("");
            Router.push("/");
        }, 2000);
    };

    const Handle_Errror = (err: any) => {
        console.log(err);
        setMessagestate("error");
        setTimeout(() => {
            setMessagestate("");
        }, 2000);
    };


    return (
        <>
            <div className="register-background">
                <div className="register-form">
                    <label>First Name: <input type="text" id="First_Name" ref={First_name}></input></label>
                    <label>Last Name: <input type="text" id="Last_Name" ref={Last_name}></input></label>
                    <label>Email: <input style={{ marginLeft: "40px" }} type="email" id="Email" ref={Email}></input></label>
                    <label>Password: &nbsp;<input type="password" id="Password" ref={Password}></input></label>
                    <label>Birth Date: &nbsp;<input type="date" id="Birth_Date" aria-describedby="date-format" min="1900-01-01" max="2022-01-01" ref={Birth_Date} /></label>
                    <label >Gender:
                        &nbsp; <select id="Gender" ref={Gender}>
                            <option>Female</option>
                            <option>Male</option>
                            <option>I do not want to specify</option>
                        </select>
                    </label>
                    <br>
                    </br>
                    <label>Profile Picture:
                        <input type="file" id="myFile" name="myFile" accept="image/*" onChange={fileChange} />
                    </label>
                    <button className="register-button" style={{ marginLeft: "80px" }} onClick={handleClick}>Register</button>
                </div>
            </div>

            {messagestate === "success" &&
                <Message type="success" place="right_bottom" mytext="Successfully Registered. Please Log in." />
            }
            {messagestate === "error" &&
                <Message type="error" place="right_bottom" mytext="Invalid Credentials, Please try again" />
            }
        </>
    );
};

export default Register;