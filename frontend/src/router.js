import Vue from 'vue';
import Router from 'vue-router';
import Dashboard from './components/Dashboard.vue';
import Settings from './components/Settings.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    { path: '/', name: 'Dashboard', component: Dashboard },
    { path: '/settings', name: 'Settings', component: Settings },
  ],
});
