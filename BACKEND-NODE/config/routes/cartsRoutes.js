const express = require("express");

const { authenticateUser } = require("../../app/middlewares/authenticateUser");
const cartsController = require("../../app/controllers/cartsController");

const router = express.Router();

router.post("/add", authenticateUser, cartsController.addItemToCart);
router.get(
  "/:cartId/complete",
  authenticateUser,
  cartsController.convertCartToOrder
);
router.get(
  "/authList",
  authenticateUser,
  cartsController.getAuthenticatedUsersCart
);
router.get("/list", cartsController.getAllCarts);

module.exports = router;
