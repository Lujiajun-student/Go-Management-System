/*
后端api接口管理
 */

import request from '@/utils/request'

export default {
    // 验证码接口
    captcha() {
        return request ({
            url: '/captcha',
            method: 'get'
        })
    },
    // 登录接口
    login(params) {
        return request({
            url: '/login',
            method: 'post',
            data: params
        })
    },
    // post岗位
    queryPostList(params) {
        return request({
            url: '/post/list',
            method: 'get',
            data: params
        })
    },
    // 批量删除岗位
    batchDeleteSysPost(ids) {
        const data = {
            ids
        }
        return request({
            url: '/post/batch/delete',
            method: 'delete',
            data: data
        })
    },
    // 根据id删除岗位
    deleteSysPost(id) {
        const data = {
            id
        }
        return request({
            url: '/post/delete',
            method: 'delete',
            data: data
        })
    },
    // 获取岗位列表
    querySysPostVOList() {
        return request({
            url: '/post/vo/list',
            method: 'get'
        })
    },
    // 添加岗位
    addPost(data) {
        return request({
            url: '/post/add',
            method: 'post',
            data: data
        })
    },
    // 根据id获取岗位信息
    postInfo(id) {
        const data = {
            id
        }
        return request({
            url: '/post/info',
            method: 'get',
            data: data
        })
    },
    // 更新岗位
    updatePost(data) {
        return request({
            url: '/post/update',
            method: 'put',
            data: data
        })
    },
    // 更新岗位状态
    updatePostStatus(id, postStatus) {
        const data = {
            id,
            postStatus
        }
        return request({
            url: '/post/updateStatus',
            method: 'put',
            data: data
        })
    }
}