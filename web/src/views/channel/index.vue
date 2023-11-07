<template>
  <el-scrollbar>
    <video-list :video-list="videoList" v-if="videoList.length !== 0"></video-list>
    <div ref="bottomRef" class="bottom-observer">{{ finished ? '已经没有了哦' : '正在加载喵~' }}</div>
  </el-scrollbar>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import VideoList from '@/components/Video/VideoList.vue';
import { Video } from '../../types/index';
import { video as videoApi } from '@/services'
import { video as VideoStore } from "@/store/video"
import { useRoute } from 'vue-router'
const route = useRoute()
const videoStore = VideoStore();

let finished = ref(false);
let bottomRef = ref<HTMLElement>();
let videoList = ref<Video[]>([]);
let max_value = ref(0);

let handleLoad = () => {
  videoApi.getPartitionVideo(max_value.value, 2, parseInt(route.params.id as string)).then(res => {
    videoList.value = res.data.video_list
    max_value.value = res.data.next_max_value
    finished.value = res.data.is_final
    videoStore.setVideoList(videoList.value)
  }).catch(() => {
    finished.value = true;
    observer.unobserve(bottomRef.value as HTMLElement)
  })
}

let observer = new IntersectionObserver(entries => {
  entries.forEach(entry => {
    if (entry.isIntersecting) {
      handleLoad()
    }
  })
})

let observeBottom = () => {
  let bottomDom = bottomRef.value as HTMLElement;
  observer.observe(bottomDom);
}
const getVideos = (partition: number) => {
  videoApi.getPartitionVideo(0, 2, partition).then(res => {
    videoList.value = res.data.video_list
    max_value.value = res.data.next_max_value
    videoStore.setVideoList(videoList.value)
  })
}

watch(() => route.params.id, (newVal) => {
  if (route.params.id) {
    videoList.value = []
    max_value.value = 0
    finished.value = false
    videoStore.setVideoList([])
    getVideos(parseInt(newVal as string))
  }
})

onMounted(() => {
  observeBottom();
})

</script>

<style scoped lang="less">
.bottom-observer {
  height: 100px;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  will-change: transform;
}
</style>
