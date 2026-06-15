<script setup>
defineProps({ title: String })
const emit = defineEmits(['close'])
</script>

<template>
  <div class="overlay" @click.self="emit('close')">
    <div class="modal card">
      <header class="modal-head">
        <h3>{{ title }}</h3>
        <button class="btn btn-ghost btn-sm" @click="emit('close')">✕</button>
      </header>
      <div class="modal-body">
        <slot />
      </div>
      <footer v-if="$slots.footer" class="modal-foot">
        <slot name="footer" />
      </footer>
    </div>
  </div>
</template>

<style scoped>
.overlay { position: fixed; inset: 0; background: rgba(15, 23, 42, .45); display: grid; place-items: center; z-index: 100; padding: 20px; }
.modal { width: 100%; max-width: 480px; max-height: 90vh; display: flex; flex-direction: column; animation: pop .18s ease; }
@keyframes pop { from { opacity: 0; transform: translateY(8px) scale(.98); } }
.modal-head { display: flex; align-items: center; justify-content: space-between; padding: 16px 20px; border-bottom: 1px solid var(--border); }
.modal-head h3 { margin: 0; font-size: 16px; }
.modal-body { padding: 20px; overflow-y: auto; }
.modal-foot { padding: 14px 20px; border-top: 1px solid var(--border); display: flex; justify-content: flex-end; gap: 10px; background: var(--surface-2); border-radius: 0 0 var(--radius) var(--radius); }
</style>
