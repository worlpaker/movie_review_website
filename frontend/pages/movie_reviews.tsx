import axios from "axios";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

// All reviews by movie.

type Data = {
    Movie_name: string | string[];
};
const Movie_Reviews = () => {
    const router = useRouter();
    const query = router.query;
    const name = query.name;
    const [movies_data, setMovies_data] = useState<any>([]);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        if (name !== undefined) {
            const myData: Data =
            {
                Movie_name: name
            };
            const Reviews = async () => {
                setLoading(true);
                await axios
                    .post("/api/show_reviews_by_movie", myData)
                    .then(res => setMovies_data(res.data))
                    .catch(err => console.log(err));
                setLoading(false);
            };
            Reviews();
        }

    }, [name]);
    return (
        <>
            < div className="watch">
                <h1>{name}</h1>

                {!loading && movies_data.map((data_review: any, index: any) =>
                    <div className="cardbody" key={index}>
                        <span className="cardbody-text">
                            {data_review.Movie_review}
                        </span>
                        <div className="cardbody-status" style={{ marginTop: "auto", marginBottom: "auto" }}>-{data_review.Email}</div>

                    </div>
                )
                }

            </div>
        </>
    );

};

export default Movie_Reviews;