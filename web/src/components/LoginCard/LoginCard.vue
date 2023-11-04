<template>
  <el-dialog v-model="visible" width="30%" destroy-on-close center style="background-color:rgb(37,38,50)">
    <template #header>
      <el-tabs v-model="activeName">
        <el-tab-pane label="登录" name="login"></el-tab-pane>
        <el-tab-pane label="注册" name="register"></el-tab-pane>
      </el-tabs>
    </template>
    <div class="flex justify-center items-center mr-6">
      <el-form label-position="right" label-width="100px" :model="formLabelAlign" class="text-center w-100 max-w-100">
        <el-form-item label="手机号">
          <el-input :maxlength="32" v-model="formLabelAlign.user" />
        </el-form-item>
        <el-form-item label="密码" v-if="loginType === '1'">
          <el-input :maxlength="32" type="password" show-password v-model="formLabelAlign.pwd" />
        </el-form-item>
        <el-form-item v-if="loginType === '2'" label="验证码">
          <div class="flex justify-center items-center">
            <el-input :maxlength="8" type="text" v-model="formLabelAlign.code" />
            <el-button type="primary" class="ml-3">发送验证码</el-button>
          </div>
        </el-form-item>
        <div v-if="activeName === 'login'">
          <div>
            <el-radio-group v-model="loginType" class="ml-4">
              <el-radio label="1" size="large">密码登录</el-radio>
              <el-radio label="2" size="large">验证码登录</el-radio>
            </el-radio-group>
          </div>
          <div>忘记密码</div>
        </div>
      </el-form>
    </div>
    <template #footer>
      <el-button type="info" @click="visible = false">
        取消
      </el-button>
      <el-button type="success" class="ml-4">
        {{ activeName == "register" ? "注册" : "登录" }}
      </el-button>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref, reactive, computed } from "vue"
const props = defineProps<{
  dialogVisible: boolean,
}>()

const emits = defineEmits(["update:dialogVisible"])
// 登录框tab
const activeName = ref("login")
// 登录方式
const loginType = ref("1")
// 登陆表单信息
const formLabelAlign = reactive({
  user: "",
  pwd: "",
  pwd2: "",
  code: ""
})

const visible = computed({
  get() {
    return props.dialogVisible
  },
  set(v: boolean) {
    emits("update:dialogVisible", v)
  }
})

</script>

<style scoped>
.el-dialog__header {
  display: flex !important;
  justify-content: center !important;
  align-items: center !important;
  border-bottom: 1px solid #4c4d4f;
}

:deep(.el-input__wrapper) {
  --el-input-border-color: none;
  background-color: rgb(51, 52, 63);
  color: white;
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

:deep(.el-tabs__item:hover) {
  color: #CFD3DC;
}

:deep(.is-active) {
  color: #CFD3DC !important;
  font-size: 18px;
}
</style>
