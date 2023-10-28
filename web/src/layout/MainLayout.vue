<template>
  <el-container class="w-100vw h-100vh bg-#131314">
    <el-aside width="200px">
      <SideBar />
    </el-aside>
    <el-container class="w-full!">
      <el-header>
        <HeaderBar />
      </el-header>
      <el-main>
        <el-empty :image-size="200" />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import SideBar from "@/layout/SideBar.vue";
import HeaderBar from "@/layout/HeaderBar.vue";
</script>

<style scoped>
:deep(.el-header) {
  --el-header-padding: 0 0 !important;
}
</style>
