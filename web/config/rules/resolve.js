const path = require("path");

module.exports = {
  extensions: [".js"],
  alias: {
    "ipo/components": path.resolve("./src/components"),
    "ipo/pages": path.resolve("./src/pages"),
  },
};
