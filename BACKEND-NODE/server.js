const express = require("express");
const morgan = require("morgan");
const bodyParser = require("body-parser");
const cookieParser = require("cookie-parser");
const cors = require("cors");
require("dotenv").config();

const configureDB = require("./config/database");
const itemsRoutes = require("./config/routes/itemsRoutes");
const cartsRoutes = require("./config/routes/cartsRoutes");
const usersRoutes = require("./config/routes/usersRoutes");
const ordersRoutes = require("./config/routes/ordersRoutes");

//app
const app = express();

//db
configureDB();

//middlewares
app.use(morgan("dev"));
app.use(bodyParser.json());
app.use(cookieParser());

//cors
app.use(cors({ origin: "*" }));

//routes middlewares
app.use("/user", usersRoutes);
app.use("/cart", cartsRoutes);
app.use("/item", itemsRoutes);
app.use("/order", ordersRoutes);

//port
const port = process.env.PORT || 7331;

app.listen(port, () => {
  console.log(
    `\x1b[94mServer is running on port ${port}\x1b[39m\n\x1b[94mvisit\x1b[39m \x1b[96mhttp://localhost:7331/\x1b[39m`
  );
});
