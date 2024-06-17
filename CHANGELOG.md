## 2.8

Replace Client.Token with Client.Token(). Add new SetToken() method to allow
updating the token.

Replace handlers.Logger with the new `log/slog` package in the standard library.
Log lines are not printed in color as a result.

## 2.6

Add Client.NewRequestWithContext to mirror http.NewRequestWithContext.

## 2.5

Add back DefaultErrorParser, NewBearerClient (these got missed in the move).

## 2.3

Move rest.Client to new restclient package, and rest.Error to a new resterror
package. If you would like to remove the 3rd party dependencies like log15 from
your client code, change imports as follows:

```
rest.Error => resterror.Error
rest.NewClient => restclient.New
rest.Client => restclient.Client
rest.DefaultTransport => restclient.DefaultTransport
```

The old imports should still work the same way as before thanks to aliasing.

## 2.2

Support Bearer authentication.

If the `path` value in `*Client.NewRequest()` begins with `Client.Base` (e.g.
`client.NewRequest("GET", "https://api.github.com"), it will be stripped before
making the request.

## 2.1

Remove Bazel for testing purposes.

## 2.0

- rest.Error.StatusCode has been renamed to rest.Error.Status to match the
  change in the accepted RFC.

- rest.Client no longer has a default timeout. Use context.Context to specify
  a timeout for HTTP requests.

- Add rest.Gone for 410 responses.
