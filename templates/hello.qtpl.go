// Code generated by qtc from "hello.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// All text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//
// Hello is a simple template function.

//line templates/hello.qtpl:5
package templates

//line templates/hello.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/hello.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/hello.qtpl:5
func StreamGenerateStarterCode(qw422016 *qt422016.Writer, url string, method string, k6Options string, body string, headers string, files string) {
//line templates/hello.qtpl:5
	qw422016.N().S(`
import http from 'k6/http';
import { sleep, check } from 'k6';
import { Counter } from 'k6/metrics';
`)
//line templates/hello.qtpl:9
	if len(files) > 0 {
//line templates/hello.qtpl:9
		qw422016.N().S(`
import { FormData } from 'https://jslib.k6.io/formdata/0.0.2/index.js';
`)
//line templates/hello.qtpl:11
	}
//line templates/hello.qtpl:11
	qw422016.N().S(`

// A simple counter for http requests

export const requests = new Counter('http_reqs');

// you can specify stages of your test (ramp up/down patterns) through the options object
// target is the number of VUs you are aiming for

export const options = `)
//line templates/hello.qtpl:20
	qw422016.N().V(k6Options)
//line templates/hello.qtpl:20
	qw422016.N().S(`;

export default function () {
  // our HTTP request, note that we are saving the response to res, which can be accessed later

  `)
//line templates/hello.qtpl:25
	if len(files) > 0 {
//line templates/hello.qtpl:25
		qw422016.N().S(`
    `)
//line templates/hello.qtpl:26
		qw422016.N().V(files)
//line templates/hello.qtpl:26
		qw422016.N().S(`
  `)
//line templates/hello.qtpl:27
	}
//line templates/hello.qtpl:27
	qw422016.N().S(`
  

  `)
//line templates/hello.qtpl:30
	if len(body) > 0 {
//line templates/hello.qtpl:30
		qw422016.N().S(` 
    `)
//line templates/hello.qtpl:31
		if len(files) > 0 {
//line templates/hello.qtpl:31
			qw422016.N().S(`
      const bodyRequest = fd.body();
    `)
//line templates/hello.qtpl:33
		} else {
//line templates/hello.qtpl:33
			qw422016.N().S(`
      const bodyRequest = `)
//line templates/hello.qtpl:34
			qw422016.N().V(body)
//line templates/hello.qtpl:34
			qw422016.N().S(` 
    `)
//line templates/hello.qtpl:35
		}
//line templates/hello.qtpl:35
		qw422016.N().S(`
  `)
//line templates/hello.qtpl:36
	} else {
//line templates/hello.qtpl:36
		qw422016.N().S(` 
    const bodyRequest = undefined 
  `)
//line templates/hello.qtpl:38
	}
//line templates/hello.qtpl:38
	qw422016.N().S(`
  const headers = `)
//line templates/hello.qtpl:39
	if len(headers) > 0 {
//line templates/hello.qtpl:39
		qw422016.N().S(` { headers: `)
//line templates/hello.qtpl:39
		qw422016.N().V(headers)
//line templates/hello.qtpl:39
		qw422016.N().S(` } `)
//line templates/hello.qtpl:39
	} else {
//line templates/hello.qtpl:39
		qw422016.N().S(` undefined `)
//line templates/hello.qtpl:39
	}
//line templates/hello.qtpl:39
	qw422016.N().S(`

  const res = http.`)
//line templates/hello.qtpl:41
	qw422016.E().S(method)
//line templates/hello.qtpl:41
	qw422016.N().S(`('`)
//line templates/hello.qtpl:41
	qw422016.E().S(url)
//line templates/hello.qtpl:41
	qw422016.N().S(`', bodyRequest, headers);

  sleep(1);

  check(res, {
    'status is 200': (r) => r.status === 200,
    'response body': (r) => r.body.indexOf('Feel free to browse') !== -1,
  });
}
`)
//line templates/hello.qtpl:50
}

//line templates/hello.qtpl:50
func WriteGenerateStarterCode(qq422016 qtio422016.Writer, url string, method string, k6Options string, body string, headers string, files string) {
//line templates/hello.qtpl:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/hello.qtpl:50
	StreamGenerateStarterCode(qw422016, url, method, k6Options, body, headers, files)
//line templates/hello.qtpl:50
	qt422016.ReleaseWriter(qw422016)
//line templates/hello.qtpl:50
}

//line templates/hello.qtpl:50
func GenerateStarterCode(url string, method string, k6Options string, body string, headers string, files string) string {
//line templates/hello.qtpl:50
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/hello.qtpl:50
	WriteGenerateStarterCode(qb422016, url, method, k6Options, body, headers, files)
//line templates/hello.qtpl:50
	qs422016 := string(qb422016.B)
//line templates/hello.qtpl:50
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/hello.qtpl:50
	return qs422016
//line templates/hello.qtpl:50
}
