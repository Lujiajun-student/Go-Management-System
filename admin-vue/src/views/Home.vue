<script>
import storage from "@/utils/storage";

export default {
  name: "Home",
  data() {
    return {
      leftMenuList: storage.getItem("leftMenuList")
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
  }
}
</script>
run
<template>
  <el-container class="home-container">
    <el-aside class="el-aside">
      <div class="logo">
        <img src="../assets/image/logo.png" class="sidebar-logo"/>
        <h3>通用后台管理系统</h3>
      </div>
      <el-menu class="el-menu" background-color="#304156" text-color="#fff" unique-opened>
<!--        无子集菜单-->
        <el-menu-item :index="'/' + item.url" v-for="item in noChildren" :key="item.menuName">
          <i :class="item.icon"></i>
          <template slot="title">
            <span>{{item.menuName}}</span>
          </template>
        </el-menu-item>
<!--        有子集菜单-->
        <el-submenu :index="item.id + ''" v-for="item in hasChildren" :key="item.id">
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
      <el-header class="el-header">Header</el-header>
      <el-main class="el-main">Main</el-main>
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
    }
    .el-main {
      background-color: #eaedf1;
    }
  }
</style>