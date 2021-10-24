import axios from "../../config/axios";
import swal from "sweetalert";

export const setUser = (user) => {
  return {
    type: "SET_USER",
    payload: user,
  };
};

export const removeUser = () => {
  return {
    type: "REMOVE_USER",
  };
};

export const startSetUser = (loginData, redirect) => {
  return (dispatch) => {
    axios
      .post("/user/login", loginData)
      .then((response) => {
        if (response.data.hasOwnProperty("errors")) {
          swal(`${response.data.errors}`, "", "error");
        } else {
          swal("Successfully Logged In!", "", "success");
          localStorage.setItem("authToken", response.data.token);
          dispatch(setUser(response.data.user));
          redirect();
        }
      })
      .catch((err) => {
        swal(`${err.response.data.error}`, "", "error");
      });
  };
};

export const startAddUser = (registerData, redirect) => {
  return (dispatch) => {
    axios
      .post("/user/create", registerData)
      .then((response) => {
        if (response.data.errors) {
          swal(`${response.data.msg}`, "", "error");
        } else {
          swal("Successfully Registered!", "", "success");
          redirect();
          dispatch(setUser(response.data.user));
        }
      })
      .catch((err) => {
        swal(`${err.response.data.error}`, "", "error");
      });
  };
};

export const startRemoveUser = () => {
  return (dispatch) => {
    dispatch({ type: "PURGE_USER" });

    localStorage.clear();
    //window.location.replace("http://localhost:3000/signin");
    window.location.replace("http://localhost:3000/signin");
  };
};
