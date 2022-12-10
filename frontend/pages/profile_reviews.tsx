import axios from "axios";
import Image from "next/image";
import Link from "next/link";
import { useEffect, useState } from "react";
import UserContext from "../components/Context";

// Reviews by user profile. Includes edit review.

const Mov_rate = (rate: number) => {
  return (
    <>
      <div className='rate-list'>
        <div className='movie-rate'>
          <div className='star'>
            <Image src="/star.svg" alt="star" width={16} height={16} />
          </div>
          {rate}
        </div>
      </div>
    </>
  );
};

const Edit_icon = (movie_name: string) => {

  return (
    <>
      <div className="edit_icon">
        <Link href={{
          pathname: "/edit_review",
          query: { name: movie_name }, // the data
        }}>
          <a>
            <Image
              src="/edit.svg"
              alt="edit"
              width={20}
              height={20}
            />
          </a>
        </Link>

      </div>
    </>
  );
};

const All_reviews = () => {

  const [user, setUser] = useState<any>([]);
  const [isLoggedin, setisLoggedin] = useState<boolean>();

  const [review_data, setReview_data] = useState<any>([]);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const Page = async () => {
      setLoading(true);
      const Data = UserContext();
      const check = Data ? true : false;
      setisLoggedin(check);
      if (check) {
        setUser(Data);
        const Reviews = async () => {
          await axios
            .post("/api/review_by_email", Data)
            .then(res => setReview_data(res.data))
            .catch(err => console.log(err));
        };
        Reviews();
      }
      setLoading(false);
    };
    Page();
  }, []);

  const review_pic = (review: any) => {
    return `/api/images/movies/${review}.jpg`;
  };

  return (
    <>
      {isLoggedin && !loading && review_data !== null && review_data.map((data_review: any, index: any) =>

        <div className='movies-background-profile' key={index}>
          <Link href={{
            pathname: "show_review",
            query: { name: data_review.Movie_name },
          }}>
            <a>
              <Image loader={() => review_pic(data_review.Movie_picture)} src={review_pic(data_review.Movie_picture)} className="review-pic" alt="latest movies" objectFit='cover' layout='fill' unoptimized={true} priority />
            </a>
          </Link>
          <div>{Mov_rate(data_review.Movie_rate)}</div>
          <div>{Edit_icon(data_review.Movie_name)}</div>

          <div className='movie-name-box'>
            <div className='movie-name'>{data_review.Movie_name}</div>
          </div>
        </div>

      )}
    </>
  );
};

const Reviews = () => {
  return (
    <>
      <div className='movies-back-profile'>
        {All_reviews()}
      </div>

    </>
  );
};

export default Reviews;