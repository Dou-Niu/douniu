<template>
    <div class="w-full h-full overflow-hidden flex scrolling">
        <div id="video" ref="videoRef" @wheel="handleWheel">
            <div class="info fw-600">
                <div class="text-6 flex gap-5 items-center">@{{ currentVideo?.author.nickname }}
                    <el-button type="danger" icon="Plus" circle plain />
                </div>
                <div class="text-4">{{ currentVideo?.title }}</div>
            </div>
            <div class="right-side text-black">
                <el-avatar :src="currentVideo?.author.avatar" :size="50" />
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
                        {{ currentVideo?.favorite_count }}
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
                        {{ currentVideo?.collection_count }}
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
                        {{ currentVideo?.comment_count }}
                    </span>
                </div>
                <div class="my-6! flex flex-col items-center justify-center">
                    <el-popover placement="left-start" :width="100" trigger="hover" content="分享"
                        popper-class="bg-#33343F! border-none! text-white! icon">
                        <template #reference>
                            <el-icon class="color-white!" :size="40" @click="handleShare">
                                <BottomRight />
                            </el-icon>
                        </template>
                    </el-popover>
                </div>
            </div>
        </div>
        <div v-if="showComments">
            <div class="w-full flex justify-center items-center text-6">评论</div>
            <el-divider class="border-#4C4D4F!" />
            <div>全部评论&nbsp;&nbsp;{{ currentVideo?.comment_count }}</div>
            <el-scrollbar>
                <Comments :videoId="currentVideo?.video_id" />
            </el-scrollbar>
        </div>
    </div>
    <el-dialog v-model="dialogVisible" width="45%" destroy-on-close center style="background-color:#141414;height:30%;">
        <div class="text-white! text-center text-6 fw-600">请复制视频URL:</div>
        <div class="text-white! text-center">{{ videoItem.play_url }}:</div>
    </el-dialog>
</template>
  
<script setup lang="ts">
import { ref, onMounted } from "vue"
import { useRoute } from "vue-router"
import { ElMessage } from "element-plus"
import Comments from "@/components/Video/Comments.vue"
import { Video } from '@/types'
import { video as videoApi } from "@/services"
import Player, { IPlayerOptions } from 'xgplayer';
import 'xgplayer/dist/index.min.css';
import { video } from '@/store/video'
import { storeToRefs } from 'pinia'
const videoStore = video()
const { currentIndex, video_list, currentVideo } = storeToRefs(videoStore)

const videoItem = ref<Video>()

const route = useRoute()
const player = ref<Player>()
const showComments = ref<boolean>(false)
const videoRef = ref()
const playerConfig = ref<IPlayerOptions>()
const nextMaxHot = ref<bigint>(0)
const url = ref("")
// 对话框
const dialogVisible = ref(false)

// 鼠标滚动
const handleWheel = (event) => {
    if (event.deltaY < 0) {
        // 上滚
        if (currentIndex.value === 0) {
            ElMessage.warning("已经是第一个视频了")
            return
        }
        videoStore.setCurrentIndex(--currentIndex.value)
        console.log(videoStore.getCurrentVideo());
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
            videoApi.getHotVideo(nextMaxHot.value).then(res => {
                nextMaxHot.value = res.data.NextMaxHot
                console.log(res.data);
                videoStore.setVideoList([...videoStore.getVideoList(), ...res.data.video_list])
            })
            return
        } else {
            videoStore.setCurrentIndex(++currentIndex.value)
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
}
// 操作
const handleFavorite = () => {
    videoApi.toLikeVideo(route.query.id, 1, currentVideo.value?.partition).then(res => {

    })
}

const handleCollect = () => {
    videoApi.toCollectVideo(BigInt(route.query.id), 1, currentVideo.value?.partition).then(res => {
        console.log(res);
    })
}

const handleShare = () => {
    videoApi.shareVideo(BigInt(route.query.id)).then(res => {
        dialogVisible.value = true
        url.value = res.data.share_url
    })
}


onMounted(() => {
    videoApi.getHotVideo(0).then(res => {
        nextMaxHot.value = res.data.NextMaxHot
        videoStore.setVideoList(res.data.video_list)
        videoStore.setCurrentIndex(0)
    })
    // videoApi.getVideoInfo(BigInt(route.query.id)).then(res => {
    //     videoItem.value = res.data.video_list[0]
    //     let index = 0
    //     videoStore.setCurrentVideo(videoItem.value)
    //     video_list.value.forEach((item: any) => {
    //         if (item.video_id === videoItem.value?.video_id) {
    //             videoStore.setCurrentIndex(index)
    //             return
    //         }
    //         index++;
    //     })
    // })
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
            // {
            //   src: "https://www.kecat.top/video/jingliu.mp4",
            //   type: "video/mp4"
            // },
            // {
            //   src: "//lf3-static.bytednsdoc.com/obj/eden-cn/nupenuvpxnuvo/xgplayer_doc/xgplayer-demo.mp4",
            //   type: "video/mp4"
            // }
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
  