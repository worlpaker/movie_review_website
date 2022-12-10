import Image from 'next/image';
import mylogo from '../public/logo.png';
import { useEffect, useState } from 'react';
import UserContext from '../components/Context';
import Link from 'next/link';

// Navbar

const Before_Login = () => {
  return (
    <div className="navbar_icons">
      <div className='className="navbar_icon"'>
        <Link href='/'>
          <a>
            <Image src='/home.svg' alt="home" width={24} height={24} />
          </a>
        </Link>
      </div>
      <div className='className="navbar_icon"'>
        <Link href='/login' >
          <a>
            <Image src='/login.svg' alt="login" width={50} height={24} />
          </a>
        </Link>
      </div>
      <div className='className="navbar_icon"'>
        <Link href='/register'>
          <a>
            <Image src='/register.svg' alt="register" width={24} height={24} />
          </a>
        </Link>
      </div>

    </div>
  );
};

const After_Login = (user: any) => {
  const user_profile = `/api/images/profiles/${user.Profile_picture}.jpg`;

  return (
    <>
      <div className='navbar_icons'>
        <div className="user_icon">
          <Link href='/' >
            <a>
              <Image src='/home.svg' alt="home" width={24} height={24} />
            </a>
          </Link>
        </div>
        <div className="user_icon">
          <Link href='/add_movie'>
            <a>
              <Image src='/add_movie.svg' alt="add_movie" width={24} height={24} />
            </a>
          </Link>
        </div>
        <div className="user_icon">
          <Link href='/add_review'>
            <a>
              <Image src='/add_review.svg' alt="add_review" width={24} height={24} />
            </a>
          </Link>
        </div>
        <div className="user_icon">
          <Link href='/add_watchlist'>
            <a>
              <Image src='/add_watchlist.svg' alt="add_watchlist" width={24} height={24} />
            </a>
          </Link>
        </div>
        <div className="user_icon">
          <Link href='/profile'>
            <a>
              <Image loader={() => user_profile} src={user_profile} alt="profile" width={24} height={24} unoptimized={true} style={{ borderRadius: "24px" }} priority />
            </a>
          </Link>
        </div>

      </div>
    </>
  );

};


const Mynavbar = ({ children }: any) => {
  const [searchvalue, setSearchvalue] = useState("");
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

  const handleChange = (e: any) => {
    setSearchvalue(e.target.value);
  };

  const handleSubmit = (e: any) => {
    e.preventDefault();
    return (
      window.location.href = `/search_movie?name=${searchvalue}`
    );

  };

  const handleKeypress = (e: any) => {
    //it triggers by pressing the enter key
    if (e.key === 'Enter') {
      handleSubmit(e);
    }
  };

  return (<>
    <div className="navbar">
      <div className="navbar_brand">
        <span className="navbar_logo">
          <Link href='/' >

            <a>
              <Image
                src={mylogo}
                alt="Website logo"
                width={150}
                height={30}
              />
            </a>
          </Link>
        </span>
      </div>

      <div className="navbar_search">
        <div className="search">
          <div className="search_icon" >
            <Image src="/search.svg" alt="search" width={16} height={16} />
          </div>
          <input type="text" placeholder="Search" className="search_input" onChange={handleChange} onKeyPress={handleKeypress} />
        </div>
      </div>

      {isLoggedin && !loading &&
        <>
          {
            After_Login(user)
          }
        </>
      }

      {!isLoggedin && !loading &&
        <>
          {
            Before_Login()
          }
        </>
      }

    </div>
    <div className="navbar_below">
      {children}
    </div>

  </>
  );
};

export default Mynavbar;