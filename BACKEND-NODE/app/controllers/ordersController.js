const Order = require("../models/Order");

const ordersController = {};

ordersController.getAllOrders = (req, res) => {
  Order.find({})
    .exec()
    .then((data) => res.json(data))
    .catch((err) => res.status(400).json({ error: errorHandler(err) }));
};

ordersController.getAllOrdersForUser = (req, res) => {
  const userId = req.user._id;

  Order.find({ user: userId })
    /* .populate("categories", "_id name slug")
    .populate("tags", "_id name slug")
    .populate("postedBy", "_id name username")
    .select(
      "_id title slug excerpt categories tags postedBy createdAt updatedAt likes"
    ) */
    //.populate("cart", "_id items")
    .populate({
      path: "cart",
      populate: {
        path: "items",
      },
    })
    .exec()
    .then((data) => {
      data ? res.json(data) : res.status(400).json({ msg: "No orders" });
    })
    .catch((err) => res.status(400).json({ error: errorHandler(err) }));
};

module.exports = ordersController;
