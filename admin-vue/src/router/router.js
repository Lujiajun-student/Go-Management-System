// 封装路由

import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/views/Login.vue'
import Home from '@/views/Home.vue'
import Welcome from '@/views/Welcome.vue'

Vue.use(Router)

const router = new Router({
    // 去掉路径的#
    mode: 'history',
    routes: [
        {path: '/', redirect: '/login'},
        {path: '/login', component: Login },
        {
            path: '/home',
            component: Home,
            redirect: '/welcome',
            children: [
                {
                path: '/welcome',
                component: Welcome
                }
            ]
        }
    ]
})
export default router