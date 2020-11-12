const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = new HtmlWebpackPlugin({
    title: 'theIPOguide',
    template: 'src/index.ejs',
    staticUrl: process.env.STATIC_URL
});
