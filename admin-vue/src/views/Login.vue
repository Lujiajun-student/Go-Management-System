
<template>
  <div class="login_container">
    <div class="login_box">
      <el-form class="login_form" ref="loginFormRef" :rules="rules" :model="loginForm">
        <div class="title">
          通用后台管理系统
        </div>
        <el-form-item prop="username">
          <el-input placeholder="账号" prefix-icon="el-icon-user-solid" v-model="loginForm.username" clearable></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input placeholder="密码" prefix-icon="el-icon-key" v-model="loginForm.password" clearable></el-input>
        </el-form-item>
        <el-form-item prop="image">
          <el-input placeholder="验证码" prefix-icon="el-icon-view" style="width: 200px; float: left; " maxlength="6" v-model="loginForm.image" clearable/>
          <el-image class="captchaImg" style="width: 150px; float: left;" :src="image" @click="getCaptcha"/>
        </el-form-item>
        <el-form-item>
          <el-row :gutter="20">
            <el-col :span="12" :offset="0">
              <el-button type="primary" style="width: 100%; font-size: large;" @click="loginBtn">登录</el-button>
            </el-col>
            <el-col :span="12" :offset="0">
              <el-button type="info" style="width: 100%; font-size: large;" @click="resetForm">重置</el-button>
            </el-col>
          </el-row>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      image: '',
      rules: {
        username: [
          {
            required: true, message:"请输入账号", trigger: "blur"
          }
        ],
        password: [
          {
            required: true, message:"请输入密码", trigger: "blur"
          }
        ],
        image: [
          {
            required: true, message:"请输入验证码", trigger: "blur"
          }
        ]
      },
      loginForm: {
        username: '',
        password: '',
        image: '',
        id_key: ''
      }
    }
  },
  methods: {
    // 获取验证码
    async getCaptcha() {
      const {data: res} = await this.$api.captcha()
      if (res.code !== 200) {
        this.$message.error(res.message)
      }else {
        this.image = res.data.image
        // 封装验证码id
        this.loginForm.id_key = res.data.idKey
      }
    },
    // 重置表单
    resetForm() {
      this.$refs.loginFormRef.resetFields()
    },
    // 登录
    loginBtn() {
      this.$refs.loginFormRef.validate(async valid => {
        // console.log("传输参数：", this.loginForm)
        if (valid) {
          const {data: res} = await this.$api.login(this.loginForm)
          // console.log("获取登录的res数据：", res)
          if (res.code !== 200) {
            this.$message.error(res.message)
          } else {
            this.$message.success("登录成功")
            this.$store.commit('saveSysAdmin', res.data.sysAdmin)
            this.$store.commit('saveToken', res.data.token)
            this.$store.commit('saveLeftMenuList', res.data.leftMenuList)
            this.$store.commit('savePermissionList', res.data.permissionList)
            this.$router.push("/home")
          }
        } else {
          return false
        }
      })
    }
  },

  created() {
    this.getCaptcha()
  }
}

</script>

<style lang="less" scoped>
  .login_container {
    background-image: url("../assets/image/login-background.jpg");
    // 拉伸背景图片
    background-size: cover;
    height: 100%;
    .login_box {
      width: 400px;
      height: 330px;
      background: #fff;
      border-radius: 1px;
      position: absolute;
      left: 50%;
      top: 50%;
      transform: translate(-50%, -50%);
      .login_form {
        position: absolute;
        bottom: 0;
        width: 100%;
        padding: 0 20px;
        box-sizing: border-box;
        .title {
            font-size: 23px;
            line-height: 1.5;
            text-align: center;
            margin-bottom: 20px;
            font-weight: bold;
            font-style: italic;
        }
        .captchaImg {
          height: 38px;
          width: 100%;
          border: 1px solid #e6e6e6;
          margin-left: 8px;
        }
      }
    }
  }

</style>