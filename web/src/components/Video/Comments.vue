<template>
  <div class="w-full h-full overflow-auto!  px-2 min-w-80">
    <div v-for="comment in commentsList" class="flex my-10 ">
      <el-avatar :size="50" :src="1" class="mr-2  " />
      <div class="flex flex-col">
        <span class="text-#4C4D4F mb-3">{{ comment.user.name }}</span>
        <span class="mb-3">{{ comment.content }}</span>
        <div class="flex justify-center items-center mb-3">
          <div class="text-#4C4D4F mx-4">回复</div>
          <div class="text-#4C4D4F mx-4">评论</div>
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
          1212
        </div>
      </div>
    </div>
    <div>

    </div>
    <div class="my-4 text-white text-4 flex">
      <el-input type="text" v-model="myComment" placeholder="请说点什么吧~" />
      <el-button type="primary" class="ml-2" @click="handleSendComment">发送</el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
interface IComment {
  id: bigint,
  content: string,
  create_time: string,
  user: {
    id: bigint
    name: string
    follow_count: number
    follower_count: number
    is_follow: boolean
    avatar: string
    background_image: string
    signature: string
    total_favorited: string
    work_count: number
    favorite_count: number
  }
}
import { ref, onMounted } from "vue"
import { comment as commentApi } from "@/services";
const props = defineProps<{
  video_id: bigint
}>()

const commentsList = ref<IComment[]>([])
const replyList = ref<IComment[]>([])
const myComment = ref("")

// 发送评论
const handleSendComment = () => {
  commentApi.sendComment(props.video_id, 1, BigInt(0), myComment.value).then(res=>{
    console.log(res);
  })
}
onMounted(() => {
  // getComments()
})
</script>

<style scoped>
:deep(.el-input__wrapper) {
  --el-input-border-color: none;
  background-color: rgb(51, 52, 63);
  color: white;
}
</style>
