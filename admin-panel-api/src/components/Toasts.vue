<script setup>
import { toasts } from '../store.js'
</script>

<template>
  <div class="toast-stack">
    <transition-group name="toast">
      <div v-for="t in toasts" :key="t.id" class="toast" :class="t.type">
        <span class="dot" />{{ t.message }}
      </div>
    </transition-group>
  </div>
</template>

<style scoped>
.toast-stack { position: fixed; top: 18px; right: 18px; z-index: 1000; display: flex; flex-direction: column; gap: 10px; }
.toast {
  display: flex; align-items: center; gap: 9px; min-width: 220px; max-width: 360px;
  padding: 12px 16px; border-radius: 10px; background: var(--surface);
  box-shadow: var(--shadow-lg); border: 1px solid var(--border); font-size: 13px; font-weight: 500;
}
.dot { width: 8px; height: 8px; border-radius: 50%; background: var(--success); flex-shrink: 0; }
.toast.error .dot { background: var(--danger); }
.toast.error { border-color: #fecaca; }
.toast-enter-active, .toast-leave-active { transition: all .25s ease; }
.toast-enter-from, .toast-leave-to { opacity: 0; transform: translateX(20px); }
</style>
