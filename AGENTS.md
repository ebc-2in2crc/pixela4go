# pixela4go

Go client library for the [Pixela](https://pixe.la/) API. Package name is `pixela` (import alias typically omitted).

## Architecture

```
client.go        — Client struct and sub-client factory methods
api.go           — HTTP machinery: requestParameter, Result, doRequest*, newHTTPRequest
retry.go         — Exponential-backoff retry (global RetryCount, max 20)
util.go          — Pointer helpers: String/StringValue, Bool/BoolValue
user.go          — User Create/Update/Delete
user_profile.go  — UserProfile Update + URL
graph.go         — Graph CRUD + SVG, Stats, Stopwatch, Add/Subtract, GetPixelDates,
                   GetLatestPixel, GetToday, UpdatePixels, URL, Get (single), Analyze
pixel.go         — Pixel Create/Get/Update/Delete + Increment/Decrement/Add/Subtract
webhook.go       — Webhook Create/GetAll/Delete/Invoke
test_common.go   — Shared test helpers and httpClientMock
e2e_*.go         — E2E tests (skipped unless PIXELA4GO_E2E_TEST_RUN=ON)
```

## Core Patterns

### Dual-method convention
Every public operation has two signatures:
```go
func (x *X) Method(input *XMethodInput) (*Result, error)
func (x *X) MethodWithContext(ctx context.Context, input *XMethodInput) (*Result, error)
```
The no-context variant calls `MethodWithContext(context.Background(), input)`.

### Input structs use pointer fields
All optional and most required input fields are `*string` or `*bool`, never bare values.
This lets callers omit fields (nil = not sent) without separate "zero value" ambiguity.
```go
type GraphCreateInput struct {
    ID    *string `json:"id"`      // required
    Color *string `json:"color"`   // required
    IsSecret *bool `json:"isSecret,omitempty"` // optional
}
```
Use `pixela.String("value")` and `pixela.Bool(true)` to create pointer literals.
Use `json:"-"` for fields that belong in the URL path, not the JSON body (e.g. `GraphID`, `Date`).

### requestParameter flow
Each method builds a `*requestParameter{Method, URL, Header, Body}` via a private
`createXxxRequestParameter` function, then passes it to one of:
- `doRequestAndParseResponse` — for standard JSON `{"message":"…","isSuccess":…}` responses
- `doRequest` — for responses that need custom unmarshalling (GraphDefinitions, Pixels, etc.)
- `mustDoRequest` — for non-JSON responses (SVG) or when HTTP status determines success

### IsSuccess determination
- `doRequestAndParseResponse` (returns `*Result`): `IsSuccess` is set directly from the API's `isSuccess` JSON field.
- Custom struct responses via `doRequest`: success responses from these endpoints typically omit `isSuccess`, so `IsSuccess` is computed explicitly:
  - Most endpoints (GetAll, Stats, GetPixelDates, Get, Pixel.Get, Webhook.GetAll): `IsSuccess = (message == "")` — no error message means success.
  - GetLatestPixel, GetToday, Analyze: `IsSuccess = (statusCode == 200)`.

### Retry on rejection
Pixela returns HTTP 503 + `{"isRejected":true}` when rate-limited.
Set `pixela.RetryCount = N` (global, default 0) before making calls to enable exponential backoff
(`2^i × 100ms`, capped at 20 retries). Context cancellation is respected mid-backoff.

## Testing

### Unit tests
Tests live alongside source files in package `pixela`. Each test creates a `Client` and
sets its `HTTPClient` field to one of the mock helpers:
```go
client := New(userName, token)
client.HTTPClient = newOKMock()           // 200 + {"message":"Success.","isSuccess":true}
client.HTTPClient = newAPIFailedMock()    // 404 + {"message":"failed.","isSuccess":false}
client.HTTPClient = newPageNotFoundMock() // 404 + "404 page not found" (triggers unmarshal error)
```
Each feature file has three test shapes:
1. `TestX_CreateXRequestParameter` — verifies the HTTP method, URL, headers, and JSON body
2. `TestX_Method` — verifies the full call succeeds via `newOKMock()`
3. `TestX_MethodFail` / `TestX_MethodError` — verifies failure paths via `newAPIFailedMock()` / `newPageNotFoundMock()`

### E2E tests
Collected in `e2e_*.go`. All run under a single `TestE2E` gate that skips unless
`PIXELA4GO_E2E_TEST_RUN=ON`. Required env vars:
```
PIXELA4GO_E2E_TEST_RUN=ON
PIXELA4GO_USER_NAME=<username>
PIXELA4GO_USER_FIRST_TOKEN=<token>
PIXELA4GO_USER_SECOND_TOKEN=<token>
PIXELA4GO_THANKS_CODE=<code>
```
E2E tests create a `Client` via `New(name, token)` (which defaults to `&http.Client{}`), hitting the real API.

## Commands

```bash
make test     # go test -v ./...
make fmt      # goimports -w on all *.go files
make lint     # go vet + golint
make deps     # go mod download
make devel-deps  # install goimports, golint, make2help
```

Run only unit tests (no E2E):
```bash
go test -v ./...   # PIXELA4GO_E2E_TEST_RUN unset → E2E skipped automatically
```

## Adding a New API Method

1. Identify which domain file it belongs to (`graph.go`, `pixel.go`, etc.).
2. Define `XxxInput` struct with pointer fields; mark URL-path fields `json:"-"`.
3. Implement `createXxxRequestParameter(input) (*requestParameter, error)`.
4. Implement `Xxx(input)` delegating to `XxxWithContext(context.Background(), input)`.
5. Implement `XxxWithContext(ctx, input)` calling the appropriate `doRequest*` helper.
6. Add unit tests:
   - `TestX_CreateXxxRequestParameter` — check Method, URL, headers, Body bytes
   - `TestX_Xxx` with `newOKMock()`
   - `TestX_XxxFail` with `newAPIFailedMock()`
7. Add an E2E function `testE2EXxx(t)` in the matching `e2e_*_test.go` file and wire it into `TestE2E`.

## Constants Reference

```go
// Graph type
GraphTypeInt / GraphTypeFloat

// Graph color
GraphColorShibafu / Momiji / Sora / Ichou / Ajisai / Kuro

// Graph selfSufficient
GraphSelfSufficientIncrement / Decrement / None

// Graph SVG mode
GraphModeShort / Badge / Line / Simple / SimpleShort

// Graph appearance
GraphAppearanceDark

// Webhook type
WebhookTypeAdd / Increment / Decrement / Stopwatch / Subtract
```

## Branch Naming

```
issue/:id   e.g. issue/42/add-graph-foo
```

## Commit Convention

Follow [Conventional Commits](https://www.conventionalcommits.org/). Format:

```
<type>[(<scope>)]: <description>
```

Scope is optional. When used, it should be the domain being changed: `graph`, `pixel`, `user`, `user-profile`, `webhook`, `api`, `retry`, `util`.

| Type | When to use | Example |
|---|---|---|
| `feat` | New public method or input field | `feat(graph): add Graph.Analyze method` |
| `fix` | Wrong URL, HTTP method, field name, IsSuccess logic | `fix(pixel): use correct endpoint for Decrement` |
| `docs` | README, AGENTS.md, godoc comments | `docs: add commit convention section` |
| `refactor` | Internal restructuring, no behavior change | `refactor(api): replace internal types with anonymous structs` |
| `test` | Unit/E2E test additions or changes only | `test(graph): add TestGraph_Analyze` |
| `chore` | Dependency updates, stdlib migrations, Makefile, go.mod | `chore: replace ioutil with io package` |
| `ci` | `.github/workflows/` changes | `ci: upgrade setup-go to v4` |

For breaking changes (removed method, changed signature), append `!` after the type (`feat!:`) or add a `BREAKING CHANGE:` footer.
