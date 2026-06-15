<script setup>
const props = defineProps({
  columns: Array,
  rows: Array,
  loading: Boolean,
  total: Number,
  limit: Number,
  offset: Number,
  lookups: { type: Object, default: () => ({}) }, // { categories: {id:name}, ... }
})
const emit = defineEmits(['page', 'edit', 'delete'])

function cellValue(row, col) {
  let v = row[col.key]
  if (col.lookup) return props.lookups[col.lookup]?.[v] ?? ('#' + v)
  if (col.format) return col.format(v)
  return v ?? '—'
}
const page = () => Math.floor(props.offset / props.limit) + 1
const pages = () => Math.max(1, Math.ceil(props.total / props.limit))
</script>

<template>
  <div class="card">
    <div class="table-wrap">
      <table class="data">
        <thead>
          <tr>
            <th v-for="c in columns" :key="c.key" :style="c.width ? { width: c.width } : null">{{ c.label }}</th>
            <th style="width: 1%; text-align: right;">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="loading">
            <td :colspan="columns.length + 1"><div class="row" style="justify-content:center; padding:24px"><span class="spinner" /></div></td>
          </tr>
          <tr v-else-if="!rows.length">
            <td :colspan="columns.length + 1"><div class="empty">No records found</div></td>
          </tr>
          <tr v-for="row in rows" :key="row.id" v-else>
            <td v-for="c in columns" :key="c.key">
              <span v-if="c.badge" class="badge" :class="'role-' + row[c.key]">{{ row[c.key] }}</span>
              <span v-else-if="c.stock" :class="row[c.key] === 0 ? 'stock-out' : row[c.key] < 10 ? 'stock-low' : ''">
                {{ row[c.key] === 0 ? 'Out of stock' : row[c.key] }}
              </span>
              <template v-else>{{ cellValue(row, c) }}</template>
            </td>
            <td style="text-align: right;">
              <div class="row" style="justify-content: flex-end; gap: 4px;">
                <button class="btn btn-ghost btn-sm" @click="emit('edit', row)">Edit</button>
                <button class="btn btn-ghost btn-sm" style="color: var(--danger)" @click="emit('delete', row)">Delete</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="pager">
      <span class="muted">{{ total }} record{{ total === 1 ? '' : 's' }}</span>
      <div class="row">
        <button class="btn btn-sm" :disabled="offset === 0 || loading" @click="emit('page', offset - limit)">‹ Prev</button>
        <span class="muted" style="font-size:12.5px">Page {{ page() }} / {{ pages() }}</span>
        <button class="btn btn-sm" :disabled="offset + limit >= total || loading" @click="emit('page', offset + limit)">Next ›</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.pager { display: flex; align-items: center; justify-content: space-between; padding: 12px 16px; border-top: 1px solid var(--border); }
</style>
