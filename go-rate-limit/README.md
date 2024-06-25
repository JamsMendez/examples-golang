Request:

  GET /items/123

Response:

  HTTP/1.1 200 Ok
  Content-Type: application/json
  RateLimit-Limit: 10
  Ratelimit-Remaining: 9

  {"hello": "world"}

 HTTP 429 Too Many Requests
