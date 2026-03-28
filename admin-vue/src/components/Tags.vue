<script>

  export default {
    name: "Tags",
    computed: {
      index() {
        return index
      }
    },
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
    },
    methods: {
      // 点击跳转
      goTo(path) {
        this.$router.push(path)
      },
      // 点击关闭
      close(i) {
        // 如果关闭当前激活的页面，且当前页面不是最后一个，那么关闭后会跳转到最后一个页面
        if (this.tags[i].path === this.$route.meta.path && i !== this.tags.length - 1) {
          this.$router.push(this.tags[this.tags.length - 1].path)
        }else if (i === this.tags.length - 1) {
          // 如果当前激活的页面是最后一个，关闭的是最后一个页面，那么关闭后会跳转到原本倒数第二个页面
          this.$router.push(this.tags[this.tags.length - 2].path)
        }
        this.tags.splice(i, 1)
      }
    }
  }
</script>

<template>
  <div class="tags">
    <el-tag class="tag" size="medium" closable :effect="item.title === $route.meta.tTitle ? 'dark' : 'plain'" v-for="(item, i) in tags" :key="item.path"
    @click="goTo(item.path)" @close="close(i)" :closable="i >0" disable-transitions >
      <i class="circular" v-show="item.title === $route.meta.tTitle"></i>
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
    .circular {
      width: 8px;
      height: 8px;
      margin-right: 4px;
      background-color: #fff;
      border-radius: 50%;
      display: inline-block;
    }
  }
</style>