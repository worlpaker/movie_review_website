import { useEffect, useState } from "react";
import Watchlist from "./profile_watchlist";
import Reviews from "./profile_reviews";
import Image from "next/image";
import UserContext from "../components/Context";
import Message from "../components/message";
import Link from "next/link";
import axios from "axios";

// User Profile

const Profile = () => {

    const [Showwatch, setShowwatch] = useState(false);
    const [Showreviews, setShowreviews] = useState(false);

    const [user, setUser] = useState<any>([]);
    const [isLoggedin, setisLoggedin] = useState<boolean>();
    const [loading, setLoading] = useState<boolean>(true);
    const [user_review, setUser_review] = useState<number>();
    const [user_watchlist, setuser_watchlist] = useState<number>();


    useEffect(() => {
        const Page = async () => {
            setLoading(true);
            const Data = UserContext();
            const check = Data ? true : false;
            setisLoggedin(check);
            if (check) {
                setUser(Data);
                const Count_Data = async () => {
                    await axios
                        .post("/api/count_reviews_by_email", Data)
                        .then(res => (setUser_review(res.data.review), setuser_watchlist(res.data.watchlist)))
                        .catch(err => console.log(err));
                };
                Count_Data();
            }
            setLoading(false);
        };
        Page();
    }, []);



    const user_profile = `/api/images/profiles/${user.Profile_picture}.jpg`;

    const handleClick_Watchlist = () => {

        setShowwatch(!Showwatch);
        setShowreviews(false);

    };
    const handleClick_Reviews = () => {
        setShowreviews(!Showreviews);
        setShowwatch(false);

    };

    return (
        <>
            {loading &&
                <>
                    <h1>Loading...</h1>
                </>
            }
            {isLoggedin && !loading &&
                <>
                    <div className="profile_section">
                        <div style={{ float: "right", position: "absolute", display: "flex", marginLeft: "670px", marginTop: "10px" }}>
                            <Link href="/update_photo">
                                <a style={{ cursor: "pointer", fontWeight: "bold", textDecoration: "none" }}>
                                    Update Photo
                                </a>
                            </Link>
                            <Link href="/update_account">
                                <a style={{ marginLeft: "10px", cursor: "pointer", fontWeight: "bold", textDecoration: "none" }}>
                                    Settings
                                </a>
                            </Link>
                            <Link href="/change_password">
                                <a style={{ marginLeft: "10px", cursor: "pointer", fontWeight: "bold", textDecoration: "none" }}>
                                    Change Password
                                </a>
                            </Link>
                            <Link href="/logout">
                                <a style={{ marginLeft: "10px", cursor: "pointer", fontWeight: "bold", textDecoration: "none" }}>
                                    Logout
                                </a>
                            </Link>
                        </div>
                        <div className="profile_image">
                            <Image
                                src={user_profile}
                                alt="user_profile"
                                layout="fill"
                                className="profile_image"
                                unoptimized={true}
                                priority
                            />
                        </div>
                        <div className="profile_header">PROFILE DETAILS</div>
                        <div className="profile_details_b">
                            <div className="profile_details">
                                <span>{user_review}</span> Reviews
                            </div>
                            <div className="profile_details">
                                <span>{user_watchlist}</span> Watchlist
                            </div>
                            <div style={{ float: "left", position: "absolute", marginTop: "80px", marginLeft: "40px", fontWeight: "bold" }}>
                                {user.First_name.toUpperCase()} {user.Last_name.toUpperCase()}
                            </div>
                            <div className="reviews" onClick={handleClick_Reviews}>
                                <a className="a-profile" href="#"> REVIEWS</a>
                            </div>
                            <div className="watchlist" onClick={handleClick_Watchlist}>
                                <a className="a-profile" href="#"> WATCHLIST</a>
                            </div>
                        </div>
                    </div>
                    <div>
                    </div>
                    <div className="content_last">
                        {Showwatch && !Showreviews && <Watchlist />}
                        {Showreviews && !Showwatch && <Reviews />}
                    </div>
                </>
            }

            {!isLoggedin && !loading &&
                <>
                    <h1 style={{ textAlign: "center", color: "red" }}>Error</h1>
                    <Message type="error" place="right_bottom" mytext="Authorization Error. Please Log in." />;
                </>
            }
        </>
    );
};

export default Profile;