<template>
    <div class="main">
        <el-scrollbar class="friend" v-if="friendList">
            <div v-for="(item, index) in friendList" class="item" :class="{ click: hoverIndex == index }"
                @click="() => { switchItem(item, index) }">
                <el-avatar :src="item.avatar" :size="32" v-if="item.id != user_id" />
                <div style="width: 70%" v-if="item.id != user_id">
                    <div style="font-size: 18px">{{ item.nickname }}</div>
                    <div class="message">{{ item.message }}</div>
                </div>
            </div>
        </el-scrollbar>
        <div class="chat" v-if="friendList">
            <!-- <div class="header">
                <div class="friendName">{{ friendList[hoverIndex].nickname }}</div>
            </div> -->
            <el-scrollbar class="message" ref="scrollRef">
                <div v-for="item in messages" class="const"
                    :class="item.from_user_id == userStore?.user_id ? 'right' : 'left'">
                    <el-avatar :src="user_info.avatar" class="avatar" v-if="item.from_user_id === user_info.id"></el-avatar>
                    <el-avatar :src="user.avatar" class="avatar" v-if="item.from_user_id === user.id"></el-avatar>
                    <div class="content">
                        {{ item.content }}
                    </div>
                </div>
            </el-scrollbar>
            <el-input v-model="content" type="textarea" placeholder="输入你想对他说的话" :maxlength="100" resize="none" />
            <el-button type="primary" plain class="send" @click="send">
                发送
            </el-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import { ElMessage, ElScrollbar } from 'element-plus';
import { storeToRefs } from 'pinia'
import { social as socialApi, chat as chatApi } from '@/services';
import { user as userInfo } from '@/store/user'
import { User } from '@/types/user';
const userStore = userInfo()
const { user_id, user_info } = storeToRefs(userStore)
let hoverIndex = ref(0);
let content = ref('')

let friendList = ref<any[]>([])
let messages = ref<any[]>([]);
let pre_msg_time = 0;
let timeOut;
let user = ref<User>();
let scrollRef = ref<InstanceType<typeof ElScrollbar>>()
function switchItem(u: User, index: number) {
    user.value = u;
    hoverIndex.value = index;
    chatApi.getMessage(u.id as number).then((res) => {
        if (res.code == 0) {
            messages.value = res.data.message_list;
            if (res.data.message_list.length > 0) {
                pre_msg_time =
                    res.data.message_list[res.data.message_list.length - 1]?.create_time;
            }
            clearInterval(timeOut);
            timeOut = setInterval(() => {
                chatApi.getMessage(u.id as number, pre_msg_time).then((res) => {
                    if (res.code == 0) {
                        if (res.data.message_list.length > 0) {
                            pre_msg_time =
                                res.data.message_list[res.data.message_list.length - 1].create_time;
                        }
                        messages.value.push(...res.data.message_list);
                    }
                });
                handleScrollToBottom();
            }, 1000);
        }
    });
}

const send = () => {
    if (content.value === '') {
        ElMessage('发了个寂寞')
        return;
    }
    chatApi.sendMessage(user.value?.id as number, 1, content.value)
        .then((res) => {
            if (res.code == 0) {
                content.value = "";
                ElMessage.success("发送成功");
                handleScrollToBottom();
                // messages.value.push({
                //   id: "-1",
                //   content: content.value,
                //   from_user_id: info?.value.id as string,
                //   to_user_id: userId.value,
                //   create_time: Date.now(),
                // });
            }
        });
}
// 滚动到底部
const handleScrollToBottom = () => {
    let chatBox = document.querySelector('.chat');
    let scrollContainer = chatBox?.children[1].children[0] as HTMLElement;
    let scrollHeight = scrollContainer?.scrollHeight;
    let containerHeight = scrollContainer?.offsetHeight;
    let scrollTop = containerHeight - scrollHeight;
    console.log(scrollHeight, containerHeight);
    scrollContainer.scrollTo(0, scrollTop);
}

