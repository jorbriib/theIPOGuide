const webpack = require("webpack");

module.exports = (environment, defaultApiUrl) => {
  const nodeEnv = process.env.NODE_ENV || environment;
  const targetEnv = process.env.TARGET_ENV || environment;
  const apiUrl = process.env.API_URL || defaultApiUrl;
  const recaptchaSiteKey = process.env.RECAPTCHA_SITE_KEY;

  if (!nodeEnv || !targetEnv || !apiUrl || !recaptchaSiteKey) {
    throw new Error(
        `Missing required environment variables NODE_ENV or TARGET_ENV`
    );
  }
  return new webpack.DefinePlugin({
    "process.env": {
      NODE_ENV: JSON.stringify(nodeEnv),
      TARGET_ENV: JSON.stringify(targetEnv),
      API_URL: JSON.stringify(apiUrl),
      RECAPTCHA_SITE_KEY: JSON.stringify(recaptchaSiteKey)
    },
  });
};
