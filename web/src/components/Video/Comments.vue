<template>
  <div class="w-full h-full overflow-auto! px-2 min-w-80">
    <div class="my-4 text-white text-4 flex flex-1">
        <el-input type="textarea" v-model="myComment" placeholder="请说点什么吧~" />
        <el-button type="primary" class="ml-2" @click="handleSendComment">发送</el-button>
      </div>
    <div class="h-80%">
      <div v-for="comment in commentsList" class="flex my-10">
        <el-avatar :size="50" :src="1" class="mr-2  " />
        <div class="flex flex-col">
          <span class="text-#4C4D4F mb-3">{{ comment.user.nickname || '昵称' }}</span>
          <span class="mb-3">{{ comment.content }}</span>
          <div class="flex flex-col justify-center items-center mb-3">
            <!-- <div class="text-#4C4D4F mx-4">评论</div> -->
          </div>
          <div class="flex items-center justify-center">
            <el-divider class="w-15! border-#4C4D4F!" />
            <div class="text-#4C4D4F! mx-4 flex whitespace-nowrap">
              <div>
                <span>展开回复</span>
                <el-icon>
                  <ArrowDown />
                </el-icon>
              </div>
              <div>
                <span>收起回复</span>
                <el-icon>
                  <ArrowUp />
                </el-icon>
              </div>
            </div>
            <el-divider class="w-15! border-#4C4D4F!" />
          </div>
          <div>
            回复列表
          </div>
        </div>
      </div>

    </div>

  </div>
</template>

<script setup lang="ts">
interface IComment {
  id: bigint,
  content: string,
  create_time: string,
  user: User
}

import { ref, onMounted } from "vue"
import { comment as commentApi } from "@/services";
import { video } from "@/store/video"
import { storeToRefs } from 'pinia'
import { User } from "@/types/user";
import { useRoute } from 'vue-router'
defineProps<{
  video_id: bigint
}>()
const route = useRoute()

const videoStore = video()
const { currentVideo } = storeToRefs(videoStore)


const commentsList = ref<IComment[]>([])
const replyList = ref<IComment[]>([])
const myComment = ref("")

// 发送评论
const handleSendComment = () => {
  console.log("哈哈哈");
  commentApi.sendComment(currentVideo.value.video_id, 1, 0, myComment.value).then(res => {
    console.log(res);
    getComments()
  })
}
const getComments = () => {
  commentApi.getVideoComment(BigInt(route.query.id), BigInt(0)).then(res => {
    commentsList.value = res.data.comment_list
  })
}
onMounted(() => {
  getComments()
})
</script>

<style scoped>
:deep(.el-textarea__inner) {
  --el-input-border-color: none;
  background-color: rgb(51, 52, 63);
  color: white;
}
</style>
