import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/account',
      name: 'account',
      component: () => import('../views/AccountPage.vue')
    },
    {
      path: '/games/all',
      name: 'allGames',
      component: () => import('../views/AllGamesPage.vue')
    },
    {
      path: '/games/info/:id',
      name: 'info',
      component: () => import('../views/GameInfo.vue')
    },
    {
      path: '/admin/login',
      name: 'login',
      component: () => import('../views/LoginPage.vue')
    },
    {
      path: '/admin/games/manage',
      name: 'manageGames',
      component: () => import('../views/GameManagementPage.vue')
    },
    {
      path: '/admin/users',
      name: 'users',
      component: () => import('../views/UsersPage.vue')
    },
  ]
})

export default router
