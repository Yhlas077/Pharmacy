<script setup>
import { useRouter } from 'vue-router'
import { useAuth, notify } from '../store.js'

const router = useRouter()
const auth = useAuth()

const nav = [
  { to: '/dashboard', label: 'Dashboard', icon: '▥' },
  { to: '/users', label: 'Users', icon: '☺' },
  { to: '/medicines', label: 'Medicines', icon: '✚' },
  { to: '/categories', label: 'Categories', icon: '❏' },
  { to: '/pharmacies', label: 'Pharmacies', icon: '⌂' },
  { to: '/orders', label: 'Orders', icon: '🛒' },
]

async function logout() {
  await auth.logout()
  notify('Logged out')
  router.push('/login')
}
</script>

<template>
  <div class="shell">
    <aside class="sidebar">
      <div class="brand">
        <span class="brand-mark">℞</span>
        <span class="brand-name">Pharmacy<small>admin</small></span>
      </div>
      <nav>
        <router-link v-for="n in nav" :key="n.to" :to="n.to" class="nav-link" active-class="active">
          <span class="nav-ico">{{ n.icon }}</span>{{ n.label }}
        </router-link>
      </nav>
      <div class="sidebar-foot muted">Live API · API.md contract</div>
    </aside>

    <div class="main">
      <header class="topbar">
        <div class="muted">Pharmacy Admin Panel</div>
        <div class="row">
          <div class="user-chip">
            <strong>{{ auth.state.user?.name }}</strong>
            <span class="badge" :class="'role-' + auth.state.user?.role">{{ auth.state.user?.role }}</span>
          </div>
          <button class="btn btn-ghost btn-sm" @click="logout">Logout</button>
        </div>
      </header>
      <main class="content">
        <slot />
      </main>
    </div>
  </div>
</template>

<style scoped>
.shell { display: flex; min-height: 100vh; }
.sidebar {
  width: var(--sidebar-w); flex-shrink: 0; background: #0f172a; color: #cbd5e1;
  display: flex; flex-direction: column; padding: 18px 14px; position: sticky; top: 0; height: 100vh;
}
.brand { display: flex; align-items: center; gap: 10px; padding: 6px 8px 18px; }
.brand-mark { width: 34px; height: 34px; display: grid; place-items: center; background: var(--primary); color: #fff; border-radius: 9px; font-size: 18px; font-weight: 700; }
.brand-name { font-size: 16px; font-weight: 700; color: #fff; display: flex; flex-direction: column; line-height: 1.1; }
.brand-name small { font-size: 10px; font-weight: 500; color: #64748b; text-transform: uppercase; letter-spacing: .12em; }
nav { display: flex; flex-direction: column; gap: 3px; margin-top: 8px; }
.nav-link { display: flex; align-items: center; gap: 11px; padding: 10px 12px; border-radius: 8px; font-weight: 500; font-size: 13.5px; transition: all .15s; }
.nav-link:hover { background: #1e293b; color: #fff; }
.nav-link.active { background: var(--primary); color: #fff; }
.nav-ico { width: 18px; text-align: center; opacity: .9; }
.sidebar-foot { margin-top: auto; font-size: 11px; padding: 10px 8px 2px; color: #475569; }

.main { flex: 1; min-width: 0; display: flex; flex-direction: column; }
.topbar { height: 60px; background: var(--surface); border-bottom: 1px solid var(--border); display: flex; align-items: center; justify-content: space-between; padding: 0 24px; position: sticky; top: 0; z-index: 5; }
.user-chip { display: flex; align-items: center; gap: 8px; font-size: 13px; }
.content { padding: 24px; flex: 1; }
</style>
