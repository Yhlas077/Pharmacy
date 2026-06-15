<script setup>
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuth, notify } from '../store.js'

const router = useRouter()
const route = useRoute()
const auth = useAuth()

const email = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)

async function submit() {
  error.value = ''
  loading.value = true
  try {
    await auth.login(email.value, password.value)
    notify('Welcome back!')
    router.push(route.query.redirect || '/dashboard')
  } catch (e) {
    error.value = e.error_msg || 'Login failed'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="login-wrap">
    <div class="login-card card">
      <div class="login-brand">
        <span class="brand-mark">℞</span>
        <div>
          <h1>Pharmacy Admin</h1>
          <p class="muted">Sign in to manage the store</p>
        </div>
      </div>

      <form @submit.prevent="submit">
        <div class="field">
          <label>Email</label>
          <input v-model="email" class="input" type="email" placeholder="you@pharmacy.tm" />
        </div>
        <div class="field">
          <label>Password</label>
          <input v-model="password" class="input" type="password" placeholder="••••••" />
        </div>
        <div v-if="error" class="field-error" style="margin-bottom:12px">{{ error }}</div>
        <button class="btn btn-primary" style="width:100%; justify-content:center" :disabled="loading">
          <span v-if="loading" class="spinner" style="width:14px;height:14px;border-color:rgba(255,255,255,.4);border-top-color:#fff" />
          Sign in
        </button>
      </form>

      <p class="hint muted">Sign in with your admin credentials.</p>
    </div>
  </div>
</template>

<style scoped>
.login-wrap { min-height: 100vh; display: grid; place-items: center; background: linear-gradient(135deg, #1e3a8a 0%, #2563eb 100%); padding: 20px; }
.login-card { width: 100%; max-width: 380px; padding: 32px; }
.login-brand { display: flex; align-items: center; gap: 14px; margin-bottom: 26px; }
.brand-mark { width: 46px; height: 46px; display: grid; place-items: center; background: var(--primary); color: #fff; border-radius: 12px; font-size: 24px; font-weight: 700; }
.login-brand h1 { margin: 0; font-size: 19px; }
.login-brand p { margin: 2px 0 0; font-size: 13px; }
.hint { font-size: 12px; text-align: center; margin: 20px 0 0; }
.hint code { background: var(--surface-2); padding: 1px 5px; border-radius: 4px; }
</style>
