const Cart = require("../models/Cart");
const Order = require("../models/Order");
const Item = require("../models/Item");

const cartsController = {};

cartsController.addItemToCart = (req, res) => {
  const { itemId } = req.body; //itemid
  const userId = req.user._id;

  Cart.findOne({ user: userId })
    .exec()
    .then((cart) => {
      /* if (cart.isPurchased) {
        res.json({ msg: "Cart already purchased" });
      } */

      Item.findOne({ _id: itemId })
        .exec()
        .then((i) => {
          Item.updateOne({ _id: itemId }, { $addToSet: { carts: cart._id } })
            .then((i) => {
              //console.log(i);
            })
            .catch((e) => {
              res.json({ msg: "Invalid Update" });
            });
        })
        .catch((err) => {
          res.json({ msg: "Invalid item" });
        });
      Cart.updateOne({ _id: cart._id }, { $addToSet: { items: itemId } })
        .then((c) => {
          res.json({ msg: "Item Added to cart" });
        })
        .catch((err) => {
          res.json(err);
        });
    })
    .catch((error) => {
      res.json(err);
    });
};

cartsController.convertCartToOrder = (req, res) => {
  const userId = req.user._id;
  const cartId = req.params.cartId;

  Cart.findOne({ _id: cartId })
    .then((c) => {
      /* if (c.isPurchased) {
        res.json({ msg: "Cart already purchased" });
      } */
      if (c.items.length === 0) {
        res.status(400).json({ msg: "Cart is Empty" });
      }

      Cart.updateOne({ _id: cartId }, { isPurchased: true })
        .then((c) => {
          const order = new Order({ user: userId, cart: cartId });
          order
            .save()
            .then((o) => {
              res.json(o);
              /* const cart = new Cart();
              cart.user = userId;
              cart
                .save()
                .then((c) => {
                  res.json(o);
                  res.json({ msg: "Order successfull", newCart: { ...c } });
                })
                .catch((err) => {
                  res.json(err);
                }); */
            })
            .catch((err) => {
              res.status(400).json({ msg: err });
            });
        })
        .catch((err) => {
          res.json({ msg: "Order Create Fail" });
        });
    })
    .catch((err) => {
      res.json({ msg: "Invalid CartId" });
    });
};

cartsController.getAllCarts = (req, res) => {
  Cart.find({})
    .exec()
    .then((data) => res.json(data))
    .catch((err) => res.status(400).json({ error: errorHandler(err) }));
};

cartsController.getAuthenticatedUsersCart = (req, res) => {
  const userId = req.user._id;
  Cart.findOne({ user: userId })
    .populate("items", "_id name")
    .exec()
    .then((data) => res.json(data))
    .catch((err) => res.status(400).json({ error: errorHandler(err) }));
};

module.exports = cartsController;
