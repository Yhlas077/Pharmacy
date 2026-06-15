<script setup>
import { ref, onMounted } from 'vue'
import { users, medicines, categories, pharmacies, orders } from '../api/index.js'
import { useAuth } from '../store.js'

const auth = useAuth()
const stats = ref([])
const recentOrders = ref([])
const lowStock = ref([])
const loading = ref(true)

const money = (v) => '₼' + Number(v).toFixed(2)

onMounted(async () => {
  const [u, m, c, p, o] = await Promise.all([
    users.list({ limit: 1 }), medicines.list({ limit: 1000 }),
    categories.list({ limit: 1 }), pharmacies.list({ limit: 1 }), orders.list({ limit: 1000 }),
  ])
  const revenue = o.data.filter((x) => x.status !== 'cancelled').reduce((s, x) => s + x.total, 0)
  stats.value = [
    { label: 'Users', value: u.meta.total, icon: '☺', color: '#2563eb' },
    { label: 'Medicines', value: m.meta.total, icon: '✚', color: '#16a34a' },
    { label: 'Categories', value: c.meta.total, icon: '❏', color: '#d97706' },
    { label: 'Pharmacies', value: p.meta.total, icon: '⌂', color: '#7c3aed' },
    { label: 'Orders', value: o.meta.total, icon: '🛒', color: '#db2777' },
    { label: 'Revenue', value: money(revenue), icon: '₼', color: '#0e7490' },
  ]
  recentOrders.value = o.data.slice(-5).reverse()
  lowStock.value = m.data.filter((x) => x.stock < 10).sort((a, b) => a.stock - b.stock)
  loading.value = false
})
</script>

<template>
  <div>
    <h2 style="margin: 0 0 4px;">Welcome, {{ auth.state.user?.name }} 👋</h2>
    <p class="muted" style="margin: 0 0 22px;">Here's what's happening in your pharmacy.</p>

    <div v-if="loading" class="row" style="justify-content:center;padding:60px"><span class="spinner" /></div>

    <template v-else>
      <div class="stats">
        <div v-for="s in stats" :key="s.label" class="card stat">
          <div class="stat-ico" :style="{ background: s.color + '18', color: s.color }">{{ s.icon }}</div>
          <div>
            <div class="stat-val">{{ s.value }}</div>
            <div class="muted stat-label">{{ s.label }}</div>
          </div>
        </div>
      </div>

      <div class="grid2">
        <div class="card panel">
          <h3>Recent orders</h3>
          <table class="data">
            <thead><tr><th>Order</th><th>Total</th><th>Status</th></tr></thead>
            <tbody>
              <tr v-for="o in recentOrders" :key="o.id">
                <td><strong>#{{ o.id }}</strong></td>
                <td>{{ money(o.total) }}</td>
                <td><span class="badge" :class="'status-' + o.status">{{ o.status }}</span></td>
              </tr>
            </tbody>
          </table>
          <router-link to="/orders" class="see-all">View all orders →</router-link>
        </div>

        <div class="card panel">
          <h3>Low / out of stock</h3>
          <table class="data" v-if="lowStock.length">
            <thead><tr><th>Medicine</th><th>Stock</th></tr></thead>
            <tbody>
              <tr v-for="m in lowStock" :key="m.id">
                <td>{{ m.name }}</td>
                <td><span :class="m.stock === 0 ? 'stock-out' : 'stock-low'">{{ m.stock === 0 ? 'Out of stock' : m.stock }}</span></td>
              </tr>
            </tbody>
          </table>
          <div v-else class="empty">All medicines well stocked ✓</div>
          <router-link to="/medicines" class="see-all">Manage medicines →</router-link>
        </div>
      </div>
    </template>
  </div>
</template>

<style scoped>
.stats { display: grid; grid-template-columns: repeat(auto-fit, minmax(170px, 1fr)); gap: 16px; margin-bottom: 22px; }
.stat { padding: 18px; display: flex; align-items: center; gap: 14px; }
.stat-ico { width: 46px; height: 46px; border-radius: 11px; display: grid; place-items: center; font-size: 20px; flex-shrink: 0; }
.stat-val { font-size: 22px; font-weight: 700; line-height: 1.1; }
.stat-label { font-size: 12.5px; }
.grid2 { display: grid; grid-template-columns: 1fr 1fr; gap: 16px; }
.panel { padding: 18px; }
.panel h3 { margin: 0 0 12px; font-size: 15px; }
.panel table.data th { background: transparent; }
.panel table.data th, .panel table.data td { padding: 9px 8px; }
.see-all { display: inline-block; margin-top: 12px; color: var(--primary); font-weight: 600; font-size: 13px; }
@media (max-width: 880px) { .grid2 { grid-template-columns: 1fr; } }
</style>
