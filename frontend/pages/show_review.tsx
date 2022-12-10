import axios from "axios";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import UserContext from "../components/Context";


// Show user review when they click it from their profile.

type Data = {
    Movie_name: string | string[],
    Email: string,
};
const Show_reviews = () => {
    const router = useRouter();
    const query = router.query;
    const name = query.name;
    const [review_data, setReview_data] = useState<any>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [data_length, setData_length] = useState<number>();
    const [user, setUser] = useState<any>([]);


    useEffect(() => {
        if (name !== undefined) {
            const Data = UserContext();
            const check = Data ? true : false;
            if (check) {
                setUser(Data);
            }
            const myData: Data =
            {
                Movie_name: name,
                Email: user.Email,
            };
            const Reviews = async () => {
                await axios
                    .post("/api/review_movie_email", myData)
                    .then(res => (setReview_data(res.data), setData_length(res.data.length)))
                    .catch(err => console.log(err));
            };
            if (myData.Email !== undefined) {
                Reviews();
                setLoading(false);
            }
        }

    }, [name, user.Email]);
    return (
        <>
            < div className="watch">

                {!loading && review_data.map((data_review: any, index: any) =>
                    <div key={index}>

                        <h1>You have {data_length} review(s)!</h1>

                        <div className="cardbody">
                            <span className="cardbody-text">
                                {data_review.Movie_review}
                            </span>
                            <div className="cardbody-status" style={{ marginTop: "auto", marginBottom: "auto" }}>-{data_review.Email}</div>

                        </div>
                    </div>
                )
                }
                {loading &&
                    <>
                        <h1>Loading...</h1>
                    </>
                }

            </div>
        </>
    );

};

export default Show_reviews;