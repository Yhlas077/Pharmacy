// Real API service layer (fetch). Same interface the components already use,
// so swapping from the mock layer required no component changes.
//
// All paths follow docs/API.md. If your live backend uses different routes,
// adjust them here only — this is the single source of endpoint truth.

import { http } from './http.js'
import { ORDER_TRANSITIONS } from './constants.js'

// ---- Auth -------------------------------------------------------------------
export const auth = {
  login({ email, password }) {
    return http.post('/auth/login', { email, password })
  },
  logout() {
    return http.post('/auth/logout')
  },
}

// ---- Generic CRUD against a base path --------------------------------------
function crud(base) {
  return {
    list(params = {}) { return http.get(base, params) },
    get(id) { return http.get(`${base}/${id}`) },
    create(payload) { return http.post(base, payload) },
    update(id, payload) { return http.put(`${base}/${id}`, payload) },
    remove(id) { return http.del(`${base}/${id}`) },
  }
}

export const users = crud('/admin/users')
export const categories = crud('/admin/categories')
export const pharmacies = crud('/admin/pharmacies')
export const medicines = crud('/admin/medicines')

// ---- Orders (read + status transitions only) --------------------------------
export const orders = {
  list(params = {}) { return http.get('/admin/orders', params) },
  get(id) { return http.get(`/admin/orders/${id}`) },
  updateStatus(id, status) { return http.patch(`/admin/orders/${id}/status`, { status }) },
  allowedNext: (status) => ORDER_TRANSITIONS[status] ?? [],
}

// ---- Lookups for select inputs / dashboard ----------------------------------
export const lookups = {
  async categories() { return (await categories.list({ limit: 1000 })).data },
  async pharmacies() { return (await pharmacies.list({ limit: 1000 })).data },
}
