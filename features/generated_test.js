import http from "k6/http";
import { sleep, check } from "k6";
import { Counter } from "k6/metrics";

import { FormData } from "https://jslib.k6.io/formdata/0.0.2/index.js";

// A simple counter for http requests

export const requests = new Counter("http_reqs");

// you can specify stages of your test (ramp up/down patterns) through the options object
// target is the number of VUs you are aiming for

export const options = {
  stages: [
    {
      duration: "5m",
      target: 60,
    },
  ],
};

export default function () {
  // our HTTP request, note that we are saving the response to res, which can be accessed later

  const fd = new FormData();
  const file0 = open("/path/to/filename", "image.mp4");
  fd.append("filename", http.file(file0, "image.mp4", "video/mp4"));
  const file1 = open("/path/to/upload", "@image.ppt");
  fd.append(
    "upload",
    http.file(file1, "@image.ppt", "application/vnd.ms-powerpoint")
  );

  const bodyRequest = fd.body();

  const headers = { headers: {} };

  const res = http.get(
    "http://localhost:8080/api/upload",
    bodyRequest,
    headers
  );

  sleep(1);

  const checkRes = check(res, {
    "status is 200": (r) => r.status === 200,
    "response body": (r) => r.body.indexOf("Feel free to browse") !== -1,
  });
}
