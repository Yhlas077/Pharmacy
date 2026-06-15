<script setup>
import { ref, reactive, watch } from 'vue'
import Modal from './Modal.vue'

const props = defineProps({
  title: String,
  fields: Array,
  model: Object,          // existing row when editing, null when creating
  lookupOptions: Object,  // { categories: [{id,name}], pharmacies: [...] }
  saving: Boolean,
})
const emit = defineEmits(['close', 'save'])

const isEdit = !!props.model
const form = reactive({})
const errors = ref({})

// initialise form values
for (const f of props.fields) {
  form[f.key] = props.model?.[f.key] ?? f.default ?? ''
}

const visibleFields = props.fields.filter((f) => !(f.createOnly && isEdit))

function optionsFor(f) {
  if (f.lookup) return (props.lookupOptions?.[f.lookup] ?? []).map((o) => ({ value: o.id, label: o.name }))
  return (f.options ?? []).map((o) => ({ value: o, label: o }))
}

function validate() {
  const e = {}
  for (const f of visibleFields) {
    if (f.required && (form[f.key] === '' || form[f.key] == null)) e[f.key] = `${f.label} is required`
  }
  errors.value = e
  return Object.keys(e).length === 0
}

function submit() {
  if (!validate()) return
  const payload = {}
  for (const f of visibleFields) {
    let v = form[f.key]
    if (f.type === 'number' || f.lookup) v = v === '' ? null : Number(v)
    payload[f.key] = v
  }
  emit('save', payload)
}
</script>

<template>
  <Modal :title="title" @close="emit('close')">
    <form @submit.prevent="submit">
      <div v-for="f in visibleFields" :key="f.key" class="field">
        <label>{{ f.label }} <span v-if="f.required" class="req">*</span></label>

        <textarea v-if="f.type === 'textarea'" v-model="form[f.key]" class="input" :placeholder="f.placeholder" />

        <select v-else-if="f.type === 'select'" v-model="form[f.key]" class="select">
          <option value="" disabled>Select…</option>
          <option v-for="o in optionsFor(f)" :key="o.value" :value="o.value">{{ o.label }}</option>
        </select>

        <input v-else v-model="form[f.key]" class="input" :type="f.type || 'text'" :placeholder="f.placeholder" />

        <div v-if="errors[f.key]" class="field-error">{{ errors[f.key] }}</div>
      </div>
    </form>

    <template #footer>
      <button class="btn" @click="emit('close')">Cancel</button>
      <button class="btn btn-primary" :disabled="saving" @click="submit">
        <span v-if="saving" class="spinner" style="width:14px;height:14px;border-color:rgba(255,255,255,.4);border-top-color:#fff" />
        {{ isEdit ? 'Save changes' : 'Create' }}
      </button>
    </template>
  </Modal>
</template>
