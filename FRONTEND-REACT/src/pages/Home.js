import React from "react";

import Grid from "@material-ui/core/Grid";
import Chip from "@material-ui/core/Chip";
import Typography from "@material-ui/core/Typography";
import Paper from "@material-ui/core/Paper";
import Button from "@material-ui/core/Button";

import swal from "sweetalert";

import axios from "../config/axios";

function Home() {
  const [items, setItems] = React.useState([]);

  const [cart, setCart] = React.useState(null);
  const [order, setOrder] = React.useState([]);

  React.useEffect(() => {
    axios
      .get("/cart/authList", {
        headers: { Authorization: `${localStorage.getItem("authToken")}` },
      })
      .then((resp) => {
        console.log(resp.data);
        setCart(resp.data);
      })
      .catch((err) => {
        console.log(err);
      });
    axios
      .get("/item/list")
      .then((resp) => {
        setItems(resp.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  const addToCart = (itemId) => {
    axios
      .post(
        "/cart/add",
        { itemId },
        {
          headers: { Authorization: `${localStorage.getItem("authToken")}` },
        }
      )
      .then((resp) => {
        swal("Successfully Added Item to Cart!", "", "success");
      })
      .catch((err) => {
        console.log(err?.response?.data);
        if (err?.response?.data?.msg) {
          swal(JSON.stringify(err?.response?.data?.msg), "", "error");
        } else {
          swal("Error Adding Item to Cart!", "", "error");
        }
      });
  };

  const getCart = () => {
    axios
      .get("/cart/authList", {
        headers: { Authorization: `${localStorage.getItem("authToken")}` },
      })
      .then((resp) => {
        console.log(resp.data);
        setCart(resp.data);
        const cart = resp.data;
        alert(JSON.stringify({ cardId: cart._id, items: cart.items }));
        /* swal(
          JSON.stringify({ cardId: cart._id, items: cart.items }),
          "",
          "success"
        ); */
      })
      .catch((err) => {
        console.log(err?.response?.data);
        if (err?.response?.data?.msg) {
          swal(JSON.stringify(err?.response?.data?.msg), "", "error");
        } else {
          swal("Error Getting Cart!", "", "error");
        }
      });
  };

  const getOrders = () => {
    axios
      .get("/order/list", {
        headers: { Authorization: `${localStorage.getItem("authToken")}` },
      })
      .then((resp) => {
        console.log(resp.data);
        setOrder(resp.data);
        const orders = resp.data.map((i) => i._id);
        //swal("Orders fetched", "", "success");
        alert(JSON.stringify(orders));
      })
      .catch((err) => {
        console.log(err?.response?.data);
        if (err?.response?.data?.msg) {
          swal(JSON.stringify(err?.response?.data?.msg), "", "error");
        } else {
          swal("Error Getting Orders!", "", "error");
        }
      });
  };

  const checkout = () => {
    if (cart) {
      axios
        .get(`/cart/${cart._id}/complete`, {
          headers: { Authorization: `${localStorage.getItem("authToken")}` },
        })
        .then((resp) => {
          console.log(resp.data);
          swal("Order placed successfully", "", "success");
        })
        .catch((err) => {
          console.log(err);
          swal("Error placing Order", "", "error");
        });
    }
  };

  return (
    <>
      <Paper style={{ width: "90%", margin: "2vh auto" }}>
        <Grid container direction="row" justify="center" alignItems="center">
          <Grid item style={{ width: "100%" }}>
            <Typography variant="h6" align="center">
              <Button style={{ margin: "5px" }} variant="outlined">
                <Typography variant="h6" align="center" onClick={checkout}>
                  CHECKOUT
                </Typography>
              </Button>
              <Button
                style={{ margin: "5px" }}
                variant="outlined"
                onClick={getCart}
              >
                <Typography variant="h6" align="center">
                  CART
                </Typography>
              </Button>
              <Button style={{ margin: "5px" }} variant="outlined">
                <Typography variant="h6" align="center" onClick={getOrders}>
                  ORDER HISTORY
                </Typography>
              </Button>
            </Typography>
          </Grid>
        </Grid>
      </Paper>

      <Paper style={{ width: "90%", margin: "3vh auto" }}>
        <Grid item style={{ width: "100%" }}>
          <Typography variant="h5" align="center">
            ITEMS
          </Typography>
        </Grid>
        {items.map((i) => (
          <Chip
            key={i._id}
            label={i.name}
            onClick={() => {
              addToCart(i._id);
            }}
            style={{ margin: "10px" }}
          />
        ))}
      </Paper>
    </>
  );
}

export default Home;
