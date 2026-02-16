// SPDX-FileCopyrightText: 2026 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

import path from "path";
import TerserPlugin from "terser-webpack-plugin";
import webpack from "webpack";
import TransformPlugin from "../../transform-plugin";

const target = "es2017";

const config: webpack.Configuration = {
  mode: "production",
  // mode: "development",
  // devtool: false,
  entry: "./src/index.ts",
  output: {
    filename: "index.js",
    path: path.resolve(__dirname, "dist"),
    chunkFormat: "commonjs",
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
  },
  externals: {
    url: "url",
  },
  plugins: [
    new webpack.DefinePlugin({ global: "globalThis" }),
    new TransformPlugin({
      test: /\.(?:js)$/,
      transform: (source: string) => {
        const SUFFIX = ";";
        if (source[source.length - 1] === SUFFIX) {
          source = source.substring(0, source.length - SUFFIX.length);
        }
        return "((module,__fetch)=>{" + source + "})";
      },
    }),
  ],
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
  target: [target, "node"],
};

export default config;
