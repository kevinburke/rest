# Changelog

This file summarizes the notable changes in each release.

## 2.14.0 - 2026-06-05

- Moved the module path to `github.com/kevinburke/rest/v2`. The module had
  been tagged `v2.x` since `2.0`, but `go.mod` lacked the `/v2` major-version
  suffix that Go modules require, so the tooling ignored every `v2.x` tag.
  Consumers must now import `github.com/kevinburke/rest/v2` (and
  `.../v2/restclient`, `.../v2/resterror`).
- Added an Install section to the README documenting the new import path.
- Made the release target's version bump configurable via `make release
  version=...`.
- Bumped `actions/checkout` to 6.0.2.

## 2.13.0 - 2026-05-21

- Trimmed development user-agent version strings so they stay stable and
  readable.
- Removed use of the deprecated `net.Dialer.DualStack` field from
  `restclient`.
- Documented how to wrap a `*Transport` to inspect response headers (e.g.
  rate limits) in the README.
- Refreshed CI coverage for newer Go versions and dependency updates.
- Added daily Dependabot updates for GitHub Actions, with a 7-day cooldown.

## 2.12.0 - 2025-07-18

- Updated the `log15` dependency to v3.

## 2.11.0 - 2025-07-16

- Added a tag prefix in the release `Makefile` flow.

## 2.10.0 - 2025-07-16

- Updated dependencies and refreshed tested Go versions.
- Made the generated user-agent version string more Go-version compatible.

## 2.9 - 2024-06-16

- Replaced the writable `Client.Token` field with `Client.Token()` and
  `SetToken()` so tokens can be updated safely at runtime.
- Expanded CI to cover Go 1.22.

## 2.8 - 2023-11-07

- Switched logging from `log15` handlers to the standard library `log/slog`
  package.
- Log lines are no longer printed in color.
- Updated module dependencies to support the logging change.

## 2.7 - 2023-03-05

- Threaded `context.Context` through `restclient` requests.
- Tightened `restclient` behavior by panicking on a nil `*Client`.
- Reduced string parsing overhead in `restclient`.
- Updated CI to test on Go 1.19.

## 2.6 - 2021-04-25

- Added `Client.NewRequestWithContext` to mirror
  `http.NewRequestWithContext`.
- Updated CI for Go 1.16.

## 2.5 - 2021-01-06

- Restored the `DefaultErrorParser` alias that was missed in the package split.

## 2.4 - 2021-01-05

- Pointed README examples at the new package locations.
- Restored the `NewBearerClient` alias and other compatibility aliases in
  `rest`.

## 2.3 - 2021-01-05

- Split the code into new `restclient` and `resterror` packages while keeping
  the old imports working through aliases.
- Added GitHub Actions and updated the release target.

If you would like to remove third-party dependencies like `log15` from your
client code, change imports as follows:

```
rest.Error => resterror.Error
rest.NewClient => restclient.New
rest.Client => restclient.Client
rest.DefaultTransport => restclient.DefaultTransport
```

## 2.2 - 2020-04-29

- Added Bearer authentication support.
- Stripped `Client.Base` from request paths when callers accidentally passed a
  full URL.
- Switched the lint target from `megacheck` to `staticcheck`.

## 2.1 - 2018-04-23

- Removed Bazel from the test and release workflow.
- Added the first project changelog and simplified release tooling.

## 2.0 - 2018-03-19

- Renamed `rest.Error.StatusCode` to `rest.Error.Status`.
- Removed the default `rest.Client` timeout so callers can control timeouts via
  `context.Context`.
- Added `rest.Gone` for HTTP 410 responses.
- Dropped Bazel from the codebase.

## 1.3 - 2018-03-06

- Added the `410 Gone` HTTP status helper.

## 1.2 - 2018-02-22

- Updated the `Error` struct.
- Removed the default timeout behavior that was later carried into 2.0.
- Fixed response body closing behavior after errors.
- Updated the Shyp import path.

## 1.1 - 2017-10-24

- Moved tests and developer workflows over to Bazel.
- Improved test logging, cache experiments, and build profiling.
- Added more examples and documentation for `NewClient`.

## 1.0 - 2017-03-18

- Added transport examples to the README and test suite.
- Added Go 1.8 to CI.
- Defaulted the debug writer instead of panicking.

## 0.18 - 2017-02-15

- Stopped panicking when `http.Client` is nil.

## 0.17 - 2017-02-15

- Added `DEBUG_HTTP_TRAFFIC` support as a `RoundTripper`.
- Made `RegisterHandler` more generic.
- Added `staticcheck`, removed Go 1.5 from CI, and refreshed project metadata.

## 0.16 - 2016-11-07

- Reduced per-request user-agent allocations.
- Improved release tooling and refreshed the license metadata.

## 0.15 - 2016-10-20

- Added support for custom `ErrorParser` implementations.
- Started running the race detector before releases.

## 0.14 - 2016-10-20

- Corrected `401 Unauthorized` status handling.

## 0.13 - 2016-10-20

- Allowed `RegisterHandler(code, nil)` to remove a handler.

## 0.12 - 2016-10-19

- Added support for registering custom error handlers.
- Expanded CI beyond `go vet`, including a Go 1.7 context-specific test.

## 0.11 - 2016-10-18

- Added a newline before debug messages for cleaner logging.

## 0.10 - 2016-10-17

- Populated the forbidden error ID when it was missing.

## 0.9 - 2016-10-12

- Added Go runtime version and architecture information to the user-agent.

## 0.8 - 2016-10-11

- Added an `Unauthorized` handler.

## 0.7 - 2016-10-10

- Added a `204 No Content` handler.

## 0.6 - 2016-09-22

- Sent additional request headers by default.

## 0.5 - 2016-09-19

- Expanded the package documentation.

## 0.4 - 2016-09-19

- Added the README and package-level documentation.

## 0.3 - 2016-09-03

- Added Travis CI, formatting and vet checks, and a release target.

## 0.2 - 2016-07-15

- Added a `404 NotFound` handler.
- Returned `nil` for `204 No Content` responses.

## 0.1 - 2016-07-15

- Initial release of the REST client.

Thanks to `burkebot` for the recent maintenance updates.
