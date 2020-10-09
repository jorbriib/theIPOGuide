const Chunks2JsonPlugin = require("chunks-2-json-webpack-plugin");

module.exports = new Chunks2JsonPlugin({
  outputDir: "public/",
  filename: "chunks.json",
  publicPath: "/",
});
