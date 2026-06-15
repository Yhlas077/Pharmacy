<script setup>
import Modal from './Modal.vue'
defineProps({ title: String, message: String, saving: Boolean })
const emit = defineEmits(['close', 'confirm'])
</script>

<template>
  <Modal :title="title || 'Confirm'" @close="emit('close')">
    <p style="margin: 0; color: var(--text-soft);">{{ message }}</p>
    <template #footer>
      <button class="btn" @click="emit('close')">Cancel</button>
      <button class="btn btn-danger" :disabled="saving" @click="emit('confirm')">Delete</button>
    </template>
  </Modal>
</template>
