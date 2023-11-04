<template>
  <div class="h-full flex flex-col bg-#060716">
    <Logo />
    <el-menu active-text-color="white" text-color="rgba(255, 255, 255, 0.5)" background-color="#060716"
      class="h-full border-none! flex flex-col items-center">
      <template v-for="item in HomeRoutes" :key="item.path">
        <el-menu-item v-if="!item.meta?.dontShow" @click="router.push(item.path)" :class="{
          'is-active': route.path === item.path && route.path !== '/err/404'
        }">
          <el-icon>
            <document />
          </el-icon>
          <span class="text-5">{{ item.name }}</span>
        </el-menu-item>
      </template>
      <el-divider class="w-15! border-#4C4D4F!" />
      <!-- <el-menu-item index="1" @click="router.push('/channel/1')">
        <el-icon>
          <document />
        </el-icon>
        <span class="text-5">游戏</span>
      </el-menu-item>
      <el-menu-item index="2" @click="router.push('/channel/2')">
        <el-icon>
          <document />
        </el-icon>
        <span class="text-5">生活</span>
      </el-menu-item>
      <el-menu-item index="3" @click="router.push('/channel/3')">
        <el-icon>
          <document />
        </el-icon>
        <span class="text-5">影视</span>
      </el-menu-item>
      <el-menu-item index="4" @click="router.push('/channel/4')">
        <el-icon>
          <document />
        </el-icon>
        <span class="text-5">动漫</span>
      </el-menu-item>
      <el-menu-item index="5" @click="router.push('/channel/5')">
        <el-icon>
          <document />
        </el-icon>
        <span class="text-5">知识</span>
      </el-menu-item> -->
      <template v-for="item in 5" :key="item">
        <el-menu-item  @click="router.push(`/channel/${item}`)" :class="{
          'is-active': route.path === `/channel/${item}` && route.path !== '/err/404'
        }">
          <el-icon>
            <document />
          </el-icon>
          <span class="text-5">{{ routeToIndex[item-1] }}</span>
        </el-menu-item>
      </template>
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
const routeToIndex=[
  "游戏","生活","影视","动漫","知识"
]
</script>

<style scoped>
:deep(.el-menu-item:hover) {
  /* background-color: #141E1F */
  color: white;
}
</style>
