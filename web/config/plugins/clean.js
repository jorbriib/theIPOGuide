const { CleanWebpackPlugin } = require("clean-webpack-plugin");

module.exports = new CleanWebpackPlugin({
  cleanStaleWebpackAssets: false,
  verbose: true,
});
