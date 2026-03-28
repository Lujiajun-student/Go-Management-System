<script>
  export default {
    name: "Tags",
    data() {
      return {
        tags:[{
          title: "首页",
          path: "/welcome"
        }]
      }
    },
    watch: {
      $route: {
        immediate: true,
        handler(val) {
          // 查看新的页面是否在当前的tags数组下
          const boolean = this.tags.find(item => {
            return val.path === item.path
          })
          // 如果不在，在tags下添加这个新的页面
          if (!boolean) {
            this.tags.push({
              title: val.meta.tTitle,
              path: val.path
            })
          }
        }
      }
    }
  }
</script>

<template>
  <div class="tags">
    <el-tag class="tag" size="medium" closable :effect="item.title === $route.meta.tTitle ? 'dark' : 'plain'" v-for="item in tags" :key="item.path">
    {{item.title}}
    </el-tag>
  </div>
</template>

<style scoped lang="less">
  .tags {
    padding-left: 20px;
    padding-top: 2px;
    padding-bottom: 2px;
  }
  .tag {
    cursor: pointer;
    margin-right: 3px;
  }
</style>