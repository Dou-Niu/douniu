<template>
  <div class="w-full h-full overflow-hidden flex scrolling">
    <div id="video" ref="videoRef" @wheel="handleWheel">
      <div class="info fw-600">
        <div class="text-6 flex gap-5 items-center">@{{ currentVideo?.author.nickname }}
          <div v-if="parseInt(user_id) != currentVideo.author.id">
            <el-button type="danger" icon="Plus" circle plain v-if="!isFollow"
              @click="handleFollow" />
            <el-button type="success" icon="Check" circle v-else-if="isFollow"
              @click="handleFollow" />
          </div>
          <div v-else-if="parseInt(user_id) == currentVideo.author.id">自己</div>
        </div>
        <div class="text-4">{{ currentVideo?.title }}</div>
      </div>
      <div class="right-side text-black">
        <el-avatar :src="currentVideo?.author.avatar" :size="50"
          @click="router.push(`/home/${currentVideo.author.id}`)" />
        <div class="my-6! flex flex-col items-center justify-center">
          <el-popover placement="left-start" :width="20" trigger="hover" content="点赞"
            popper-class="bg-#33343F! border-none! text-white! icon">
            <template #reference>
              <el-image style="width: 40px; height: 40px" src="/img/favorite-fill.png" fit="cover" v-if="isFavorite"
                @click="handleFavorite" />
              <el-image style="width: 40px; height: 40px" src="/img/favorite.png" fit="cover" v-else-if="!isFavorite"
                @click="handleFavorite" />
              <!-- <el-icon class="color-white!" :size="40" @click="handleFavorite">
                <Star />
              </el-icon> -->
            </template>
          </el-popover>
          <span class="color-white">
            {{ currentVideo?.favorite_count +Number(isFavorite)}}
          </span>
        </div>
        <div class="my-6! flex flex-col items-center justify-center">
          <el-popover placement="left-start" :width="20" trigger="hover" content="收藏"
            popper-class="bg-#33343F! border-none! text-white! icon">
            <template #reference>
              <el-image style="width: 40px; height: 40px" src="/img/collect.png" fit="cover" @click="handleCollect"
                v-if="!isCollect" />
              <el-image style="width: 40px; height: 40px" src="/img/collect-fill.png" fit="cover" @click="handleCollect"
                v-else-if="isCollect" />
            </template>
          </el-popover>
          <span class="color-white">
            {{ currentVideo?.author.collection_count+Number(isCollect) }}
          </span>
        </div>
        <div @click="showComments = !showComments" class="my-6! flex flex-col items-center justify-center">
          <el-popover placement="left-start" :width="100" trigger="hover" content="评论"
            popper-class="bg-#33343F! border-none! text-white! icon">
            <template #reference>
              <el-image style="width: 40px; height: 40px" src="/img/message.png" fit="cover" />
              <!-- <el-icon class="color-white!" :size="40">
                <ChatDotRound />
              </el-icon> -->
            </template>
          </el-popover>
          <span class="color-white">
            {{ currentVideo?.comment_count}}
          </span>
        </div>
      </div>
    </div>
    <div v-if="showComments">
      <div class="w-full flex justify-center items-center text-6">评论</div>
      <el-divider class="border-#4C4D4F!" />
      <div>全部评论&nbsp;&nbsp;{{ currentVideo?.comment_count }}</div>
      <el-scrollbar>
        <Comments :videoId="currentVideo?.video_id" @update:comments="handleUpdateComments" />
      </el-scrollbar>
    </div>
  </div>
  <!-- <el-dialog v-model="dialogVisible" width="45%" destroy-on-close center style="background-color:#141414;height:30%;">
    <div class="text-white! text-center text-6 fw-600">请复制视频URL:</div>
    <div class="text-white! text-center">{{ videoItem.play_url }}:</div>
  </el-dialog> -->
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from "element-plus"
import Comments from "@/components/Video/Comments.vue"
import { Video } from '@/types'
import { social as socialApi, video as videoApi } from "@/services"
import Player, { IPlayerOptions } from 'xgplayer';
import 'xgplayer/dist/index.min.css';
import { video } from '@/store/video'
import { user as userInfo } from '@/store/user'

import { storeToRefs } from 'pinia'
const props = defineProps<{
  id: string
}>()



const videoStore = video()
const userStore = userInfo()

const route = useRoute()
const router = useRouter()

const { user_id } = storeToRefs(userStore)
const { currentIndex, video_list, currentVideo } = storeToRefs(videoStore)

const videoItem = ref<Video>()

