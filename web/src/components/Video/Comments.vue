<template>
  <div class="w-full overflow-auto! comments px-4">
    <div v-for="comment in 10" class="flex my-10">
      <el-avatar :size="50" :src="1" class="mr-2  " />
      <div class="flex flex-col">
        <span class="text-#4C4D4F mb-3">1</span>
        <span class="mb-3">1</span>
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
          回复列表
        </div>
      </div>
    </div>
    <div class="my-4 text-white text-4">
        <span>名字</span>
        <el-input type="text" v-model="myComment" />
      </div>
  </div>
</template>

<script setup lang="ts">
interface IComment {
  id: number,
  content: string,
  create_time: string,
  user: {
    id: number
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
import { ref } from "vue"
const commentsList = ref<IComment[]>([])
const replyList = ref<IComment[]>([])
const myComment = ref("")
</script>

<style scoped></style>
