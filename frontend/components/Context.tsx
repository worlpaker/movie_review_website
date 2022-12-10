import jwt_decode from "jwt-decode";

type Info = {
    First_name: string,
    Last_name: string,
    Email: string,
    Birth_Date: number,
    Gender: string,
    Profile_picture: string;
};

const UserContext = () => {
    const isLoggedin = localStorage.getItem("access");
    if (!isLoggedin) {
        return false;
    }
    const user: Info = jwt_decode(isLoggedin);

    return user;
};
export default UserContext;