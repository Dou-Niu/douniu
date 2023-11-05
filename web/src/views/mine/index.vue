<template>
  <el-scrollbar>
    <div class="my_container">
      <div class="header">
        <div>
          <el-avatar :src="user.avatar" :size="110" />
        </div>
        <div class="header-mid">
          <div class="name">{{ user.nickname }}</div>
          <div class="user-data">
            <div class="user-data-item">关注 <span class="data">{{ user.follow_count }}</span></div>
            <div class="user-data-item">粉丝 <span class="data">{{ user.follower_count }}</span></div>
            <div class="user-data-item">获赞 <span class="data">{{ user.total_favorited }}</span></div>
          </div>
          <div class="user-info">
            <div class="user-sign">斗牛号: {{ user.id }}</div>
            <!-- <div class="user-year">{{ '男' }}{{ 20 }}岁</div>
            <div class="user-ip">{{ "江苏" }}·{{ "徐州" }}</div>
            <div class="user-others">{{ "中国矿业大学" }}</div> -->
          </div>
          <div class="user-produce">
            {{ user.signature }}
          </div>

        </div>
        <div class="header-right">
          <div class="btn editor" @click="dialogVisible = true">编辑资料</div>
        </div>
      </div>
      <div class="mid_nav">
        <div class="nav-item" :class="{ 'active': activeName === 'works' }" @click="activeName = 'works'">作品 <span
            class="data">{{ user.work_count }}</span></div>
        <div class="nav-item" :class="{ 'active': activeName === 'favorite' }" @click="activeName = 'favorite'">喜欢 <span
            class="data">{{ user.favorite_count }}</span></div>
        <div class="nav-item" :class="{ 'active': activeName === 'collect' }" @click="activeName = 'collect'">收藏 <span
            class="data">{{ user.collection_count }}</span></div>
      </div>
      <div class="line mb-6"></div>
      <div class="footer_list w-80%">
        <div v-if="activeName === 'works'">
          <showVideoList :videoList="videoList" />
        </div>
        <div v-else-if="activeName === 'like'">
          <showVideoList :videoList="videoList" />
        </div>
        <div v-else>
          <showVideoList :videoList="videoList" />
        </div>
      </div>
    </div>
    <el-dialog v-model="dialogVisible" title="编辑资料" width="30%" center align-center destroy-on-close
      style="background-color:rgb(37,38,50);color:red;height:auto;">
      <!-- <div class="my-4 text-white flex justify-center flex-col items-center">
        <div class="mb-4 text-4">头像</div>
        <el-upload class="" :show-file-list="false" :on-success="handleAvatarSuccess">
          <el-avatar :src="user.avatar" :size="100" />
        </el-upload>
      </div> -->
      <!-- <div class="my-4 flex justify-center flex-col items-center">
        <div class="mb-4 text-4 text-white">背景图片</div>
        <el-upload class="w-full" drag :on-success="handleAvatarSuccess" action="">
          <el-icon :size="50"><upload-filled /></el-icon>
        </el-upload>
      </div> -->
      <div class="my-4 text-white text-4">
        <span>名字</span>
        <el-input type="text" v-model="formUserInfo.nickname" />
      </div>
      <div class="mt-8 text-white text-4">
        <span class="mb-2">简介</span>
        <el-input type="textarea" v-model="formUserInfo.signature" :rows="4" :maxlength="40" />
      </div>
      <template #footer>
        <span>
          <el-button type="primary" size="large" @click="handleSubmit" color="rgb(124,101,109)">
            保存
          </el-button>
          <el-button @click="dialogVisible = false" size="large" color="rgb(60,62,73)">取消</el-button>
        </span>
      </template>
    </el-dialog>
  </el-scrollbar>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, watch } from "vue"
import type { UploadProps } from 'element-plus'
import { storeToRefs } from 'pinia'
import { type User } from "@/services/user";
import showVideoList from "@/components/Video/showVideoList.vue";
import { Video } from '@/types'
import { user as userInfo } from '@/store/user';
import { user as userApi, video as videoApi } from "@/services"

const userStore = userInfo();
const { user_info } = storeToRefs(userStore);

const activeName = ref<string>("works")
const dialogVisible = ref<boolean>(false)
const videoList = ref<Video[]>([])
//个人信息
const user = reactive<User>({
  id: BigInt(0),
  phone: '',
  nickname: '',
  follow_count: 0,
  follower_count: 0,
  is_follow: false,
  avatar: 'https://www.kecat.top/avatar.webp',
  background_image: '',
  signature: '',
  total_favorited: '',
  work_count: 0,
  favorite_count: 0,
  collection_count: 0
})

// 表单修改
const formUserInfo = reactive({
  nickname: '',
  signature: '',
  avatar: '',
  background_image: '',
})

