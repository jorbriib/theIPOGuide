const MiniCssExtractPlugin = require("mini-css-extract-plugin");

module.exports = {
  test: /\.css$/,
  include: /node_modules/,
  use: [MiniCssExtractPlugin.loader, "css-loader"],
};
