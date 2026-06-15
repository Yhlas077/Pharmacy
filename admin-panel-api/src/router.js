import { createRouter, createWebHistory } from 'vue-router'
import { useAuth } from './store.js'

import LoginView from './views/LoginView.vue'
import DashboardView from './views/DashboardView.vue'
import CrudView from './views/CrudView.vue'
import OrdersView from './views/OrdersView.vue'

const routes = [
  { path: '/login', name: 'login', component: LoginView, meta: { public: true } },
  { path: '/', redirect: '/dashboard' },
  { path: '/dashboard', name: 'dashboard', component: DashboardView },
  { path: '/users', name: 'users', component: CrudView, props: { resource: 'users' } },
  { path: '/categories', name: 'categories', component: CrudView, props: { resource: 'categories' } },
  { path: '/medicines', name: 'medicines', component: CrudView, props: { resource: 'medicines' } },
  { path: '/pharmacies', name: 'pharmacies', component: CrudView, props: { resource: 'pharmacies' } },
  { path: '/orders', name: 'orders', component: OrdersView },
]

const router = createRouter({ history: createWebHistory(), routes })

router.beforeEach((to) => {
  const { isAuthed } = useAuth()
  if (!to.meta.public && !isAuthed) return { name: 'login', query: { redirect: to.fullPath } }
  if (to.name === 'login' && isAuthed) return { name: 'dashboard' }
})

export default router
