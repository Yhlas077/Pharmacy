<script setup>
import { ref, reactive, computed, watch, onMounted } from 'vue'
import { resources } from '../resources.js'
import { lookups as lookupApi } from '../api/index.js'
import { notify } from '../store.js'
import DataTable from '../components/DataTable.vue'
import FormModal from '../components/FormModal.vue'
import ConfirmDialog from '../components/ConfirmDialog.vue'

const props = defineProps({ resource: String })

const cfg = computed(() => resources[props.resource])

const rows = ref([])
const total = ref(0)
const loading = ref(false)
const limit = 8
const offset = ref(0)
const filters = reactive({})

// Lookup tables for select fields & table display (category_id → name, etc.)
const lookupOptions = reactive({}) // { categories: [...], pharmacies: [...] }
const lookupMaps = reactive({})     // { categories: {1:'Pain Relief'}, ... }

function neededLookups() {
  const set = new Set()
  cfg.value.columns.forEach((c) => c.lookup && set.add(c.lookup))
  cfg.value.fields.forEach((f) => f.lookup && set.add(f.lookup))
  cfg.value.filters.forEach((f) => f.lookup && set.add(f.lookup))
  return [...set]
}

async function loadLookups() {
  for (const key of neededLookups()) {
    const data = await lookupApi[key]()
    lookupOptions[key] = data
    lookupMaps[key] = Object.fromEntries(data.map((d) => [d.id, d.name]))
  }
}

async function load() {
  loading.value = true
  try {
    const params = { limit, offset: offset.value }
    for (const f of cfg.value.filters) {
      const v = filters[f.key]
      if (v !== '' && v != null) params[f.key] = v
    }
    const res = await cfg.value.api.list(params)
    rows.value = res.data
    total.value = res.meta.total
  } catch (e) {
    notify(e.error_msg || 'Failed to load', 'error')
  } finally {
    loading.value = false
  }
}

function resetAndLoad() { offset.value = 0; load() }
function changePage(o) { offset.value = Math.max(0, o); load() }

// debounce search-type filters
let t
watch(filters, () => { clearTimeout(t); t = setTimeout(resetAndLoad, 300) })

// re-init when navigating between resources
watch(() => props.resource, init)

function init() {
  offset.value = 0
  Object.keys(filters).forEach((k) => delete filters[k])
  cfg.value.filters.forEach((f) => (filters[f.key] = ''))
  loadLookups().then(load)
}
onMounted(init)

// ---- modals ----
const formOpen = ref(false)
const editing = ref(null)
const saving = ref(false)

function openCreate() { editing.value = null; formOpen.value = true }
function openEdit(row) { editing.value = row; formOpen.value = true }

async function save(payload) {
  saving.value = true
  try {
    if (editing.value) {
      await cfg.value.api.update(editing.value.id, payload)
      notify(`${cfg.value.singular} updated`)
    } else {
      await cfg.value.api.create(payload)
      notify(`${cfg.value.singular} created`)
    }
    formOpen.value = false
    load()
  } catch (e) {
    notify(e.error_msg || 'Save failed', 'error')
  } finally {
    saving.value = false
  }
}

const confirmRow = ref(null)
async function doDelete() {
  saving.value = true
  try {
    await cfg.value.api.remove(confirmRow.value.id)
    notify(`${cfg.value.singular} deleted`)
    confirmRow.value = null
    if (rows.value.length === 1 && offset.value > 0) offset.value -= limit
    load()
  } catch (e) {
    notify(e.error_msg || 'Delete failed', 'error')
  } finally {
    saving.value = false
  }
}
</script>

<template>
  <div>
    <div class="spread" style="margin-bottom: 18px;">
      <div>
        <h2 style="margin: 0 0 2px;">{{ cfg.title }}</h2>
        <p class="muted" style="margin: 0; font-size: 13px;">Manage {{ cfg.title.toLowerCase() }}</p>
      </div>
      <button class="btn btn-primary" @click="openCreate">+ New {{ cfg.singular }}</button>
    </div>

    <!-- Filters toolbar -->
    <div class="toolbar card" v-if="cfg.filters.length">
      <template v-for="f in cfg.filters" :key="f.key">
        <input v-if="f.type === 'search'" v-model="filters[f.key]" class="input" :placeholder="f.placeholder" style="max-width: 280px;" />
        <select v-else-if="f.type === 'select'" v-model="filters[f.key]" class="select" style="max-width: 200px;">
          <option value="">{{ f.allLabel || 'All' }}</option>
          <option v-for="o in (f.lookup ? lookupOptions[f.lookup] : f.options)" :key="f.lookup ? o.id : o" :value="f.lookup ? o.id : o">
            {{ f.lookup ? o.name : o }}
          </option>
        </select>
      </template>
    </div>

    <DataTable
      :columns="cfg.columns" :rows="rows" :loading="loading"
      :total="total" :limit="limit" :offset="offset" :lookups="lookupMaps"
      @page="changePage" @edit="openEdit" @delete="(r) => (confirmRow = r)"
    />

    <FormModal
      v-if="formOpen"
      :title="(editing ? 'Edit ' : 'New ') + cfg.singular"
      :fields="cfg.fields" :model="editing" :lookup-options="lookupOptions" :saving="saving"
      @close="formOpen = false" @save="save"
    />

    <ConfirmDialog
      v-if="confirmRow"
      :title="'Delete ' + cfg.singular"
      :message="`Are you sure you want to delete this ${cfg.singular.toLowerCase()}? This cannot be undone.`"
      :saving="saving"
      @close="confirmRow = null" @confirm="doDelete"
    />
  </div>
</template>

<style scoped>
.toolbar { display: flex; gap: 12px; padding: 14px 16px; margin-bottom: 16px; flex-wrap: wrap; }
</style>
