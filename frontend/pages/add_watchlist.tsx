import axios from "axios";
import Router from "next/router";
import { useEffect, useRef, useState } from "react";
import UserContext from "../components/Context";
import Message from "../components/message";


// Add private watchlist, only visible to user.

type Data = {
    Email: string,
    Movie_name: string,
};

const Add_watchlist = () => {

    const Movie_name = useRef<any>(null);
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
            Email: user.Email,
            Movie_name: Movie_name.current?.value,
        };
        Post_Data(myData);
    };
    const Post_Data = async (data: any) => {
        await axios
            .post('/api/add_watchlist', data)
            .then(res => (Success_Post(res.data)))
            .catch(err => (Error_Post(err)));
    };

    const Success_Post = (res: any) => {
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
                <div className="login-background">
                    <div className="login-form">
                        <label>Movie Name:<input type="text" id="Movie_name" ref={Movie_name}></input></label>
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
                <Message type="success" place="right_bottom" mytext="Successfully added a watchlist." />
            }
            {messagestate === "error" &&
                <Message type="error" place="right_bottom" mytext="Invalid Credentials, Please try again" />
            }
        </>
    );
};

export default Add_watchlist;