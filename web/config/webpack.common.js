const path = require("path");

module.exports = {
  entry: ["./src/index.js"],
  output: {
    filename: '[name].bundle.js',
    path: path.resolve(__dirname, "../public"),
  },
  devtool: "source-map",
  module: {
    rules: [
      require("./rules/babelLoader"),
      require("./rules/cssLoader"),
      require("./rules/cssModulesLoader"),
      require("./rules/filesLoader"),
    ],
  },
  optimization: {
    splitChunks: {
      cacheGroups: {
        vendor: {
          test: /node_modules/,
          chunks: "initial",
          name: "vendor",
        },
      },
    },
  },
  plugins: [
    require("./plugins/clean"),
    require("./plugins/chunks2Json"),
    require("./plugins/copy"),
    require("./plugins/html"),

  ],
  resolve: require("./rules/resolve"),
};
