<template>
  <div class="w-full h-full overflow-hidden flex" @wheel="handleWheel">
    <div id="video" ref="videoRef">
      <div class="info fw-600">
        <div class="text-6 flex gap-5 items-center">@视频信息
          <el-icon :size="20" class="color-red!">
            <Plus />
          </el-icon>
        </div>
        <div class="text-4">视频简介</div>
      </div>
      <div class="right-side text-black">
        <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" :size="50" />
        <div class="my-6! flex flex-col items-center justify-center">
          <el-popover placement="left-start" :width="20" trigger="hover" content="点赞"
            popper-class="bg-#33343F! border-none! text-white! icon">
            <template #reference>
              <el-icon class="color-white!" :size="40" @click="handleFavorite">
                <Star />
              </el-icon>
            </template>
          </el-popover>
          <span class="color-white">
            111万
          </span>
        </div>
        <div class="my-6! flex flex-col items-center justify-center">
          <el-popover placement="left-start" :width="20" trigger="hover" content="收藏"
            popper-class="bg-#33343F! border-none! text-white! icon">
            <template #reference>
              <el-icon class="color-white!" :size="40" @click="handleCollect">
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
        <!-- <div class="my-6! flex flex-col items-center justify-center">
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
        </div> -->
        <div class="my-6! flex flex-col items-center justify-center">
          <el-popover placement="left-start" :width="100" trigger="hover" content="分享"
            popper-class="bg-#33343F! border-none! text-white! icon">
            <template #reference>
              <el-icon class="color-white!" :size="40" @click="handleShare">
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
      <div class="w-full flex justify-center items-center text-6">评论</div>
      <el-divider class="border-#4C4D4F!" />
      <div>全部评论&nbsp;&nbsp;{{ 1 }}</div>
      <el-scrollbar>
        <Comments />
      </el-scrollbar>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import { useRoute } from "vue-router"
import { favorite, shareVideo, collection } from "@/services/operation"
import Comments from "@/components/Video/Comments.vue"

import Player from 'xgplayer';
import 'xgplayer/dist/index.min.css';

const route = useRoute()
const player = ref<Player>()
const showComments = ref<boolean>(false)
const videoRef = ref()

// 鼠标滚动
const handleWheel = (event) => {
  if (event.deltaY < 0) {
    // 上滚
    console.log("上滚");
  } else {
    // 下滚
    console.log("下滚");
    console.log(player.value?.url);
    console.log(videoRef.value.children[1].children);
  }
}
// 操作
const handleFavorite = async () => {
  let res = await favorite(1, 1)
  console.log(res);
}

const handleCollect = async () => {
  let res = await collection(1, 1, 1)
  console.log(res);
}

const handleShare = async () => {
  let res = await shareVideo(1)
  console.log(res);
}
onMounted(() => {
  player.value = new Player({
    id: "video",
    el: document.getElementById("video") as HTMLElement,
    width: "100%",
    height: "100%",
    url: [
      {
        src: "https://www.kecat.top/video/jingliu.mp4",
        type: "video/mp4"
      },
      {
        src: "//lf3-static.bytednsdoc.com/obj/eden-cn/nupenuvpxnuvo/xgplayer_doc/xgplayer-demo.mp4",
        type: "video/mp4"
      }
    ],
    playsinline: true,
    poster: "//lf9-cdn-tos.bytecdntp.com/cdn/expire-1-M/byted-player-videos/1.0.0/poster.jpg",
    lang: 'zh',
    autoplay: true,
    marginControls: false,
    dynamicBg: {
      disable: false
    }
  });
  console.log(route.query.url);
})
</script>

<style scoped>
.info {
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
  font-size: 20px;
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
}
</style>
