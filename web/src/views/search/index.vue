<template>
  <el-scrollbar>
    <div>
      全部结果{{ videoList.length }}
    </div>
    <video-list :video-list="videoList"></video-list>
    <div ref="bottomRef" class="bottom-observer">{{ finished ? '已经没有了哦' : '正在加载喵~' }}</div>
  </el-scrollbar>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { useRoute } from 'vue-router';
import VideoList from '@/components/Video/VideoList.vue';
import { Video } from '../../types/index';
import { video as videoApi } from "@/services"
import { video as videoInfo} from "@/store/video"
const videoStore = videoInfo()

const route = useRoute()

let finished = ref(false);
let bottomRef = ref<HTMLElement>();
let videoList = ref<Video[]>([]);
let curPage = ref(1);


let handleLoad = () => {
  videoApi.searchVideo(route.query.key_words as string, curPage.value).then(res => {
    videoList.value = res.data.video_list
    finished.value = res.data.is_final;
    videoStore.setVideoList(videoList.value)
    if (res.data.is_final) {
      observer.unobserve(bottomRef.value as HTMLElement);
      curPage.value = 1
    } else {
      curPage.value++;
    }
  }).catch(() => {
    observer.unobserve(bottomRef.value as HTMLElement);
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

watch(() => route.query.key_words, (newVal) => {
  if (newVal) {
    curPage.value = 1;
    videoList.value = [];
    videoApi.searchVideo(newVal as string, curPage.value).then(res => {
      videoList.value = res.data.video_list
      finished.value = res.data.is_final;
      videoStore.setVideoList(videoList.value)
      curPage.value++;
    })
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
