<template>
  <el-dialog v-model="visible" width="45%" destroy-on-close center style="background-color:#141414;">
    <template #header>
      <el-tabs v-model="activeName">
        <el-tab-pane label="登录" name="login"></el-tab-pane>
        <el-tab-pane label="注册" name="register"></el-tab-pane>
      </el-tabs>
    </template>
    <div class="flex justify-center items-center mr-6">
      <el-form ref="formRef" :rules="rules" label-position="right" label-width="100px" :model="formLabelAlign"
        class="text-center w-100 max-w-100">
        <el-form-item label="手机号" prop="phone">
          <el-input :maxlength="32" v-model="formLabelAlign.phone" />
        </el-form-item>
        <el-form-item label="密码" v-if="loginType === '1' && activeName !== 'register'" prop="pwd">
          <el-input :maxlength="32" type="password" show-password v-model="formLabelAlign.pwd" />
        </el-form-item>
        <!-- <el-form-item v-if="activeName === 'register' " label="重复" prop="pwd2">
          <el-input :maxlength="32" v-model="formLabelAlign.pwd2" />
        </el-form-item> -->
        <el-form-item v-if="loginType === '2' || activeName === 'register'" label="验证码">
          <div class="flex justify-center items-center">
            <el-input :maxlength="8" type="text" v-model="formLabelAlign.code" />
            <el-button type="primary" class="ml-3" @click="handleSendCode">发送验证码</el-button>
          </div>
        </el-form-item>
        <div v-if="activeName === 'login'">
          <el-radio-group v-model="loginType" class="ml-4">
            <el-radio label="1" size="large">密码登录</el-radio>
            <el-radio label="2" size="large">验证码登录</el-radio>
          </el-radio-group>
          <div>忘记密码</div>
        </div>
      </el-form>
    </div>
    <template #footer>
      <el-button type="info" @click="visible = false">
        取消
      </el-button>
      <el-button type="success" class="ml-4" @click="handleLogin">
        {{ activeName == "register" ? "注册" : "登录" }}
      </el-button>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { login } from '@/services';
import { FormInstance, FormRules, ElMessage } from 'element-plus';
import { ref, reactive, computed } from 'vue';

import { user as userInfo } from '@/store/user';
const userStore = userInfo();

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
  phone: "",
  pwd: "",
  pwd2: "",
  code: ""
})
const formRef = ref<FormInstance>();

const validatePass = (_: any, value: any, callback: any) => {
  if (formLabelAlign.pwd !== '') {
    if (value !== formLabelAlign.pwd) {
      callback(new Error('两次输入密码不一致！'))
    }
  }
  callback()
}
const rules: FormRules = {
  phone: [
    { required: true, trigger: 'blur', message: '手机号不能为空' }
  ],
  pwd: [
    { required: true, trigger: 'blur', message: '密码不能为空' },
    { min: 6, trigger: 'blur', message: '密码最短6位' }
  ],
  pwd2: [
    { required: true, trigger: 'blur', message: '密码不能为空' },
    { min: 6, trigger: 'blur', message: '密码最短6位' },
    { validator: validatePass, trigger: 'blur', message: '与第一次输入密码不一致' }
  ]
}

const visible = computed({
  get() {
    return props.dialogVisible
  },
  set(v: boolean) {
    emits("update:dialogVisible", v)
  }
})

// 发送验证码
const handleSendCode = () => {
  try {
    formRef.value?.validate(async (isOk: boolean) => {
      if (isOk) {
        let res = await login.getCode(formLabelAlign.phone);
        console.log('%c [ res ]-105', 'font-size:13px; background:pink; color:#bf2c9f;', res)
      }
    });
  } catch (error: any) {
    ElMessage.error(error.message);
  }
}

const handleLogin = async () => {
  if (loginType.value === '1') {
    // 密码登录
    let res = await login.loginByPass(formLabelAlign.phone, formLabelAlign.pwd);
    userStore.setUserId(res.data.user_id.toString());
    userStore.setToken('access_token', res.data.access_token);
    userStore.setToken('refresh_token', res.data.refresh_token);
    userStore.getUserInfo();
  } else {
    // 验证码登录
    let res = await login.loginByCode(formLabelAlign.phone, formLabelAlign.code);
    // 存储token信息
    console.log(res.data);
    userStore.setUserId(res.data.user_id.toString());
    userStore.setToken('access_token', res.data.access_token);
    userStore.setToken('refresh_token', res.data.refresh_token);
    userStore.getUserInfo();
  }
  visible.value = false;
}

</script>

<style scoped>
.el-dialog__header {
  display: flex !important;
  justify-content: center !important;
  align-items: center !important;
  border-bottom: 1px solid #4c4d4f;
}

:deep(.el-form-item__label) {
  color: #CFD3DC;
}

:deep(.el-input__wrapper) {
  --el-input-border-color: none;
  background-color: rgb(51, 52, 63);
  color: white;
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
  color: #CFD3DC;
  font-size: 18px;
}
</style>
@/services