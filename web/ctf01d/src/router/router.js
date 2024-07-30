import { createWebHistory, createRouter } from 'vue-router'

import WelcomeView from '../components/welcome/welcomeView.vue'
import GamesView from '../components/games/gamesView.vue'
import ServicesView from '../components/services/servicesView.vue'
import TeamsView from '../components/teams/teamsView.vue'
import UsersView from '../components/users/usersView.vue'

const routes = [
    { path: '/', component: WelcomeView },
    { path: '/games', component: GamesView },
    { path: '/services', component: ServicesView },
    { path: '/teams', component: TeamsView },
    { path: '/users', component: UsersView },
    { path: '/:pathMatch(.*)*', component: WelcomeView}
  ]
  
  const router = createRouter({
    history: createWebHistory(),
    routes,
  })

  export default router;

  