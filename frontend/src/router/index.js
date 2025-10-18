import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RegaView from '@/views/RegaView.vue'
import LoginView from '@/views/LoginView.vue'

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
    }
  ]
})

export default router