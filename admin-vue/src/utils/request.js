/*
封装axios
 */
import {Message} from "element-ui"
import axios from 'axios'
import router from "@/router/router";
import storage from "@/utils/storage";

// 创建axios对象
const service = axios.create({
    baseURL: process.env["VUE_APP_BASE_API"],
    timeout: 8000
})

// 请求拦截，加上token
service.interceptors.request.use((req) => {
    const headers = req.headers
    // 从localStorage中获取token
    const token = storage.getItem("token") || {}

    if (!headers.Authorization) {
        headers.Authorization = 'Bearer ' + token
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
            // 清除localStorage
            storage.clearAll()
            router.push("/login")
        }, 1500)
    }else if (code === 406) {
        // token过期
        Message.error(message)
        setTimeout(() => {
            // 清空localStorage
            storage.clearAll()
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