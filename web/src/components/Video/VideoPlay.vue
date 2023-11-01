<template>
  <div class="w-full h-full container flex">
    <!-- <video src="@/assets/1.mp4" autoplay controls poster="@/assets/logo.png" class="w-full h-150" ref="videoRef"></video> -->
    <div id="video">
      <div class="info fw-600 text-6">
        <h1>视频信息</h1>
      </div>
      <div class="right-side text-black">
        <div class="my-6! flex flex-col items-center justify-center">
          <el-popover placement="left-start" :width="20" trigger="hover" content="点赞"
            popper-class="bg-#33343F! border-none! text-white! icon">
            <template #reference>
              <el-icon class="color-white!" :size="40">
                <Star />
              </el-icon>
            </template>
          </el-popover>
          <span class="color-white">
            111万
          </span>
        </div>
        <div @click="showComments = !showComments" class="my-6! flex flex-col items-center justify-center">
          <el-popover placement="left-start" :width="100" trigger="hover" content="评论"
            popper-class="bg-#33343F! border-none! text-white! icon">
            <template #reference>
              <el-icon class="color-white!" :size="40">
                <ChatDotRound />
              </el-icon>
            </template>
          </el-popover>
          <span class="color-white">
            111万
          </span>
        </div>
        <div class="my-6! flex flex-col items-center justify-center">
          <el-popover placement="left-start" :width="100" trigger="hover" content="收藏"
            popper-class="bg-#33343F! border-none! text-white! icon">
            <template #reference>
              <el-icon class="color-white!" :size="40">
                <Star />
              </el-icon>
            </template>
          </el-popover>
          <span class="color-white">
            111万
          </span>
        </div>
        <div class="my-6! flex flex-col items-center justify-center">
          <el-popover placement="left-start" :width="100" trigger="hover" content="分享"
            popper-class="bg-#33343F! border-none! text-white! icon">
            <template #reference>
              <el-icon class="color-white!" :size="40">
                <BottomRight />
              </el-icon>
            </template>
          </el-popover>
          <span class="color-white">
            111万
          </span>
        </div>
      </div>
    </div>
    <div v-if="showComments">
      <el-tabs v-model="activeName">
        <el-tab-pane label="TA的作品" name="ta">TA的作品</el-tab-pane>
        <el-tab-pane label="评论" name="comments">
          <Comments />
        </el-tab-pane>
        <el-tab-pane label="相关推荐" name="recommend">相关推荐</el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import Comments from "@/components/Video/Comments.vue"
// import { formatSeconds } from "@/utils/format"

import Player from 'xgplayer';
import 'xgplayer/dist/index.min.css';

const player = ref<Player>()
const showComments = ref<boolean>(false)

const activeName = ref<string>("comments")
onMounted(() => {
  player.value = new Player({
    id: "video",
    el: document.getElementById("video") as HTMLElement,
    width: "100%",
    height: "100%",
    url: "//lf3-static.bytednsdoc.com/obj/eden-cn/nupenuvpxnuvo/xgplayer_doc/xgplayer-demo.mp4",
    playsinline: true,
    poster: "//lf9-cdn-tos.bytecdntp.com/cdn/expire-1-M/byted-player-videos/1.0.0/poster.jpg",
    lang: 'zh',
    autoplay: true,
    marginControls: false,
    dynamicBg: {
      disable: false
    }
  });
})
</script>

<style scoped>
.container {
  overflow: hidden;
}

.info{
  position: absolute;
  left: 40px;
  bottom: 160px;
}

.right-side {
  position: absolute;
  right: 40px;
  bottom: 100px;
}

:deep(.el-tabs__item) {
  color: #526368;
  font-size:20px;
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
  background-color: red;
  height: 2px;
  border-radius: 0;
  margin-top: -2px;
}

:deep(.el-tabs__item:hover) {
  color: #CFD3DC;
}

:deep(.is-active) {
  color: #CFD3DC;
}</style>
