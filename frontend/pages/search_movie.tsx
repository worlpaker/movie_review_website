import axios from "axios";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

// Search Movies by name. 

type Data = {
    Movie_name: string | string[];
};
const Search_Movies = () => {
    const router = useRouter();
    const query = router.query;
    const name = query.name;
    const [movies_data, setMovies_data] = useState<any>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [data_length, setData_length] = useState<number>();

    useEffect(() => {
        if (name !== undefined) {
            const myData: Data =
            {
                Movie_name: name
            };
            const Reviews = async () => {
                setLoading(true);
                await axios
                    .post("/api/search_movies", myData)
                    .then(res => (setMovies_data(res.data), setData_length(res.data.length)))
                    .catch(err => console.log(err));
                setLoading(false);
            };
            Reviews();
        }

    }, [name]);
    return (
        <>
            < div className="watch">
                <h1>Found {data_length} movie(s)!</h1>

                {!loading && movies_data.map((data_search: any, index: any) =>
                    <div className="cardbody" key={index}>
                        <span className="cardbody-text">
                            Movie name: {data_search.Movie_name}({data_search.Movie_year}) - Rate: {data_search.Movie_rate} - Type: {data_search.Movie_type} - Category: {data_search.Movie_cat}
                        </span>
                    </div>
                )
                }

            </div>
        </>
    );

};

export default Search_Movies;