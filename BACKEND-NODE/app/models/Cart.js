const mongoose = require("mongoose");
const { ObjectId } = mongoose.Schema;

const cartSchema = new mongoose.Schema(
  {
    isPurchased: {
      type: Boolean,
      default: false,
    },
    items: [{ type: ObjectId, ref: "Item" }],
    user: { type: ObjectId, ref: "User", required: true },
  },
  { timestamps: true }
);

const Cart = mongoose.model("Cart", cartSchema);

module.exports = Cart;
