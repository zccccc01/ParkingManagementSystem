const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const HappyPack = require('happypack');
require('dotenv').config();

module.exports = {
  entry: './src/index.js',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'dist'),
  },
  module: {
    rules: [
      // 处理 JavaScript 文件
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,
        use: ['happypack/loader?id=babel'],
      },
      // 处理 .module.scss 文件
      {
        test: /\.module\.s[ac]ss$/,
        use: [
          MiniCssExtractPlugin.loader,
          {
            loader: 'css-loader',
            options: {
              modules: {
                localIdentName: '[name]__[local]--[hash:base64:5]',
              },
            },
          },
          'sass-loader',
        ],
      },
      // 处理普通的 .scss 文件
      {
        test: /\.s[ac]ss$/,
        exclude: /\.module\.s[ac]ss$/,
        use: [MiniCssExtractPlugin.loader, 'css-loader', 'sass-loader'],
      },
      // 处理普通的 .css 文件
      {
        test: /\.css$/,
        use: [MiniCssExtractPlugin.loader, 'css-loader'],
      },
    ],
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: '[name].css',
    }),
    new HappyPack({
      id: 'babel',
      loaders: ['babel-loader'],
      threads: 4, // 并发数
    }),
  ],
};
