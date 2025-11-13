import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import HomeView from '@/views/HomeView.vue'
import RegaView from '@/views/RegaView.vue'
import LoginView from '@/views/LoginView.vue'
import AddMaterialView from '@/views/AddMaterialView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/rega',
      name: 'rega',
      component: RegaView
      // meta: { requiresGuest: true } 
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
      // meta: { requiresGuest: true } 
    },
    {
      path: '/materials/add',
      name: 'addMaterial',
      component: AddMaterialView,
      meta: { requiresAuth: true }
    },
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/login')
    return
  }

  if (to.meta.requiresGuest && authStore.isAuthenticated) {
    next('/')
    return
  }

  next()
})


export default router