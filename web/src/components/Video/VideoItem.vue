<template>
    <div ref="videoRef" class="video-item" @click="handlePlay" :style="{ transform: `scale(${scale})` }">
        <div class="video-item-top">
            <el-image @load="handleImgLoad" :src="videoItem.cover_url">
                <template #error>
                    <el-image @load="handleImgLoad" src="https://www.kecat.top/video/logo.jpg
" alt="默认图片" />
                </template>
            </el-image>
            <div class="video-item-likeCount">{{ videoItem.favorite_count }}</div>
            <!-- <div class="video-item-duration">{{ "00:12" }}</div> -->
        </div>
        <div class="video-item-bottom">
            <div class="video-item-title">{{ videoItem.title }}</div>
            <div class="video-item-other">
                <div class="elipsis">{{ videoItem.author.nickname }}</div>
                <div class="video-item-time">{{ formatTime(videoItem.create_time) }}</div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { Video } from '../../types/index';
import { ref } from 'vue';
import { flowList } from '../../utils/flowList/index';
import { formatTime } from '@/utils/format';
import { useRouter } from "vue-router"
const router = useRouter()

let videoRef = ref<HTMLElement>();
let scale = ref(0);
let props = defineProps<{
    videoItem: Video,
    listInstance: flowList
}>()

let handleImgLoad = async () => {
    await props.listInstance.setPosition(props.videoItem, (videoRef.value?.parentElement as HTMLElement));
    scale.value = 1;
}

const handlePlay = () => {
    router.push({
        path: "/play",
        query: {
            id: BigInt(props.videoItem.video_id).toString(),
        }
    })
}

</script>

<style lang="less" scoped>
.elipsis {
    max-width: 50%;
    text-overflow: ellipsis;
    overflow: hidden;
}

.video-item {
    position: relative;
    background-color: rgb(37, 38, 50);
    border-radius: 10px;
    overflow: hidden;
    width: 100%;

    &-top {
        position: relative;
        width: 100%;

        img {
            width: 100%;
            object-fit: contain;
        }

        .video-item-duration {
            position: absolute;
            bottom: 0.35em;
            right: 0.3em;
            font-size: 14px;
            padding: 2px 4px;
            background-color: #000;
            border-radius: 4px;
        }

        .video-item-likeCount {
            position: absolute;
            bottom: 0.35em;
            left: 0.3em;
            text-shadow: 0 0.712644px 0.712644px rgba(0, 0, 0, .15);
            font-size: 14px;
        }
    }

    &-bottom {
        padding: 12px;
        display: flex;
        flex-direction: column;

        .video-item-title {
            cursor: pointer;
            overflow: hidden;
            text-overflow: ellipsis;
            word-break: break-all;
            -webkit-box-orient: vertical;
            -webkit-line-clamp: 2;
            color: rgba(255, 255, 255, .9);
            display: -webkit-box;
            font-size: 15px;
            line-height: 23px;
        }

        .video-item-other {
            display: flex;
            justify-content: space-between;
            color: rgba(22, 24, 35, .34);
            color: rgb(138, 138, 145);
            display: inline-flex;
            font-size: 14px;
            line-height: 22px;
            margin-top: 8px;
            white-space: nowrap;
        }
    }
}
</style>