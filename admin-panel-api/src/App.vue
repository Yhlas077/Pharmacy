<script setup>
import { useRoute } from 'vue-router'
import { computed } from 'vue'
import AppLayout from './components/AppLayout.vue'
import Toasts from './components/Toasts.vue'

const route = useRoute()
const isPublic = computed(() => route.meta.public)
</script>

<template>
  <Toasts />
  <component :is="isPublic ? 'div' : AppLayout">
    <router-view />
  </component>
</template>
