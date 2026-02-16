// SPDX-FileCopyrightText: 2026 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

import { Compiler, ModuleFilenameHelpers } from "webpack";
import { RawSource } from "webpack-sources";

export interface Args extends MatchObject {
  transform: Transformer;
}

export interface MatchObject {
  test?: Matcher;
  include?: Matcher;
  exclude?: Matcher;
}

export type Matcher = string | RegExp | ((str: string) => boolean) | (string | RegExp | ((str: string) => boolean))[];

export type Transformer = (source: string) => string;

export default class TransformPlugin {
  private matchObject: MatchObject;
  private transform: Transformer;
  constructor(args: Args) {
    this.matchObject = {
      test: args.test,
      include: args.include,
      exclude: args.exclude,
    };
    this.transform = args.transform;
  }

  apply(compiler: Compiler) {
    compiler.hooks.compilation.tap(TransformPlugin.name, (compilation) => {
      compilation.hooks.afterProcessAssets.tap({ name: TransformPlugin.name }, (assets) => {
        for (const [path, source] of Object.entries(assets)) {
          if (!ModuleFilenameHelpers.matchObject(this.matchObject, path)) {
            continue;
          }
          assets[path] = new RawSource(this.transform(source.buffer().toString()));
        }
      });
    });
  }
}
