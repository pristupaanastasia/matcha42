import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Profile from '../views/Profile.vue'
import Sendmail from '../views/Sendmail.vue';

const routes = [
    {
      path: "/",
      name: "Home",
      component: Home
    },
    {
      path: "/profile",
      name: "Profile",
      component: Profile,

    },
    {
      path: "/login",
      name: "Login",
      component: Login,

    },
    {
      path: "/register",
      name: "Register",
      component: Register,

    },
    {
        path: "/sendmail",
        name: "Sendmail",
        component: Sendmail,

    }
    ];
const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})
export default router
