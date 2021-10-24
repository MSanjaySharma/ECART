const bcryptjs = require("bcryptjs");
const jwt = require("jsonwebtoken");
const shortId = require("shortid");

const User = require("../models/User");
const Cart = require("../models/Cart");

const usersController = {};

usersController.register = (req, res) => {
  User.findOne({ email: req.body.email }).exec((err, user) => {
    if (user) {
      return res.json({
        error: "Email is taken",
      });
    } else {
      const body = req.body;
      let username = `${shortId.generate()}|${req.body.name
        .split(" ")
        .join("-")}`;
      body.username = username;

      const user = new User(body);
      bcryptjs.genSalt().then((salt) => {
        bcryptjs.hash(user.password, salt).then((encrypted) => {
          user.password = encrypted;
          user
            .save()
            .then((u) => {
              //res.json(user);
              const cart = new Cart();
              cart.user = u._id;
              cart
                .save()
                .then((c) => {
                  res.json({ msg: "Signup success! Please Signin" });
                })
                .catch((err) => {
                  res.status(404).json({ msg: err });
                });
            })
            .catch((err) => {
              res.status(404).json({ msg: err });
            });
        });
      });
    }
  });
};

usersController.login = (req, res) => {
  const body = req.body;
  User.findOne({ email: body.email }).then((user) => {
    if (!user) {
      res.status(404).json({ error: "invalid email or password" });
      return;
    } else {
      bcryptjs.compare(body.password, user.password).then((match) => {
        if (match) {
          const tokenData = {
            _id: user._id,
            email: user.email,
            username: user.username,
          };
          const token = jwt.sign(tokenData, process.env.JWT_SIGN_KEY, {
            expiresIn: "30d",
          });
          res.cookie("token", token, { expiresIn: "30d" });
          res.json({
            token: `Bearer ${token}`,
            user: {
              _id: user._id,
              email: user.email,
              username: user.username,
            },
          });
        } else {
          res.status(404).json({ error: "invalid email or password" });
        }
      });
    }
  });
};

usersController.getAllUsers = (req, res) => {
  User.find({})
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

module.exports = usersController;
