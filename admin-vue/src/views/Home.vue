<script>
import storage from "@/utils/storage";

export default {
  name: "Home",
  data() {
    return {
      leftMenuList: storage.getItem("leftMenuList"),
      activePath: '',
      collapseBtnClass: "el-icon-s-fold",
      isCollapse: false,
    }
},
  computed: {
    // 无子集
    noChildren() {
      return this.leftMenuList.filter(item => !item.menuSVoList)
    },
    // 有子集
    hasChildren() {
      return this.leftMenuList.filter(item => item.menuSVoList)
    }
  },
  methods: {
    // 保持路由激活
    saveNavState(activePath) {
      storage.setItem('activePath', activePath)
      this.activePath = activePath
    },
    // 顶部栏展开和折叠
    toggleCollapse() {
      this.isCollapse = !this.isCollapse
      if (this.isCollapse) {
        this.collapseBtnClass = 'el-icon-s-unfold'
      } else {
        this.collapseBtnClass = 'el-icon-s-fold'
      }
    }
  }
}
</script>
run
<template>
  <el-container class="home-container">
    <el-aside class="el-aside" :width="isCollapse? '64px' : '200px' ">
      <div class="logo">
        <img src="../assets/image/logo.png" class="sidebar-logo"/>
        <h3 v-show="!isCollapse">通用后台管理系统</h3>
      </div>
      <el-menu class="el-menu" background-color="#304156" text-color="#fff" unique-opened router :default-active="$route.path"
               :collapse="isCollapse" :collapse-transition="false">
<!--        无子集菜单-->
        <el-menu-item :index="'/' + item.url" v-for="item in noChildren" :key="item.menuName"
        @click="saveNavState('/' + item.url)">
          <i :class="item.icon"></i>
          <template slot="title">
            <span>{{item.menuName}}</span>
          </template>
        </el-menu-item>
<!--        有子集菜单-->
        <el-submenu :index="item.id + ''" v-for="item in hasChildren" :key="item.id"
                    @click="saveNavState('/' + item.url)">
          <template slot="title">
            <i :class="item.icon"></i>
            <span>{{ item.menuName }}</span>
          </template>
          <el-menu-item :index="'/' + subItem.url" v-for="subItem in item.menuSVoList" :key="subItem.id">
            <template slot="title">
              <i :class="subItem.icon"></i>
              <span>{{subItem.menuName}}</span>
            </template>
          </el-menu-item>
        </el-submenu>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header class="el-header">
<!--        顶部栏设置-->
        <div class="fold-btn">
          <i :class="collapseBtnClass" @click="toggleCollapse"></i>
        </div>
<!--        面包屑功能-->
        <div class="bread-btn">
<!--          当前处于首页，面包屑首个元素为首页-->
          <el-breadcrumb separator="/" v-if="$router.currentRoute.path !== '/welcome'">
            <el-breadcrumb-item :to="{path: '/welcome'}">首页</el-breadcrumb-item>
            <el-breadcrumb-item>{{$route.meta.sTitle}}</el-breadcrumb-item>
            <el-breadcrumb-item>{{$route.meta.tTitle}}</el-breadcrumb-item>
          </el-breadcrumb>
<!--          不在首页，不需要显示首页-->
          <el-breadcrumb separator="/" v-else>
            <el-breadcrumb-item>{{$route.meta.sTitle}}</el-breadcrumb-item>
            <el-breadcrumb-item>{{$route.meta.tTitle}}</el-breadcrumb-item>
          </el-breadcrumb>
        </div>
      </el-header>
      <el-main class="el-main">
        <router-view/>
      </el-main>
    </el-container>
  </el-container>
</template>

<style lang = "less" scoped>
  .home-container {
    height: 100%;
    .el-aside {
      background-color: #304156;
      .logo {
        margin-top: 5px;
        display: flex;
        align-items: center;
        font-size: 13px;
        height: 50px;
        color: #fff;
        font-style: italic;
        .sidebar-logo {
          width: 32px;
          height: 32px;
          margin: 0 16px;
        }
      }
      .el-menu {
        border-right: none;
      }
    }
    .el-header {
      background-color: #f9fafc;
      align-items: center;
      justify-content: space-between;
      display: flex;
      .fold-btn {
        padding-top: 2px;
        font-size: 23px;
        cursor: pointer;
      }
    .bread-btn {
      padding-top: 2px;
      position: fixed;
      margin-left: 40px;
    }
    }
    .el-main {
      background-color: #eaedf1;
    }
  }
</style>