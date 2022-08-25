# AdamsTaskForce

------------------------
First Task:
1. csinalj egy web alkalmazast ami :8080 PORTon figyel a gepeden es "Hello, world!" el valaszol HTTP -n ha a GET / hivjak rajta

ha ezt lefuttatod:
$ curl -X GET http://localhost:8080/ 

ezt kapod:
Hello, world!

szukseged lesz ezekre a stdlibekre (standard library)
- net/http

ezekrol a komponensekrol fogsz tanulni:
- http.Handler (interface)
- http.ServeMux
- http.ListenAndServe

2. a "/foo" http utvonalat "POST" methodussal hivjak, akkor az alklalmazas jegyezze meg a request.Body tartalmat.
amikor a GET "/foo" -t hivjak rajta, adja vissza ezt a tartalmat

a tesztelesben ez a csomag nagyon sokat segithet:
- net/http/httptest

ezek a komponensek segithetnek a teszteles soran:
- httptest.NewServer
- httptest.ResponseRecorder
- httptest.NewRequest
- http.MethodGet
- http.MethodPost
- bytes.Buffer

a bytes.Buffert lehet helyetesiteni akar strings.NewReader -el is

|---------------------------|
| GET PATH                           |
| Header1: header1-value    |
| Header2: header2-value   |
| Header2: header2-value   |
|                                              |
| a request body tartalma    |
| akar hany soron keresztul |
| ...                                          |

https://pkg.go.dev/net/http

https://pkg.go.dev/net/http/httptest

https://pkg.go.dev/net/http/httptest#example-ResponseRecorder
-----------------------------------
