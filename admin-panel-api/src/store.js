import { reactive, readonly } from 'vue'
import { auth } from './api/index.js'

// --- Auth state (persisted to localStorage) ----------------------------------
const saved = JSON.parse(localStorage.getItem('pharmacy_auth') || 'null')
const authState = reactive({
  token: saved?.token || null,
  user: saved?.user || null,
})

export const useAuth = () => ({
  state: readonly(authState),
  get isAuthed() { return !!authState.token },
  async login(email, password) {
    const res = await auth.login({ email, password })
    authState.token = res.data.token
    authState.user = res.data.user
    localStorage.setItem('pharmacy_auth', JSON.stringify({ token: authState.token, user: authState.user }))
    return res
  },
  async logout() {
    await auth.logout()
    authState.token = null
    authState.user = null
    localStorage.removeItem('pharmacy_auth')
  },
})

// --- Toast notifications -----------------------------------------------------
export const toasts = reactive([])
let toastId = 0
export function notify(message, type = 'success') {
  const id = ++toastId
  toasts.push({ id, message, type })
  setTimeout(() => {
    const i = toasts.findIndex((t) => t.id === id)
    if (i !== -1) toasts.splice(i, 1)
  }, 3500)
}
