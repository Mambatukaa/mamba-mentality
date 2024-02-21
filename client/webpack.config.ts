import path from 'path';
import { Configuration } from 'webpack';
const HtmlWebpackPlugin = require('html-webpack-plugin');

const config: Configuration = {
  mode:
    (process.env.NODE_ENV as 'production' | 'development' | undefined) ??
    'development',
  entry: './src/entrypoint.tsx',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'dist')
  },
  module: {
    rules: [
      {
        test: /.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/
      },
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader', 'postcss-loader']
      }
    ]
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js']
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: path.resolve(__dirname, 'public/index.html'),
      inject: true
    })
  ]
};

export default config;
