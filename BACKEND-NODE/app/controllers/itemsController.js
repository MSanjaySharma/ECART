const Item = require("../models/Item");

const itemsController = {};

itemsController.createItem = (req, res) => {
  const { name } = req.body;
  const item = new Item({ name });
  console.log("ITEM CONTROLLER");
  item
    .save()
    .then((item) => {
      res.json(item);
    })
    .catch((err) => {
      return res.status(400).json({ error: errorHandler(err) });
    });
};

itemsController.getAllItems = (req, res) => {
  Item.find({})
    /* .populate("categories", "_id name slug")
    .populate("tags", "_id name slug")
    .populate("postedBy", "_id name username")
    .select(
      "_id title slug excerpt categories tags postedBy createdAt updatedAt likes"
    ) */
    .exec()
    .then((data) => res.json(data))
    .catch((err) => res.status(400).json({ error: errorHandler(err) }));
};

module.exports = itemsController;
