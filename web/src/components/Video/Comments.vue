<template>
  <div class="w-full h-full overflow-auto! px-2 min-w-80">
    <div class="my-4 text-white text-4 flex flex-1">
      <el-input type="textarea" v-model="myComment" placeholder="请说点什么吧~" />
      <el-button type="primary" class="ml-2" @click="handleSendComment(0)">发送</el-button>
    </div>
    <div class="h-80%">
      <div v-for="comment in commentsList" class="flex my-10">
        <el-avatar :size="50" :src="comment.user.avatar" class="mr-2  " />
        <div class="flex flex-col w-full">
          <span class="text-#4C4D4F mb-3">{{ comment.user.nickname }}</span>
          <span class="mb-3">{{ comment.content }}</span>
          <div class="flex justify-start items-center mb-3">
            <div class="text-#4C4D4F">
              <el-button type="primary" @click="showReplyList(comment.id, comment.user.nickname)" plain>评论</el-button>
            </div>
          </div>
          <!-- <div class="flex items-center justify-center">
            <el-divider class="w-15! border-#4C4D4F!" />
            <div class="text-#4C4D4F! mx-4 flex whitespace-nowrap">
              <div>
                <span @click="showReplyList(comment.id)">展开回复</span>
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
          </div> -->
          <el-divider class="border-#4C4D4F!" />
          <div v-for="replys in replyList">
            <template v-for="rep in replys">
              <div class="flex" v-if="rep?.parent_id == comment.id.toString()">
                <el-avatar :size="50" :src="rep.user.avatar" class="mr-2" />
                <div class="flex flex-col">
                  <span class="text-#4C4D4F mb-3">{{ rep?.user?.nickname }}</span>
                  <span class="mb-3">{{ rep?.content }}</span>
                  <div>
                  </div>
                </div>
              </div>
            </template>
          </div>
        </div>
      </div>
    </div>
    <div class="my-4 text-white text-4 fixed bottom-0 mx-4" v-if="replyVisible">
      <div class="flex justify-center w-full">
        <el-input type="text" v-model="reply" :placeholder="`回复@${replyName}`" class="w-250px!" />
        <el-button type="primary" class="ml-2" @click="handleSendComment(replyId)">发送</el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface IComment {
  id: number,
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
  videoId: number
}>()
const route = useRoute()

const videoStore = video()
const { currentVideo } = storeToRefs(videoStore)

const commentsList = ref<IComment[]>([])
const replyList = ref<any[]>([])
const myComment = ref("")
const reply = ref("")
const replyId = ref(0)
const replyName = ref("")
const replyVisible = ref(false)
const showReplyList = (id: number, name: string) => {
  replyId.value = id
  replyName.value = name
  replyVisible.value = !replyVisible.value
}


// 发送评论
const handleSendComment = (parent_id?: number) => {
  if (parent_id === undefined) parent_id = 0
  if (parent_id !== 0) {
    // 评论回复
    commentApi.sendComment(currentVideo.value.video_id, 1, parent_id, reply.value).then(() => {
      getReply(parent_id as number)
      reply.value = ""
    })
  } else {
    // 评论视频
    commentApi.sendComment(currentVideo.value.video_id, 1, 0, myComment.value).then(() => {
      getComments()
      myComment.value = ""
    })

  }
}
const getComments = () => {
  commentApi.getVideoComment(parseInt(route.query.id as string), 0).then(res => {
    commentsList.value = res.data.comment_list
    for (let comment of commentsList.value) {
      getReply(comment.id)
    }
  })
}

const getReply = (parent_id: number) => {
  commentApi.getComment(parent_id).then(res => {
    replyList.value.push(res.data.comment_list)
    console.log(replyList.value);
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

:deep(.el-input__wrapper) {
  --el-input-border-color: none;
  background-color: rgb(51, 52, 63);
  color: white;
}
</style>
