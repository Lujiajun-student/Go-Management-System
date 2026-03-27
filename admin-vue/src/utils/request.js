/*
封装axios
 */
import {Message} from "element-ui"
import axios from 'axios'
import router from "@/router/router";

// 创建axios对象
const service = axios.create({
    baseURL: process.env["VUE_APP_BASE_API"],
    timeout: 8000
})

// 请求拦截，加上token
service.interceptors.request.use((req) => {
    const headers = req.headers
    // todo token
    if (!headers.Authorization) {
        headers.Authorization = 'Bearer + Lu'
    }
    return req
})

// 响应拦截
service.interceptors.response.use((res) => {
    // 与后端的result结构体对应
    const {code, data, message} = res.data
    // 403 无权限
    if (code === 403) {
        Message.error(message)
        setTimeout(() => {
            // todo 清除存储信息
            router.push("/login")
        }, 1500)
    }else if (code === 406) {
        // token过期
        Message.error(message)
        setTimeout(() => {
            router.push("/login")
        }, 1500)
    }else {
        return res
    }
})

// 请求核心函数
function request(options) {
    options.method = options.method || 'get'
    if (options.method.toLowerCase() === 'get') {
        options.params = options.data
    }
    service.defaults.baseURL = process.env["VUE_APP_BASE_API"]
    return service(options)
}

export default request