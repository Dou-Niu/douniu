<template>
  <div class="h-full flex flex-col bg-#060716">
    <Logo />
    <el-menu active-text-color="white" text-color="rgba(255, 255, 255, 0.5)" background-color="#060716" router
      class="h-full border-none! flex flex-col items-center" :default-active="currentRoute">
      <el-menu-item index="/index">
        <el-icon>
          <document />
        </el-icon>
        <span class="text-5">首页</span>
      </el-menu-item>
      <el-menu-item index="/recommend">
        <el-icon>
          <document />
        </el-icon>
        <span class="text-5">推荐</span>
      </el-menu-item>
      <el-menu-item index="/mine">
        <el-icon>
          <document />
        </el-icon>
        <span class="text-5">我的</span>
      </el-menu-item>
    </el-menu>
    3131231
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue"
import Logo from "@/components/SideBar/SideLogo.vue";

const currentRoute = ref("/recommend")

</script>

<style scoped>
:deep(.el-menu-item:hover) {
  /* background-color: #141E1F */
  color: white;
}
</style>
