import axios from 'axios';
import Router from 'next/router';
import { useEffect, useRef, useState } from 'react';
import UserContext from '../components/Context';
import Message from '../components/message';

// Add movie to backend.
// Any user can add a new movie, if it doesn't exist.

type Data = {
    Email: string,
    Movie_name: string,
    Movie_type: string,
    Movie_cat: string,
    Movie_year: number,
    Movie_rate: number,
};

const Add_movie = () => {

    const Movie_name = useRef<any>(null);
    const Movie_type = useRef<any>(null);
    const Movie_cat = useRef<any>(null);
    const Movie_year = useRef<any>(null);

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
        const myData: Data = {
            Movie_name: Movie_name.current?.value,
            Movie_type: Movie_type.current?.value,
            Movie_cat: Movie_cat.current?.value,
            Email: user.Email,
            Movie_year: parseInt(Movie_year.current?.value),
            Movie_rate: 0,
        };

        Post_Movie(myData);

    };

    const Post_Movie = async (data: Data) => {
        await axios
            .post('/api/add_movie', data)
            .then(res => (Form_Success(res.data)))
            .catch(err => (Handle_Errror(err)));
    };

    const Form_Success = (data: any) => {
        formData.append("movie_pic", myimage);
        formData.append("movie_id", data);
        Post_Photo(formData);
    };

    const Post_Photo = async (photo_data: any) => {
        await axios
            .post('/api/add_movie_photo', photo_data)
            .then(res => (Post_Success(res.data)))
            .catch(err => (Handle_Errror(err)));
    };

    const Post_Success = (data: any) => {
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
            {isLoggedin && !loading &&
                <div className="login-background">
                    <div className="login-form">
                        <label>Movie Name: <input style={{ float: "right", marginRight: "20px" }} type="text" id="Movie_name" ref={Movie_name}></input></label>
                        <label>Movie Type:
                            <select id="Type" ref={Movie_type} style={{ float: "right", marginRight: "20px", width: "177px" }} >
                                <option>Film</option>
                                <option>TV Series</option>
                            </select>
                        </label>
                        <label>Movie Category:
                            <select id="Category" ref={Movie_cat} style={{ float: "right", marginRight: "20px", width: "177px" }} >
                                <option>Action</option>
                                <option>Dram</option>
                                <option>Sci-Fi</option>
                            </select>
                        </label>
                        <label>Movie year:<input type="number" id="Movie_year" ref={Movie_year} style={{ float: "right", marginRight: "20px" }} ></input></label>
                        <br>
                        </br>
                        <label>Movie Picture:
                            <input type="file" id="movie_pic" name="movie_pic" accept="image/*" onChange={fileChange} style={{ float: "right", marginRight: "20px", width: "177px" }} />
                        </label>
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
                <Message type="success" place="right_bottom" mytext="Successfully added movie." />
            }

            {messagestate === "error" &&
                <Message type="error" place="right_bottom" mytext="Invalid Credentials, Please try again" />
            }
        </>
    );
};
export default Add_movie;