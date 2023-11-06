<template>
  <el-dialog v-model="visible" width="40%" destroy-on-close center style="background-color:#141414;">
    <template #header>
      <h1>{{ listType }}列表</h1>
    </template>
    <div class="flex flex-col justify-center items-start mr-6 color-white">
      <div v-for="user in list" class="flex my-4" @click="handlePush(user)">
        <el-avatar :size="40" src="/logo.jpg" class="mr-4" />
        <div class="flex flex-col whitespace-nowrap">
          <div class="text-5">{{ user.nickname }}</div>
          <div class="opacity-50">斗牛号:{{ user.id }}</div>
        </div>
      </div>
      <div v-if="list.length === 0" class="flex self-center">
        <el-empty description="一个人也没有" />
      </div>
    </div>
    <template #footer>

    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref, computed,onMounted } from 'vue';

import { user as userInfo } from '@/store/user';
import { useRouter} from 'vue-router'

const router = useRouter()

const props = defineProps<{
  dialogVisible: boolean,
  listType: string,
  list: any
}>()

const emits = defineEmits(["update:dialogVisible"])

const visible = computed({
  get() {
    return props.dialogVisible
  },
  set(v: boolean) {
    emits("update:dialogVisible", v)
  }
})

const handlePush = (user: User) => {
  router.push({
    path: `/home/${user.id}`
  })
}


onMounted(() => {

})
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
</style>