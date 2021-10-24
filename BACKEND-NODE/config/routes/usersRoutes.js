const express = require("express");

//controllers
const usersController = require("../../app/controllers/usersController");

//validators
const { runValidation } = require("../validators");
const usersValidator = require("../validators/usersValidator");

const router = express.Router();

//REGISTER
router.post(
  "/create",
  usersValidator.register,
  runValidation,
  usersController.register
);

//LOGIN
router.post(
  "/login",
  usersValidator.login,
  runValidation,
  usersController.login
);

// GET ALL USERS
router.get("/list", usersController.getAllUsers);

module.exports = router;
