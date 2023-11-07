<template>
  <el-scrollbar>
    <video-list :video-list="videoList"></video-list>
    <div ref="bottomRef" class="bottom-observer">{{ finished ? '已经没有了哦' : '正在加载喵~' }}</div>
  </el-scrollbar>
</template>

<script setup lang="ts">
import VideoList from '@/components/Video/VideoList.vue';
import { ref, onMounted } from 'vue';
import { Video } from '../../types/index';
const props = defineProps<{
  videoList: Video[]
}>()

let finished = ref(false);
let bottomRef = ref<HTMLElement>();
// let videoList = ref<Video[]>([]);
let limit = 25;


// let handleLoad = async () => {

// }

let observer = new IntersectionObserver(entries => {
  entries.forEach(entry => {
    if (entry.isIntersecting) {
      // handleLoad()
    }
  })
})

let observeBottom = () => {
  let bottomDom = bottomRef.value as HTMLElement;
  observer.observe(bottomDom);
}

onMounted(() => {
  if(props.videoList.length < limit) {
    finished.value = true;
  }
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
