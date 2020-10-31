const { merge } = require("webpack-merge");
const path = require("path")
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

const common = require("./webpack.common.js");

module.exports = merge(common, {
  mode: "development",
  plugins: [
    require("./plugins/define")("development", "http://localhost", "http://localhost"),
    new MiniCssExtractPlugin(),
  ],
  devServer: {
    contentBase: path.join(__dirname, '../public'),
    compress: true,
    port: 9000,
    disableHostCheck: true,
    host: "0.0.0.0",
    historyApiFallback: true,
  },
});
