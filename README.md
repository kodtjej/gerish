
# Gerish - for all your API-testing error, timeout, slowish needs

Available commands:
```
slow, s starts a slow http server - default sleeps for 11 seconds before returning a response.

faulty, f starts a faulty http server that only responds with error codes.

unstable, u starts a http server that is unstable and responds with an error code every other request (by default).
```