const player = ref<Player>()
const showComments = ref<boolean>(false)
const videoRef = ref()
const playerConfig = ref<IPlayerOptions>()
const isFavorite = ref(false)
const isCollect = ref(false)
const isFollow = ref(false)



const comments = ref(currentVideo.value.comment_count)
const handleUpdateComments = (val: number) => {
  comments.value = val
}
// 鼠标滚动
const handleWheel = (event) => {
  if (event.deltaY < 0) {
    // 上滚
    if (currentIndex.value === 0) {
      ElMessage.warning("已经是第一个视频了")
      return
    }
    videoStore.setCurrentIndex(--currentIndex.value)
    showComments.value = false
    router.push(`/play?id=${video_list.value[currentIndex.value].video_id}`)
    player?.value?.playNext({
      id: "video",
      el: document.getElementById("video") as HTMLElement,
      width: "100%",
      height: "100%",
      url: [
        {
          // src: videoItem?.value?.play_url,
          src: videoStore.getCurrentVideo().play_url,
          type: "video/mp4"
        },
      ],
      playsinline: true,
      poster: videoStore.getCurrentVideo().cover_url,
      lang: 'zh',
      autoplay: true,
      marginControls: false,
      dynamicBg: {
        disable: false
      }
    })
  } else {
    // 下滚
    if (currentIndex.value === video_list.value.length - 1) {
      ElMessage.warning("已经是最后一个视频了")
      return
    }
    showComments.value = false
    videoStore.setCurrentIndex(++currentIndex.value)
    router.push(`/play?id=${videoStore.getCurrentVideo().video_id}`)
    player?.value?.playNext({
      id: "video",
      el: document.getElementById("video") as HTMLElement,
      width: "100%",
      height: "100%",
      url: [
        {
          // src: videoItem?.value?.play_url,
          src: videoStore.getCurrentVideo().play_url,
          type: "video/mp4"
        },
      ],
      playsinline: true,
      poster: videoStore.getCurrentVideo().cover_url,
      lang: 'zh',
      autoplay: true,
      marginControls: false,
      dynamicBg: {
        disable: false
      }
    })
  }
}
// 操作
const handleFavorite = () => {
  videoApi.toLikeVideo(parseInt(props.id), isFavorite.value ? 2 : 1, currentVideo.value?.partition).then(() => {
    isFavorite.value = !isFavorite.value;
    ElMessage.success(isFavorite.value ? "点赞成功" : "取消点赞成功");
    initVideo()
  })
}

const handleCollect = () => {
  videoApi.toCollectVideo(parseInt(props.id), isCollect.value ? 2 : 1, currentVideo.value?.partition).then(() => {
    isCollect.value = !isCollect.value;
    ElMessage.success(isCollect.value ? "收藏成功" : "取消收藏成功");
    initVideo()
  })
}

const handleFollow = () => {
  socialApi.toFollow(currentVideo.value.author.id, isFollow.value ? 2 :1).then(() => {
    isFollow.value = !isFollow.value
    ElMessage.success(isFollow.value ? "关注成功" : "取消关注成功");
    initVideo()
  })
}


const initVideo = () => {
  videoApi.getVideoInfo(parseInt(route.query.id as string)).then(res => {
    videoItem.value = res.data.video_list[0]
    isFavorite.value = res.data.video_list[0].is_favorite
    isCollect.value = res.data.video_list[0].is_collect
    isFollow.value = res.data.video_list[0].author.is_follow
    let index = 0
    videoStore.setCurrentVideo(videoItem.value as Video)
    video_list.value.forEach((item: any) => {
      if (item.video_id === videoItem.value?.video_id) {
        videoStore.setCurrentIndex(index)
        return
      }
      index++;
    })
    playerConfig.value = {
      id: "video",
      el: document.getElementById("video") as HTMLElement,
      width: "100%",
      height: "100%",
      url: [
        {
          src: currentVideo?.value?.play_url || videoItem?.value?.play_url,
          type: "video/mp4"
        },
      ],
      playsinline: true,
      poster: currentVideo?.value?.cover_url || videoItem?.value?.cover_url,
      lang: 'zh',
      autoplay: true,
      marginControls: false,
      dynamicBg: {
        disable: false
      }
    }
    player.value = new Player(playerConfig.value);
  })
}
onMounted(() => {
  if (route.query.id) {
    initVideo()
  }
})

</script>

<style scoped>
/* .scrolling {
  position: relative;
  animation: scrollUp 4s linear 1;
}

@keyframes scrollUp {
  0% {
    transform: translateY(0);
  }

  100% {
    transform: translateY(-100%);
  }
} */

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
