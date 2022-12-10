import axios from 'axios';
import { useEffect, useState } from 'react';
import UserContext from '../components/Context';
import Message from '../components/message';

// Update account informations

type Data = {
    First_name: string,
    Last_name: string,
    Email: string,
    Birth_Date: number,
    Gender: string;
};

const Update_account = () => {


    const [messagestate, setMessagestate] = useState<string>("");
    const [user, setUser] = useState<any>([]);
    const [isLoggedin, setisLoggedin] = useState<boolean>();
    const [loading, setLoading] = useState<boolean>(true);
    const [values, setValues] = useState<any>();


    useEffect(() => {
        const Page = async () => {
            setLoading(true);
            const Data = UserContext();
            const check = Data ? true : false;
            setisLoggedin(check);
            if (check) {
                setUser(Data);
                setValues(Data);
            }
            setLoading(false);
        };
        Page();
    }, []);




    const handleClick = () => {
        const myData: Data = {
            First_name: values.First_name,
            Last_name: values.Last_name,
            Email: values.Email,
            Birth_Date: values.Birth_Date,
            Gender: values.Gender
        };
        Post_Data(myData);
    };

    const handleInputChange = (e: any) => {
        const { name, value } = e.target;
        setValues({
            ...values,
            [name]: value,
        });
    };
    const Post_Data = async (data: Data) => {
        await axios('/api/profile/updateprofile', {
            method: 'PUT',
            headers: {
                "Content-Type": "application/json"
            },
            withCredentials: true,
            data: data,
        })
            .then(res => (Success_Post(res.data)))
            .catch(err => (Error_Post(err)));
    };

    const Error_Post = (err: any) => {
        console.log(err);
        setMessagestate("error");
        setTimeout(() => {
            setMessagestate("");
        }, 2000);
    };

    const Success_Post = (res: any) => {
        localStorage.setItem("access", res);
        setMessagestate("success");
        setTimeout(() => {
            setMessagestate("");
            window.location.href = "/";
        }, 2000);
    };


    return (
        <>
            {isLoggedin && !loading &&

                <div className="register-background">

                    <div className="register-form">
                        <label>First Name: <input style={{ marginLeft: "8px" }} type="text" id="First_Name" name="First_name" value={values.First_name} onChange={handleInputChange}></input></label>
                        <label>Last Name: <input style={{ marginLeft: "8px" }} type="text" id="Last_Name" name="Last_name" value={values.Last_name} onChange={handleInputChange}></input></label>
                        <label>Email:<input style={{ marginLeft: "43px" }} type="email" id="Email" name="Email" value={values.Email} onChange={handleInputChange}></input></label>
                        <label>Birth Date: <input style={{ marginLeft: "12px" }} type="date" id="Birth_Date" aria-describedby="date-format" min="1900-01-01" max="2022-01-01" name="Birth_Date" value={values.Birth_Date} onChange={handleInputChange} /></label>
                        <label >Gender:
                            &nbsp;
                            <select id="Gender" name="Gender" value={values.Gender} onChange={handleInputChange}>
                                <option>Female</option>
                                <option>Male</option>
                                <option>I do not want to specify</option>
                            </select>
                        </label>

                        <button className="register-button" onClick={handleClick}>Save</button>
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
                <Message type="success" place="right_bottom" mytext="Successfully update account." />
            }
            {messagestate === "error" &&
                <Message type="error" place="right_bottom" mytext="Invalid Credentials, Please try again" />
            }
        </>
    );
};

export default Update_account;