<template>
  <el-menu default-active="1" class="w-full!" mode="horizontal" background-color="bg-#131314"
    :ellipsis="false" text-color="rgba(255,255,255,0.6)" active-text-color="rgba(255,255,255,1)">
    <div class="user flex justify-center items-center px-4 py-4">
      <el-button type="success" round v-if="!isLogin" @click="dialogVisible = true">
        登录
      </el-button>
      <div v-else class="flex justify-center items-center info">
        <el-avatar :src="info?.avatar" />
        <span>Hi,{{ info?.name }}</span>
        <a @click="logout">注销</a>
      </div>
    </div>
    <div class="grow-1" />
    <el-menu-item index="1">首页</el-menu-item>
  </el-menu>
  <el-dialog v-model="dialogVisible" width="40%" destroy-on-close center style="background-color:#141414;">
    <template #header>
      <el-tabs v-model="activeName">
        <el-tab-pane label="登录" name="login"></el-tab-pane>
        <el-tab-pane label="注册" name="register"></el-tab-pane>
      </el-tabs>
    </template>
    <div class="flex justify-center items-center">
      <el-form label-position="right" label-width="100px" :model="formLabelAlign" class="text-center w-100 max-w-100">
        <el-form-item label="手机号">
          <el-input :maxlength="32" v-model="formLabelAlign.user" />
        </el-form-item>
        <el-form-item label="密码" v-if="loginType==='1'">
          <el-input :maxlength="32" type="password" show-password v-model="formLabelAlign.pwd" />
        </el-form-item>
        <el-form-item v-if="activeName === 'register'" label="重复">
          <el-input :maxlength="32" type="password" show-password v-model="formLabelAlign.pwd2" />
        </el-form-item>
        <el-form-item v-if="loginType==='2'" label="验证码">
          <div class="flex justify-center items-center">
            <el-input :maxlength="8" type="text" v-model="formLabelAlign.code" />
            <el-button type="primary" class="ml-3">发送验证码</el-button>
          </div>
        </el-form-item>
        <div v-if="activeName === 'login'">
          <el-radio-group v-model="loginType" class="ml-4">
            <el-radio label="1" size="large">密码登录</el-radio>
            <el-radio label="2" size="large">验证码登录</el-radio>
          </el-radio-group>
        </div>
      </el-form>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="info" @click="dialogVisible = false">
          取消
        </el-button>
        <el-button type="success" @click="login">
          {{ activeName == "register" ? "注册" : "登录" }}
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive } from "vue";
const dialogVisible = ref(true);
const activeName = ref("login");
// 登录方式
const loginType=ref("1");
const formLabelAlign = reactive({
  user: "",
  pwd: "",
  pwd2: "",
  code: ""
});
</script>

<style scoped>
.el-menu--horizontal{
  border-bottom: 1px solid #4C4D4F;
}
.el-menu-item:hover {
  background-color: #141E1F !important;
}

.el-dialog__header {
  display: flex !important;
  justify-content: center !important;
  align-items: center !important;
  border-bottom: 1px solid #4c4d4f;
}

:deep(.el-form-item__label) {
  color: #CFD3DC;
}


:deep(.el-tabs__nav-wrap::after) {
  color: #414243;
  background-color: #414243;
}

:deep(.el-tabs__nav-scroll) {
  display: flex;
  justify-content: center;
  align-items: center;
}

:deep(.el-tabs__active-bar) {
  background-color: #CFD3DC;
  height: 2px;
  border-radius: 0;
  margin-top: -2px;
}

:deep(.el-tabs__item) {
  color: #CFD3DC;
  font-size: 18px;
}

.info>a {
  font-size: 11px;
  cursor: pointer;
  text-decoration: underline;

  &:hover {
    color: #cd3420;
  }
}

.info>* {
  margin: 0 3px;
}

</style>
