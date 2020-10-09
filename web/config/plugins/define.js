const webpack = require("webpack");

module.exports = (environment) => {
  const nodeEnv = process.env.NODE_ENV || environment;
  const targetEnv = process.env.TARGET_ENV || environment;

  if (!nodeEnv || !targetEnv) {
    throw new Error(
        `Missing required environment variables NODE_ENV or TARGET_ENV`
    );
  }
  return new webpack.DefinePlugin({
    "process.env": {
      NODE_ENV: JSON.stringify(nodeEnv),
      TARGET_ENV: JSON.stringify(targetEnv),
    },
  });
};
