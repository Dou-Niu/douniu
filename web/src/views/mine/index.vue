<template>
  <el-scrollbar>
    <div class="my_container">
      <div class="header">
        <div>
          <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" :size="110" />
        </div>
        <div class="header-mid">
          <div class="name">{{ "星野露比" }}</div>
          <div class="user-data">
            <div class="user-data-item">关注 <span class="data">{{ 13 }}</span></div>
            <div class="user-data-item">粉丝 <span class="data">{{ 13 }}</span></div>
            <div class="user-data-item">获赞 <span class="data">{{ 444 }}</span></div>
          </div>
          <div class="user-info">
            <div class="user-sign">斗牛号: {{ 213131313 }}</div>
            <!-- <div class="user-year">{{ '男' }}{{ 20 }}岁</div>
            <div class="user-ip">{{ "江苏" }}·{{ "徐州" }}</div>
            <div class="user-others">{{ "中国矿业大学" }}</div> -->
          </div>
          <div class="user-produce">
            {{ "矿大老学长" }}
          </div>

        </div>
        <div class="header-right">
          <div class="btn editor" @click="dialogVisible = true">编辑资料</div>
        </div>
      </div>
      <div class="mid_nav">
        <div class="nav-item" :class="{ 'active': activeName === 'works' }" @click="activeName = 'works'">作品 <span
            class="data">{{ 13 }}</span></div>
        <div class="nav-item" :class="{ 'active': activeName === 'like' }" @click="activeName = 'like'">喜欢 <span
            class="data">{{ 13 }}</span></div>
        <div class="nav-item" :class="{ 'active': activeName === 'collect' }" @click="activeName = 'collect'">收藏 <span
            class="data">{{ 13 }}</span></div>
      </div>
      <div class="line mb-6"></div>
      <div class="footer_list w-80%">
        <div v-if="activeName === 'works'">
          <showVideoList />
        </div>
        <div v-else-if="activeName === 'like'">
          <showVideoList />
        </div>
        <div v-else>
          <showVideoList />
        </div>
      </div>
    </div>
    <el-dialog v-model="dialogVisible" title="编辑资料" width="30%" center align-center destroy-on-close
      style="background-color:rgb(37,38,50);color:red;height:auto;">
      <div class="my-4 text-white flex justify-center flex-col items-center">
        <div class="mb-4 text-4">头像</div>
        <el-upload class="" :show-file-list="false">
          <el-avatar :src="userInfo.avatar" :size="100" />
        </el-upload>
      </div>
      <div class="my-4 flex justify-center flex-col items-center">
        <div class="mb-4 text-4 text-white">背景图片</div>
        <el-upload class="w-full" drag>
          <el-icon :size="50"><upload-filled /></el-icon>
        </el-upload>
      </div>
      <div class="my-4 text-white text-4">
        <span>名字</span>
        <el-input type="text" v-model="userInfo.name" />
      </div>
      <div class="mt-8 text-white text-4">
        <span class="mb-2">简介</span>
        <el-input type="textarea" v-model="userInfo.description" :rows="4" :maxlength="40" />
      </div>
      <template #footer>
        <span>
          <el-button type="primary" size="large" @click="dialogVisible = false" color="rgb(124,101,109)">
            保存
          </el-button>
          <el-button @click="dialogVisible = false" size="large" color="rgb(60,62,73)">取消</el-button>
        </span>
      </template>
    </el-dialog>
  </el-scrollbar>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue"
import showVideoList from "@/components/Video/showVideoList.vue";
const activeName = ref<string>("works")
const dialogVisible = ref<boolean>(false)
// 修改个人信息
const userInfo = reactive({
  name: "",
  description: "",
  avatar: "",
  background: ""
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
