import axios from "axios";
import Router from "next/router";
import { useState, useEffect } from "react";
import UserContext from "../components/Context";
import Message from "../components/message";

// Update Account Photo.

const Update_photo = () => {
    const formData = new FormData();
    const [myimage, setMyimage] = useState<any>();

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

    const fileChange = (e: any) => {
        setMyimage(e.target.files[0]);
    };
    const handleClick = () => {
        formData.append("myFile", myimage);
        formData.append("user_profile", user.Profile_picture);
        Post_Photo(formData);
    };
    const Post_Photo = async (photo_data: any) => {
        await axios
            .post('/api/uploadphoto', photo_data)
            .then(res => (Success_Post(res.data)))
            .catch(err => (Error_Post(err)));
    };

    const Success_Post = (data: any) => {
        console.log(data);
        setMessagestate("success");
        setTimeout(() => {
            setMessagestate("");
            Router.push("/");
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
                <div className="login-background" style={{ height: "100px" }}>
                    <label style={{ marginLeft: "20px" }} >Profile Picture:
                        <input type="file" id="myFile" name="myFile" accept="image/*" onChange={fileChange} />
                    </label>
                    <button className="register-button" style={{ marginRight: "15px" }} onClick={handleClick}>Submit</button>

                </div>
            }
            {!isLoggedin && !loading &&
                <>
                    <h1 style={{ textAlign: "center", color: "red" }}>Error</h1>
                    <Message type="error" place="right_bottom" mytext="Authorization Error. Please Log in." />;
                </>
            }
            {messagestate === "success" &&
                <Message type="success" place="right_bottom" mytext="Photo Successfully Updated" />
            }
            {messagestate === "error" &&
                <Message type="error" place="right_bottom" mytext="Invalid Credentials, Please try again" />
            }
        </>
    );

};
export default Update_photo;