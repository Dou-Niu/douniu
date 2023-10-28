<template>
  <div class="h-full flex flex-col bg-#060716">
    <Logo />
    <el-menu active-text-color="white" text-color="rgba(255, 255, 255, 0.5)" background-color="#060716"
      class="h-full border-none! flex flex-col items-center" default-active="1" @open="handleOpen" @close="handleClose">
      <el-menu-item index="1">
        <el-icon>
          <document />
        </el-icon>
        <span class="text-5">推荐</span>
      </el-menu-item>
      <el-menu-item index="2">
        <el-icon>
          <document />
        </el-icon>
        <span class="text-5">我的</span>
      </el-menu-item>
    </el-menu>
    3131231
  </div>
</template>

<script setup>
import Logo from "@/components/SideBar/SideLogo.vue";
</script>

<style scoped>
:deep(.el-menu-item:hover) {
  /* background-color: #141E1F */
  color: white;
}
</style>
