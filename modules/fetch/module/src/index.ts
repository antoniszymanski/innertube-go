// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

import { Headers, HeadersList, HeadersObject, headersToList } from "headers-polyfill";

declare function __fetch(req: request): response;

type request = {
  url: string;
  body: string;
  headers: Record<string, string[]>;
  method: string;
};

type response = {
  body: string;
  headers: Record<string, string[]>;
  ok: boolean;
  redirected: boolean;
  status: number;
  statusText: string;
  url: string;
};

type Request = {
  body: string;
  headers: HeadersInit | HeadersObject | HeadersList;
  method: string;
};

// oxlint-disable-next-line no-unused-vars
export async function fetch(url: string, options?: Request) {
  if (options === undefined) {
    options = {
      body: "",
      headers: {},
      method: "GET",
    };
  }

  const req: request = {
    url: url,
    body: options.body,
    method: options.method,
    headers: {},
  };
  // @ts-expect-error
  for (const [key, value] of headersToList(new Headers(options.headers))) {
    if (typeof value === "string") {
      req.headers[key] = [value];
    } else {
      req.headers[key] = value;
    }
  }

  const resp = __fetch(req);
  return new Response(resp);
}

export class Response {
  readonly body: string;
  bodyUsed: boolean;
  readonly headers: Headers;
  readonly ok: boolean;
  readonly redirected: boolean;
  readonly status: number;
  readonly statusText: string;
  readonly url: string;
  constructor(resp: response) {
    this.body = resp.body;
    this.headers = new Headers(resp.headers);
    this.ok = resp.ok;
    this.redirected = resp.redirected;
    this.status = resp.status;
    this.statusText = resp.statusText;
    this.url = resp.url;
  }
  blob(): Promise<Blob> {
    return new Promise((resolve) => {
      const blob = new Blob([this.body]);
      this.bodyUsed = true;
      resolve(blob);
    });
  }
  json(): Promise<any> {
    return new Promise((resolve, reject) => {
      try {
        const json = JSON.parse(this.body);
        this.bodyUsed = true;
        resolve(json);
      } catch (error) {
        reject(error);
      }
    });
  }
  async text() {
    this.bodyUsed = true;
    return this.body;
  }
}

export { Headers } from "headers-polyfill";

export default fetch;
