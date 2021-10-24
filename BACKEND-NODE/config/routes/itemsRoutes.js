const express = require("express");

const itemsController = require("../../app/controllers/itemsController");

const router = express.Router();

router.post("/create", itemsController.createItem);
router.get("/list", itemsController.getAllItems);

module.exports = router;
