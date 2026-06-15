# Pharmacy Admin Panel — Live API

Same UI as [`../admin-panel`](../admin-panel), but it talks to the **real
backend over `fetch`** instead of mock data. All requests/responses follow
[`../docs/API.md`](../docs/API.md).

## Configure

Set the API base URL in `.env` (copy from `.env.example`):

```
VITE_API_BASE=http://localhost:8080/api
```

## Run

```bash
npm install
npm run dev      # http://localhost:5173
npm run build    # production build → dist/
```

You log in with **real credentials** via `POST /auth/login`; the returned token
is stored in `localStorage` and sent as `Authorization: Bearer <token>` on every
request.

## Where the API lives

Everything network-related is in two files:

| File                | Responsibility                                             |
|---------------------|-----------------------------------------------------------|
| `src/api/http.js`   | `fetch` wrapper: base URL, Bearer header, envelope parsing, errors |
| `src/api/index.js`  | Endpoint paths per resource (the single source of route truth)     |

If your live backend uses **different routes** than `API.md` (e.g. the current
Go code uses `?token=` and `/admin/category`), change them **only** in those two
files — components and views stay untouched.

### Expected contract (see API.md)

- `POST /auth/login` → `{ data: { token, expires_at, user } }`
- `POST /auth/logout`
- `GET|POST /admin/users`, `GET|PUT|DELETE /admin/users/:id` (same shape for
  `categories`, `medicines`, `pharmacies`)
- `GET /admin/orders`, `GET /admin/orders/:id`, `PATCH /admin/orders/:id/status`
- List responses include `meta: { total, limit, offset }`
- Errors: `{ success: false, error_msg, error_code }`

> ⚠️ The backend must implement this contract for the panel to work. The current
> Go server does **not** yet — see [`../docs/TODO.md`](../docs/TODO.md). Until
> then, use the mock panel in `../admin-panel`, or point `VITE_API_BASE` at a
> server that matches.

## Structure

Identical to the mock panel except the `src/api/` folder:

```
src/api/
  constants.js   enums + order status-transition map
  http.js        fetch wrapper (Bearer auth, envelope handling)
  index.js       per-resource endpoints
```
