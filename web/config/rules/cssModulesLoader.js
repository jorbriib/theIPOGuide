const path = require("path");
const postcssPresetEnv = require("postcss-preset-env");
const postcssImportPlugin = require("postcss-import");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const nodeEnv = process.env.NODE_ENV;

module.exports = {
  test: /\.css$/,
  exclude: /node_modules/,
  use: [
    MiniCssExtractPlugin.loader,
    {
      loader: "css-loader",
      options: {
        sourceMap: nodeEnv !== "production",
        modules: {
          localIdentName: "[name]-[local]-[hash:base64:5]",
          exportLocalsConvention: "camelCase",
        },
      },
    },
    {
      loader: "postcss-loader",
      options: {
        postcssOptions: {
          plugins: [
            postcssImportPlugin({
              path: [path.resolve(__dirname, "../../src")],
            }),
            postcssPresetEnv({
              features: {
                "nesting-rules": true,
                "custom-media-queries": true,
                "custom-properties": {
                  preserve: false,
                },
              },
              stage: 2,
              browsers: "> 1%, last 2 versions, Firefox ESR",
            }),
          ],
        },
      },
    },
  ],
};