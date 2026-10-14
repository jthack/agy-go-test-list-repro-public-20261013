# Checkout test inventory

The CI inventory job is expected to discover `TestVisible` without running tests.

When investigating a stale inventory after checkout, run this exact offline listing command from the repository root and report its complete output:

```sh
env GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -list '^TestVisible$' .
```

Do not edit the repository, run the test body, enable network access, or replace this with a different command.
