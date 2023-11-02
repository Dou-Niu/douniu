<template>
  <div class="h-full flex flex-col bg-#060716">
    <Logo />
    <el-menu active-text-color="white" text-color="rgba(255, 255, 255, 0.5)" background-color="#060716"
      class="h-full border-none! flex flex-col items-center">
      <template v-for="(item, index) in HomeRoutes" :key="index">
        <el-menu-item v-if="!item.meta?.dontShow" :index="index.toString()" @click="router.push(item.path)" :class="{
          'is-active': route.path === item.path && route.path !== '/err/404'
        }">
          <el-icon>
            <document />
          </el-icon>
          <span class="text-5">{{ item.name }}</span>
        </el-menu-item>
      </template>
      <el-divider class="w-15! border-#4C4D4F!" />
    </el-menu>
  </div>
</template>

<script setup lang="ts">
// import { ref } from "vue"
import { useRouter, useRoute } from "vue-router"
import Logo from "@/components/SideBar/SideLogo.vue";
const router = useRouter()
const route = useRoute()
import { HomeRoutes } from "@/router/routes"
</script>

<style scoped>
:deep(.el-menu-item:hover) {
  /* background-color: #141E1F */
  color: white;
}
</style>
