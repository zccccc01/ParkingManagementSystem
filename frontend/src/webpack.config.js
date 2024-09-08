// webpack.config.js
module.exports = {
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['react-app'],
            plugins: ['@babel/plugin-proposal-private-property-in-object']
          }
        }
      }
    ]
  }
};