// // 轮询接口函数
// let timer: any = null;
// const getMessage = () => {
//     timer = setInterval(() => {
//         // 轮询接口
//         // 有变化
//         // 更新数据
//         nextTick(() => {
//             handleScrollToBottom();
//         })
//     }, 100)
// }

onMounted(() => {
    handleScrollToBottom();
    if (user_id?.value) {
        socialApi.getFriendsList(parseInt(user_id.value)).then((res) => {
            if (res.code == 0) {
                friendList.value = res.data.user_list;
            }
        });
    } else {
        ElMessage.error("请先登录");
        // if (loginDialog) {
        //   loginDialog.value = true;
        // }
    }
})

onUnmounted(() => {
    clearInterval(timeOut);
});
</script>
  
<style lang="scss" scoped>
:deep(.el-textarea__inner) {
    color: rgba(255, 255, 255, .9);
    background-color: rgba(255, 255, 255, .2);
    font-size: 15px;
    box-shadow: 0 0 0 1px rgba(255, 255, 255, .5) inset;

    &:focus {
        box-shadow: 0 0 0 1px rgba(255, 255, 255, .9) inset;
    }
}

.main {
    overflow: hidden;
    background-color: rgb(37, 38, 50);
    border-radius: 15px;
    overflow: hidden;
}

.item {
    display: flex;
    justify-content: center;
    padding: 5px;
    font-size: 13px;
    width: 100%;
    box-sizing: border-box;
    justify-content: space-around;
    align-items: center;
    cursor: pointer;
    border-bottom: 1px solid rgb(59, 60, 71);
    transition: background .5s;

    &:hover {
        background: linear-gradient(90deg, #252632, #33343f 25.44%, #33343f 75.56%, #252632);
    }

    &.click {
        background: linear-gradient(90deg, #252632, #33343f 25.44%, #33343f 75.56%, #252632);
    }

    .message {
        display: inline-block;
        white-space: nowrap;
        width: 100%;
        overflow: hidden;
        text-overflow: ellipsis;
    }
}

.chat {
    // width: max-content;
    height: 100%;
    position: relative;
    flex: 1;
    padding-top: 75px;
    box-sizing: border-box;

    .header {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 75px;
        border-bottom: 1px solid rgb(59, 60, 71);
        display: flex;
        align-items: center;
        padding-left: 20px;
        font-size: 18px;
    }

    .message {
        height: 79%;
        box-sizing: border-box;

        .const {
            width: 100%;
            margin: 10px 0;
            padding: 20px;
            min-height: 30px;
            box-sizing: border-box;

            &.left {
                display: flow-root;
                position: relative;

                & .avatar {
                    float: left;
                    width: 40px;
                    height: 40px;
                }

                & .content {
                    background-color: rgb(65, 66, 76);
                    color: rgba(255, 255, 255, .9);
                    border-radius: 10px;
                    position: absolute;
                    left: 70px;
                    padding: 10px 15px;
                }
            }

            &.right {
                display: flow-root;
                position: relative;

                & .avatar {
                    float: right;
                    width: 40px;
                    height: 40px;
                }

                & .content {
                    color: rgba(255, 255, 255, .9);
                    background-color: rgb(41, 140, 255);
                    border-radius: 10px;
                    position: absolute;
                    right: 70px;
                    padding: 10px 15px;
                }
            }
        }
    }

    .el-textarea {
        height: 20%;
        box-sizing: border-box;

        :deep(textarea) {
            height: 100%;
        }
    }

    .send {
        position: absolute;
        bottom: 10px;
        right: 10px;
    }
}

.friend {
    width: 20%;
    max-width: 400px;
    height: 100%;
    border-right: 1px solid rgb(59, 60, 71);
}

.main {
    margin: -10px;
    width: calc(100% + 10px);
    height: calc(100% + 10px);
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>
  