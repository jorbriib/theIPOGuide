const { merge } = require("webpack-merge");
const TerserPlugin = require("terser-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

const common = require("./webpack.common.js");

module.exports = merge(common, {
  mode: "production",
  devtool: "hidden-source-map",
  output: {
    filename: "[name]-[chunkhash:12].js",
    chunkFilename: "[id]-[chunkhash:12].js",
  },
  optimization: {
    minimizer: [
      new TerserPlugin({
        sourceMap: true,
        parallel: true,
        terserOptions: {
          extractComments: "all",
          compress: {
            inline: false,
            drop_console: true,
          },
        },
      }),
    ],
  },
  plugins: [
    require("./plugins/define")("production"),
    new MiniCssExtractPlugin({
      filename: "[name]-[hash:12].css",
      chunkFilename: "[id]-[hash:12].css",
    }),
  ],
});
