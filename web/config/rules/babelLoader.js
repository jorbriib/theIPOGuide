module.exports = {
  test: /\.js$/,
  use: {
    loader: "babel-loader",
  },
  exclude: {
    test: /node_modules\/(?!slate|slate-history|slate-react\/)/,
  },
};
