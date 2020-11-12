const webpack = require("webpack");

module.exports = (environment, defaultApiUrl, defaultAppUrl, defaultStaticUrl) => {
  const nodeEnv = process.env.NODE_ENV || environment;
  const targetEnv = process.env.TARGET_ENV || environment;
  const appUrl = process.env.APP_URL || defaultAppUrl;
  const staticUrl = process.env.STATIC_URL || defaultStaticUrl;
  const apiUrl = process.env.API_URL || defaultApiUrl;
  const recaptchaSiteKey = process.env.RECAPTCHA_SITE_KEY;

  if (!nodeEnv || !targetEnv || !appUrl || !apiUrl || !staticUrl || !recaptchaSiteKey) {
    throw new Error(
        `Missing required environment variables NODE_ENV or TARGET_ENV or APP_UR or API_URLL or STATIC_URL or RECAPTCHA_SITE_KEY`
    );
  }
  return new webpack.DefinePlugin({
    "process.env": {
      NODE_ENV: JSON.stringify(nodeEnv),
      TARGET_ENV: JSON.stringify(targetEnv),
      APP_URL: JSON.stringify(appUrl),
      API_URL: JSON.stringify(apiUrl),
      STATIC_URL: JSON.stringify(staticUrl),
      RECAPTCHA_SITE_KEY: JSON.stringify(recaptchaSiteKey)
    },
  });
};
