import type { NextPage } from 'next';
import Head from 'next/head';
import { Navigation, A11y, Autoplay as autoplay } from 'swiper';
import { Swiper, SwiperSlide, } from 'swiper/react';
import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/navigation";

import Image from 'next/image';
import { useEffect, useState } from 'react';
import axios from 'axios';
import Link from "next/link";

// Index page: show top rated movies, all reviews.

const Toprated_Movies = () => {
  const [movie_data, setMovie_data] = useState<any>([]);
  const [loading, setLoading] = useState<boolean>(true);
  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      await axios
        .get("/api/show_all_movies")
        .then(res => (setMovie_data(res.data)))
        .catch(err => console.log(err));
      setLoading(false);
    };
    fetchData();
  }, []);

  const Show_Movies = () => {
    const movie_pic = (movie: any) => {
      return `/api/images/movies/${movie}.jpg`;
    };
    return (
      <>
        {!loading && movie_data !== null && movie_data.map((data: any, index: any) =>
          <SwiperSlide key={index}>
            <div className='slides-background'>
              <Link href={{
                pathname: "/movie_reviews",
                query: { name: data.Movie_name }, // the data
              }}>
                <a >
                  <Image loader={() => movie_pic(data.Id)} src={movie_pic(data.Id)} alt="latest movies" layout='fill' className="review-pic" unoptimized={true} priority />
                  <div >{Mov_rate(data.Movie_rate)}</div>
                </a>
              </Link>
            </div>
          </SwiperSlide>
        )}
      </>
    );
  };

  return (
    <>

      <div className='slides-back'>
        <div className='slides-body'>
          <Swiper
            slidesPerView={5}
            spaceBetween={30}
            slidesPerGroup={3}
            autoplay={{ delay: 5000, disableOnInteraction: false }}

            navigation={true}
            modules={[Navigation, A11y, autoplay]}
            className="mySwiper"
          >

            {Show_Movies()}


          </Swiper>
        </div>
      </div>
    </>
  );
};

const Mov_rate = (rate: number) => {
  return (
    <>
      <div className='rate-list'>
        <div className='movie-rate'>
          <div className='star'>
            <Image src="/star.svg" alt="star" width={16} height={16} />
          </div>
          <div style={{ fontSize: "12px" }}>{rate}</div>
        </div>
      </div>
    </>
  );
};

const All_reviews = () => {
  const [review_data, setReview_data] = useState<any>([]);
  const [loading, setLoading] = useState<boolean>(true);
  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      await axios
        .get("/api/show_all_reviews")
        .then(res => (setReview_data(res.data)))
        .catch(err => console.log(err));
      setLoading(false);
    };
    fetchData();
  }, []);

  const Show_Reviews = () => {
    const review_pic = (review: any) => {
      return `/api/images/movies/${review}.jpg`;
    };
    return (
      <>
        {!loading && review_data !== null && review_data.map((data_review: any, index: any) =>

          <div className='movies-background' key={index}>
            <Link href={{
              pathname: "/movie_reviews",
              query: { name: data_review.Movie_name }, // the data
            }}>
              <a>
                <Image loader={() => review_pic(data_review.Movie_picture)} src={review_pic(data_review.Movie_picture)} className="review-pic" alt="latest movies" layout='fill' objectFit='cover' unoptimized={true} priority />
                <div >{Mov_rate(data_review.Movie_rate)}</div>
                <div className='movie-name-box'>
                  <div className='movie-name'>{data_review.Movie_name}</div>
                </div>
              </a>
            </Link>
          </div>

        )}
      </>
    );
  };
  return (
    <>
      <div className='movies-back'>
        {Show_Reviews()}
      </div>
    </>
  );
};

const Home: NextPage = () => {
  return (

    <>
      <Head>
        <title>Movie Review</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>


      {<Toprated_Movies />}

      {<All_reviews />}

    </>

  );
};

export default Home;
