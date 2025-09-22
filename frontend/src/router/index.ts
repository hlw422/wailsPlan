import { createRouter, createWebHashHistory, createWebHistory } from 'vue-router'

// 路由规则
const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/HomeView.vue')
  },
    {
    path: '/home',
    name: 'HomeMain',
    component: () => import('../views/HomeView.vue')
  }
]

// 创建路由实例
const router = createRouter({
  history: createWebHashHistory(), 
  routes: routes
})
export default router
