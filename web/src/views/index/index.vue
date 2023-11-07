<template>
  <el-scrollbar>
    <video-list :video-list="videoList" v-if="videoList.length !== 0"></video-list>
    <div ref="bottomRef" class="bottom-observer">{{ finished ? '已经没有了哦' : '正在加载喵~' }}</div>
  </el-scrollbar>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import VideoList from '@/components/Video/VideoList.vue';
import { Video } from '../../types/index';
import { video as videoApi } from '@/services'
import { video as VideoStore } from "@/store/video"

const videoStore = VideoStore();


let finished = ref(false);
let bottomRef = ref<HTMLElement>();
let videoList = ref<Video[]>([]);
let curTime = ref(0);

let handleLoad = () => {
  videoApi.getHomeVideo(curTime.value).then(res => {
    videoList.value = res.data.video_list
    curTime.value = res.data.next_time
    videoStore.setVideoList(videoList.value)
  }).catch(()=>{
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

onMounted(() => {
  // videoApi.getHomeVideo(0).then(res => {
  //   console.log(res);
  //   videoList.value = res.data.video_list
  //   curTime.value = res.data.next_time
  // })
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
