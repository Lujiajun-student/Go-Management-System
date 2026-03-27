/*
后端api接口管理
 */

import request from '@/utils/request'

export default {
    captcha() {
        return request ({
            url: '/captcha',
            method: 'get'
        })
    }
}