// 保存信息
const handleSubmit = () => {
  // 修改类型 1:昵称 2:个性签名 3:头像 4:背景图
  if (formUserInfo.nickname) {
    userApi.changeUserInfo(
      1,
      formUserInfo.nickname
    ).then(() => {
      initInfo()
    })
  }
  if (formUserInfo.signature) {
    userApi.changeUserInfo(
      2,
      formUserInfo.signature
    ).then(() => {
      initInfo()
    })
  }
  if (formUserInfo.avatar) {
    userApi.changeUserInfo(
      3,
      formUserInfo.avatar
    ).then(() => {
      initInfo()
    })
  }
  if (formUserInfo.background_image) {
    userApi.changeUserInfo(
      4,
      formUserInfo.background_image
    ).then(() => {
      initInfo()
    })
  }
  dialogVisible.value = false
}

//获取个人信息
const initInfo = () => {
  user.id = user_info.value.id
  userStore.getUserInfo().then(res=>{
    console.log(res);
    user.id = BigInt(res.id)
    user.nickname = res.nickname
    user.avatar = res.avatar
    user.background_image = res.background_image
    user.nickname = res.nickname
    user.signature = res.signature
    user.avatar = res.background_image
    user.background_image = res.background_image
    user.follow_count = res.follow_count
    user.follower_count = res.follower_count
    user.is_follow = res.is_follow
    user.total_favorited = res.total_favorited
    user.work_count = res.work_count
    user.favorite_count = res.favorite_count
    user.collection_count = res.collection_count
    getVideos()
  })
}

// 上传图片
const handleAvatarSuccess: UploadProps['onSuccess'] = (
  response,
  uploadFile
) => {
  console.log(response, uploadFile.raw);
  formUserInfo.avatar = URL.createObjectURL(uploadFile.raw!)
}


const getVideos = () => {
  if (activeName.value === 'works') {
    videoList.value = []

    videoApi.getUserAllVideo(BigInt(user.id), 9999999999999, 2).then(res => {
      videoList.value = res.data.video_list
    })
  } else if (activeName.value === 'collect') {
    videoList.value = []
    videoApi.getCollectVideoList(BigInt(user.id), 1).then(res => {
      videoList.value = res.data.video_list
    })
  } else if (activeName.value === 'favorite') {
    videoList.value = []
    videoApi.getLikeVideoList(BigInt(user.id), 1).then(res => {
      videoList.value = res.data.video_list
    })
  }
}

watch(activeName, () => {
  getVideos()
})

onMounted(() => {
  initInfo()
})

</script>

<style scoped lang="less">
:deep(.el-textarea__inner) {
  --el-input-border-color: none;
  background-color: rgb(51, 52, 63);
}

:deep(.el-input__wrapper) {
  --el-input-border-color: none;
  background-color: rgb(51, 52, 63);
}

:deep(.el-dialog__title) {
  color: white;
}

.my_container {
  min-width: 540px;
  box-sizing: border-box;
  padding: 20px;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;

  .header {
    max-width: 1208px;
    display: flex;
    width: 100%;
    gap: 20px;

    .data {
      color: rgba(255, 255, 255, .9);
      font-size: 16px;
      line-height: 24px;
    }

    .header-mid {
      flex: 1;
      display: flex;
      flex-direction: column;
      gap: 5px;

      .name {
        font-family: PingFang SC, DFPKingGothicGB-Medium, sans-serif;
        font-size: 20px;
        font-weight: 700;
        line-height: 28px;
        max-width: 300px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }

      .user-data,
      .user-info {
        display: flex;

        .user-sign {
          color: rgba(255, 255, 255, .5);
          font-size: 12px;
          line-height: 20px;
          margin-right: 20px;
        }
      }

      .user-produce {
        font-size: 12px;
        line-height: 20px;
      }

      .user-data {
        .user-data-item {
          position: relative;
          color: rgba(255, 255, 255, .5);
          font-size: 14px;
          line-height: 22px;
          margin-right: 6px;
        }

        .user-data-item+.user-data-item {
          margin-left: 20px;

          &::before {
            content: "";
            width: 1px;
            height: 16px;
            background-color: rgba(89, 89, 90, 0.8);
            position: absolute;
            left: -10px;
            top: 50%;
            transform: translateY(-50%);
          }
        }
      }
    }

    .header-right {
      min-width: 120px;
      position: relative;

      .editor {
        position: absolute;
        right: 20px;
        bottom: 0;
      }

      .btn {
        background: rgba(65, 66, 76, 1);
        color: rgba(255, 255, 255, .9);
        padding: 7px 17px;
        border-radius: 5px;
        font-size: 14px;
        cursor: pointer;
      }
    }
  }

  .line {
    width: 400%;
    height: 1px;
    background-color: gray;
    transform: scale(0.25);
  }

  .mid_nav {
    margin-top: 16px;
    max-width: 1208px;
    width: 100%;
    display: flex;
    padding: 11px 0;
    height: 26px;

    .nav-item {
      line-height: 26px;
      font-size: 18px;
      color: rgba(255, 255, 255, .34);
      cursor: pointer;
    }

    .nav-item+.nav-item {
      margin-left: 40px;
    }

    .active {
      color: rgba(255, 255, 255, .9);
    }
  }
}
</style>
