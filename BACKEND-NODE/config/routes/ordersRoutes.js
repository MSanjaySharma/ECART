const express = require("express");

const ordersController = require("../../app/controllers/ordersController");

const { authenticateUser } = require("../../app/middlewares/authenticateUser");

const router = express.Router();

router.get("/list", authenticateUser, ordersController.getAllOrdersForUser);
router.get("/listAll", ordersController.getAllOrders);

module.exports = router;
