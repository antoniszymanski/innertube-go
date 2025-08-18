import path from "path";
import TerserPlugin from "terser-webpack-plugin";
import webpack from "webpack";

const target = "es2017";

const config: webpack.Configuration = {
  mode: "production",
  // mode: "development",
  // devtool: false,
  entry: "./youtubei/src/index.ts",
  output: {
    filename: "index.js",
    path: path.resolve(__dirname, "dist"),
    chunkFormat: false,
    library: {
      type: "commonjs2",
    },
  },
  module: {
    rules: [
      {
        test: /\.(?:js|ts)$/,
        loader: "esbuild-loader",
        options: {
          target: target,
        },
      },
    ],
  },
  resolve: {
    extensions: [".js", ".ts"],
    fallback: {
      crypto: require.resolve("crypto-browserify"),
      stream: require.resolve("stream-browserify"),
      vm: require.resolve("vm-browserify"),
    },
  },
  externals: {
    url: "url",
    "node-fetch": "fetch",
  },
  plugins: [new webpack.DefinePlugin({ global: "globalThis" })],
  optimization: {
    chunkIds: "total-size",
    moduleIds: "size",
    minimize: true,
    minimizer: [
      new TerserPlugin({
        terserOptions: {
          format: {
            comments: false,
          },
        },
        extractComments: false,
      }),
    ],
  },
  performance: {
    hints: false,
  },
  target: [target],
};

export default config;
