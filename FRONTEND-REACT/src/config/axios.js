import Axios from "axios";

const url = process.env.URL || "https://ecommerce-ecom.herokuapp.com" || "http://localhost:7331";
const axios = Axios.create({
  //baseURL: "/",
  baseURL: url,
});

export default axios;
