const mongoose = require("mongoose");
const { ObjectId } = mongoose.Schema;

const itemSchema = new mongoose.Schema(
  {
    name: {
      type: String,
      required: true,
      trim: true,
      minlength: 4,
      maxlength: 160,
    },
    //tags: [{ type: ObjectId, ref: "Tag", required: true }],
    carts: [{ type: ObjectId, ref: "Cart" }],
  },
  { timestamps: true }
);

const Item = mongoose.model("Item", itemSchema);

module.exports = Item;
