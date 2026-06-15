// Thin fetch wrapper. Speaks the docs/API.md envelope:
//   success → { success: true, data, meta }
//   error   → { success: false, error_msg, error_code }
//
// On any non-2xx response or `success: false`, it throws that error object
// (augmented with `status`) so callers can `catch (e) { e.error_msg }`.

const BASE = (import.meta.env.VITE_API_BASE || 'http://localhost:8080/api').replace(/\/$/, '')

function token() {
  try {
    return JSON.parse(localStorage.getItem('pharmacy_auth') || 'null')?.token || null
  } catch {
    return null
  }
}

function buildQuery(params = {}) {
  const q = new URLSearchParams()
  for (const [k, v] of Object.entries(params)) {
    if (v === '' || v == null) continue
    q.append(k, v)
  }
  const s = q.toString()
  return s ? `?${s}` : ''
}

export async function request(path, { method = 'GET', params, body } = {}) {
  const headers = { Accept: 'application/json' }
  const t = token()
  if (t) headers.Authorization = `Bearer ${t}`
  if (body !== undefined) headers['Content-Type'] = 'application/json'

  let res
  try {
    res = await fetch(BASE + path + buildQuery(params), {
      method,
      headers,
      body: body !== undefined ? JSON.stringify(body) : undefined,
    })
  } catch {
    throw { success: false, error_msg: 'Network error — is the API running?', error_code: 'network', status: 0 }
  }

  // 204 No Content (e.g. DELETE) — nothing to parse
  if (res.status === 204) return { success: true, data: null }

  let payload = null
  try {
    payload = await res.json()
  } catch {
    payload = null
  }

  if (!res.ok || payload?.success === false) {
    throw {
      success: false,
      error_msg: payload?.error_msg || `Request failed (${res.status})`,
      error_code: payload?.error_code || 'internal',
      status: res.status,
    }
  }
  return payload ?? { success: true, data: null }
}

export const http = {
  get: (path, params) => request(path, { method: 'GET', params }),
  post: (path, body) => request(path, { method: 'POST', body }),
  put: (path, body) => request(path, { method: 'PUT', body }),
  patch: (path, body) => request(path, { method: 'PATCH', body }),
  del: (path) => request(path, { method: 'DELETE' }),
}
