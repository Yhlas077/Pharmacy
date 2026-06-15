<script setup>
import { ref, reactive, onMounted } from 'vue'
import { orders } from '../api/index.js'
import { ORDER_STATUSES } from '../api/constants.js'
import { notify } from '../store.js'
import Modal from '../components/Modal.vue'

const rows = ref([])
const total = ref(0)
const loading = ref(false)
const limit = 8
const offset = ref(0)
const filters = reactive({ status: '' })

async function load() {
  loading.value = true
  try {
    const params = { limit, offset: offset.value }
    if (filters.status) params.status = filters.status
    const res = await orders.list(params)
    rows.value = res.data
    total.value = res.meta.total
  } catch (e) {
    notify(e.error_msg || 'Failed to load orders', 'error')
  } finally {
    loading.value = false
  }
}
function setStatus(v) { filters.status = v; offset.value = 0; load() }
function changePage(o) { offset.value = Math.max(0, o); load() }
onMounted(load)

const money = (v) => '₼' + Number(v).toFixed(2)

// detail drawer
const selected = ref(null)
const saving = ref(false)
const nextStatus = ref('')

function open(order) {
  selected.value = order
  nextStatus.value = ''
}
async function applyStatus() {
  if (!nextStatus.value) return
  saving.value = true
  try {
    const res = await orders.updateStatus(selected.value.id, nextStatus.value)
    notify(`Order #${selected.value.id} → ${nextStatus.value}`)
    selected.value = res.data
    nextStatus.value = ''
    load()
  } catch (e) {
    notify(e.error_msg || 'Update failed', 'error')
  } finally {
    saving.value = false
  }
}
</script>

<template>
  <div>
    <div class="spread" style="margin-bottom: 18px;">
      <div>
        <h2 style="margin: 0 0 2px;">Orders</h2>
        <p class="muted" style="margin: 0; font-size: 13px;">View orders and update fulfilment status</p>
      </div>
    </div>

    <div class="toolbar card">
      <button class="chip" :class="{ active: filters.status === '' }" @click="setStatus('')">All</button>
      <button v-for="s in ORDER_STATUSES" :key="s" class="chip" :class="{ active: filters.status === s }" @click="setStatus(s)">
        {{ s }}
      </button>
    </div>

    <div class="card">
      <div class="table-wrap">
        <table class="data">
          <thead>
            <tr><th>Order</th><th>Items</th><th>Total</th><th>Address</th><th>Status</th><th></th></tr>
          </thead>
          <tbody>
            <tr v-if="loading"><td colspan="6"><div class="row" style="justify-content:center;padding:24px"><span class="spinner" /></div></td></tr>
            <tr v-else-if="!rows.length"><td colspan="6"><div class="empty">No orders</div></td></tr>
            <tr v-for="o in rows" :key="o.id" v-else>
              <td><strong>#{{ o.id }}</strong></td>
              <td>{{ o.items.reduce((n, i) => n + i.quantity, 0) }} item(s)</td>
              <td>{{ money(o.total) }}</td>
              <td class="muted">{{ o.address }}</td>
              <td><span class="badge" :class="'status-' + o.status">{{ o.status }}</span></td>
              <td style="text-align:right"><button class="btn btn-ghost btn-sm" @click="open(o)">View</button></td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="pager">
        <span class="muted">{{ total }} order{{ total === 1 ? '' : 's' }}</span>
        <div class="row">
          <button class="btn btn-sm" :disabled="offset === 0 || loading" @click="changePage(offset - limit)">‹ Prev</button>
          <button class="btn btn-sm" :disabled="offset + limit >= total || loading" @click="changePage(offset + limit)">Next ›</button>
        </div>
      </div>
    </div>

    <!-- Order detail -->
    <Modal v-if="selected" :title="`Order #${selected.id}`" @close="selected = null">
      <div class="spread" style="margin-bottom:14px">
        <span class="badge" :class="'status-' + selected.status">{{ selected.status }}</span>
        <strong>{{ money(selected.total) }}</strong>
      </div>

      <table class="data" style="border:1px solid var(--border); border-radius:8px; overflow:hidden; margin-bottom:16px">
        <thead><tr><th>Item</th><th>Qty</th><th>Price</th><th>Subtotal</th></tr></thead>
        <tbody>
          <tr v-for="(i, idx) in selected.items" :key="idx">
            <td>{{ i.name }}</td><td>{{ i.quantity }}</td><td>{{ money(i.price) }}</td><td>{{ money(i.price * i.quantity) }}</td>
          </tr>
        </tbody>
      </table>

      <div class="field"><label>Delivery address</label><div>{{ selected.address }}</div></div>
      <div class="field" v-if="selected.notes"><label>Notes</label><div class="muted">{{ selected.notes }}</div></div>

      <div class="field" v-if="orders.allowedNext(selected.status).length">
        <label>Update status</label>
        <div class="row">
          <select v-model="nextStatus" class="select" style="max-width:220px">
            <option value="" disabled>Move to…</option>
            <option v-for="s in orders.allowedNext(selected.status)" :key="s" :value="s">{{ s }}</option>
          </select>
          <button class="btn btn-primary" :disabled="!nextStatus || saving" @click="applyStatus">Apply</button>
        </div>
      </div>
      <p v-else class="muted" style="font-size:12.5px">No further status changes available.</p>
    </Modal>
  </div>
</template>

<style scoped>
.toolbar { display: flex; gap: 8px; padding: 12px 14px; margin-bottom: 16px; flex-wrap: wrap; }
.chip { padding: 6px 14px; border-radius: 999px; border: 1px solid var(--border); background: var(--surface); cursor: pointer; font-size: 12.5px; font-weight: 600; text-transform: capitalize; color: var(--text-soft); }
.chip:hover { background: var(--surface-2); }
.chip.active { background: var(--primary); color: #fff; border-color: var(--primary); }
</style>
