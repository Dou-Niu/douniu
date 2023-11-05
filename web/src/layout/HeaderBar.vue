<template>
  <el-menu default-active="1" class="w-full!" mode="horizontal" background-color="bg-#131314" :ellipsis="false"
    text-color="rgba(255,255,255,0.6)" active-text-color="rgba(255,255,255,1)" router>
    <div class="user flex justify-center items-center px-4 py-4">
      <el-button type="success" round v-if="!user_info" @click="dialogVisible = true">
        登录
      </el-button>
      <div v-else class="flex justify-center items-center">
        <!-- <el-avatar :src="user_info?.avatar" /> -->
        <span class="whitespace-nowrap">Hi,{{ user_info.nickname }}</span>
        <!-- <el-button type="info" round>
          注销
        </el-button> -->
      </div>
    </div>
    <div class="grow-1 flex items-center justify-center">
      <el-input v-model="search" placeholder="" class="w-120!" @keyup.enter="handleSearch">
        <template #suffix>
          <div>
            <el-button plain class="bg-#131314! btn" @click="handleSearch" >
              <el-icon size="18">
                <Search />
              </el-icon>
              <span class="fw-600 ml-2 text-4">搜索</span>
            </el-button>
          </div>
        </template>
      </el-input>
    </div>
    <!-- <el-menu-item index="1">首页</el-menu-item> -->
    <el-menu-item v-for="item in  HeaderRoutes " :key="item.name" :index="item.path">
      <span :style="{ width: '100%', height: '100%', display: 'flex', justifyContent: 'center', alignItems: 'center' }"
        v-if="item.path === '/mine'">
        <el-avatar fit="contain" :src="user_info.avatar"></el-avatar>
      </span>
      <span v-else>{{ item.name }}</span>
    </el-menu-item>
  </el-menu>
  <LoginCard :dialogVisible="dialogVisible" @update:dialogVisible="setDialogVisible" />
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { HeaderRoutes } from "@/router/routes";
import LoginCard from "@/components/LoginCard/LoginCard.vue"
import { useRouter } from 'vue-router';
import { user } from '@/store/user'
import { video } from '@/store/video'
import { storeToRefs } from 'pinia'
import { video as videoApi } from "@/services"

const { video_list } = storeToRefs(video())
const { setVideoList } = video()

const { user_info } = storeToRefs(user());
const router = useRouter();

// 搜索框
const search = ref("")

// 登录框
const dialogVisible = ref(false)
const setDialogVisible = (v: boolean) => {
  dialogVisible.value = v
}

const handleSearch = () => {
  videoApi.searchVideo(search.value, 1).then(res => {
    setVideoList(res.data.video_list)
  })
  if (search.value) {
    router.push(`/search?key_words=${search.value}`)
  }
  search.value = ""
}

</script>

<style scoped>
.el-menu--horizontal {
  border-bottom: 1px solid #4C4D4F;
}

.el-menu-item:hover {
  background-color: #141E1F !important;
}

.el-input {
  border-color: white !important;
}

:deep(.el-input__wrapper) {
  background-color: #131314 !important;
}

:deep(.el-input__inner) {
  color: white !important;
}

.el-button.btn {
  color: white !important;
  border-radius: 0%;
  border-left: none;
  position: absolute;
  /* 绝对定位图标 */
  right: 0;
  /* 将图标放置在搜索框的右侧 */
  top: 50%;
  /* 将图标垂直居中 */
  transform: translateY(-50%);
  /* 垂直居中图标 */
}

.el-button:hover {
  background-color: white !important;
  color: black !important;
}

:deep(.is-focus) {
  border-color: red !important;
  box-shadow: 0 0 0 1px white !important;
}
</style>
