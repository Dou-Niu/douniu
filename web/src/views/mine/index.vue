<template>
  <el-scrollbar>
    <div class="my_container">
      <div class="header">
        <div class="header-left">
          <img
            src="https://p6-pc-sign.douyinpic.com/aweme-avatar/tos-cn-i-0813_ogr9sAA3yCAsxYjwCOfhgtzlAWAeCsAIEotsNH~tplv-8yspqt5zfm-300x300.webp?x-expires=1698847200&x-signature=HgAaiu%2FtFCV3nLnBe%2Bykwe9mT1Y%3D&from=2480802190"
            alt="avatar">
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
          <div class="btn editor">编辑资料</div>
        </div>
      </div>
      <div class="mid_nav">
        <div class="nav-item active">作品 <span class="data">{{ 13 }}</span></div>
        <div class="nav-item">喜欢 <span class="data">{{ 13 }}</span></div>
        <div class="nav-item">收藏 <span class="data">{{ 13 }}</span></div>
      </div>
      <div class="line"></div>
      <div class="footer_list">
        <!-- 视频列表，暂未开发 -->
        <!-- <video-list></video-list> -->
      </div>
    </div>
  </el-scrollbar>
</template>

<script setup lang="ts">
</script>

<style scoped lang="less">
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

    .header-left {
      width: 112px;
      height: 112px;
      border-radius: 100%;
      overflow: hidden;

      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
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
