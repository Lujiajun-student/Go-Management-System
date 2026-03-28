// 封装路由

import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/views/Login.vue'
import Home from '@/views/Home.vue'
import Welcome from '@/views/Welcome.vue'
import storage from "@/utils/storage";
import Personal from '@/views/Personal'
import Admin from '@/views/base/Admin'
import Role from '@/views/base/Role'
import Dept from '@/views/base/Dept'
import Post from '@/views/base/Post'
import LoginLog from '@/views/monitor/LoginLog'
import Operator from '@/views/monitor/Operator'
import Menu from '@/views/base/Menu'

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
                },
                {
                    path: '/personal',
                    component: Personal
                },
                {
                    path: '/base/admin',
                    component: Admin
                },
                {
                    path: '/base/role',
                    component: Role
                },
                {
                    path: '/base/menu',
                    component: Menu
                },
                {
                    path: '/base/dept',
                    component: Dept
                },
                {
                    path: '/base/post',
                    component: Post
                },
                {
                    path: '/monitor/loginLog',
                    component: LoginLog
                },
                {
                    path: '/monitor/operator',
                    component: Operator
                }
            ]
        }
    ]
})

// 挂载路由导航数据
router.beforeEach((to, from, next) => {
    if (to.path === '/login') {
        return next()
    }
    const tokenStr = storage.getItem("token")
    if (!tokenStr) {
        return next('/login')
    }
    next()
})

export default router