import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RegaView from '@/views/RegaView.vue'
import LoginView from '@/views/LoginView.vue'
import AddMaterialView from '../views/AddMaterialView.vue'

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
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/materials/add',
      name: 'addMaterial',
      component: AddMaterialView,
      meta: { requiresAuth: true }
    }
  ]
})

router.beforeEach((to, from, next) => {
  const isAuthenticated = document.cookie.includes('jwt=')

  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login')
  } else {
    next()
  }
})


export default router