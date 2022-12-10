import axios from "axios";
import { useEffect, useState } from "react";
import UserContext from "../components/Context";
import Message from "../components/message";

// Private watchlist show to user.

type Data = {
  Email: string,
  Movie_name: string[];
};

const myarray: string[] = [];

const Toggle = (props: string) => {

  const handleChange = (e: any) => {
    let isChecked = e.target.checked;
    let name = e.target.name;
    (isChecked ? myarray.push(name) : myarray.indexOf(name) !== -1 && myarray.splice(myarray.indexOf(name), 1));
    console.log(myarray);
  };

  return (
    <>
      <div className="container">Status{" "}
        <div className="toggle-switch">
          <input type="checkbox" className="checkbox"
            name={props} id={props} onChange={handleChange} />
          <label className="label" htmlFor={props} >
            <span className="inner" />
            <span className="switch" />
          </label>
        </div>
      </div>
    </>
  );
};


const Watchlist = () => {
  const [messagestate, setMessagestate] = useState<string>("");
  const [watchlist_data, setWatchlist_data] = useState<any>([]);
  const [user, setUser] = useState<any>([]);
  const [isLoggedin, setisLoggedin] = useState<boolean>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    setLoading(true);
    const Data = UserContext();
    const check = Data ? true : false;
    setisLoggedin(check);
    if (check) {
      setUser(Data);
      const Watch_List = async () => {
        await axios
          .post("/api/show_watchlist", Data)
          .then(res => setWatchlist_data(res.data))
          .catch(err => console.log(err));
      };
      Watch_List();
    }
    setLoading(false);

  }, []);

  const Post_Data = async (myData: Data) => {
    await axios
      .post('/api/delete_watchlist', myData)
      .then(res => (Success_Post(res.data)))
      .catch(err => (Error_Post(err)));
  };

  const Success_Post = (res: any) => {
    setMessagestate("success");
    setTimeout(() => {
      setMessagestate("");
    }, 2000);
  };

  const Error_Post = (err: any) => {
    console.log(err);
    setMessagestate("error");
    setTimeout(() => {
      setMessagestate("");
    }, 2000);
  };

  const handleSave = () => {
    const myData: Data = {
      Email: user.Email,
      Movie_name: myarray
    };

    Post_Data(myData);
  };

  return (
    <>
      < div className="watch">
        <div className="help-text">After you turn them Yes, kindly Click Save button</div>
        <div className="button-nav" ><button onClick={handleSave}>Save</button></div>

        {isLoggedin && !loading && watchlist_data !== null && watchlist_data.map((data_watchlist: any, index: any) =>
          <div className="cardbody" key={index}>
            <span className="cardbody-text">
              {data_watchlist.Movie_name}
            </span>
            <div className="cardbody-status">{Toggle(data_watchlist.Movie_name)}</div>
          </div>
        )
        }

        {
          messagestate === "success" &&
          <Message type="success" place="right_bottom" mytext="Successfully edit watchlist." />
        }
        {
          messagestate === "error" &&
          <Message type="error" place="right_bottom" mytext="Invalid Credentials, Please try again" />
        }
      </div>


    </>
  );
};
export default Watchlist;