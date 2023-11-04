<template>
  <MainLayout>
    <router-view></router-view>
  </MainLayout>
</template>

<script setup>
import MainLayout from "@/layout/MainLayout.vue"
</script>

<style>
body {
  padding: 0;
  margin: 0;
  color:white;
}

* {
  font-family: PingFang SC,DFPKingGothicGB-Regular,sans-serif;
  outline: none;
  user-select: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  -webkit-user-drag: none;

  --el-input-focus-border-color:none;
  --el-input-text-color:white;
}
</style>
