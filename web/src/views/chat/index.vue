<template>
    <div class="main">
        <el-scrollbar class="friend">
            <div v-for="(item, index) in friendList" class="item" :class="{ click: hoverIndex == index }"
                @click="handleChangeChatingFriend(index)">
                <el-avatar :src="item.avatar" :size="32" />
                <div style="width: 70%">
                    <div style="font-size: 18px">{{ item.name }}</div>
                    <div class="message">{{ item.message }}</div>
                </div>
            </div>
        </el-scrollbar>
        <div class="chat">
            <div class="header">
                <div class="friendName">{{ friendList[hoverIndex].name }}</div>
            </div>
            <el-scrollbar class="message">
                <div v-for="item in friendList[hoverIndex].messages" class="const"
                    :class="{ right: item.name === '星野露比', left: item.name !== '星野露比' }">
                    <el-avatar :src="item.avatar" class="avatar"></el-avatar>
                    <div class="content">
                        {{ item.content }}
                    </div>
                </div>
            </el-scrollbar>
            <el-input v-model="content" type="textarea" placeholder="我来讲两句." :maxlength="100" resize="none" />
            <el-button type="primary" plain class="send" @click="send">
                发送
            </el-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick } from 'vue';
import { ElMessage } from 'element-plus';

let hoverIndex = ref(0);
let content = ref('')

let friendList = ref([
    {
        avatar: 'https://www.kecat.top/other/avatar.webp',
        name: 'kecat1',
        message: '你好啊',
        messages: [
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat1',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat1',
                content: '我叫kecat1'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '我叫星野露比'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat1',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat1',
                content: '我叫kecat1'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '我叫星野露比'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat1',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat1',
                content: '我叫kecat1'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '我叫星野露比'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat1',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat1',
                content: '我叫kecat1'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '我叫星野露比'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat1',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat1',
                content: '我叫kecat1'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '我叫星野露比'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat1',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat1',
                content: '我叫kecat1'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '我叫星野露比'
            },
        ]
    },
    {
        avatar: 'https://www.kecat.top/other/avatar.webp',
        name: 'kecat2',
        message: '你好啊',
        messages: [
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat2',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat2',
                content: '我叫kecat2'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '我叫星野露比'
            }
        ]
    },
    {
        avatar: 'https://www.kecat.top/other/avatar.webp',
        name: 'kecat3',
        message: '你好啊',
        messages: [
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat3',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat3',
                content: '我叫kecat3'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '我叫星野露比'
            }
        ]
    },
    {
        avatar: 'https://www.kecat.top/other/avatar.webp',
        name: 'kecat4',
        message: '你好啊',
        messages: [
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat4',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/other/avatar.webp',
                name: 'kecat4',
                content: '我叫kecat4'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '你好啊'
            },
            {
                avatar: 'https://www.kecat.top/avatar.webp',
                name: '星野露比',
                content: '我叫星野露比'
            }
        ]
    }
])


const handleScrollToBottom = () => {
    let chatBox = document.querySelector('.chat');
    let scrollContainer = chatBox?.children[1].children[0] as HTMLElement;
    let scrollHeight = scrollContainer?.scrollHeight;
    let containerHeight = scrollContainer?.offsetHeight;
    let scrollTop = scrollHeight - containerHeight;
    scrollContainer.scrollTo(0, scrollTop);
}
const handleChangeChatingFriend = (index: number) => {
    if (hoverIndex.value === index) return;
    hoverIndex.value = index
    nextTick(() => {
        handleScrollToBottom();
    })
}
const send = () => {
    if (content.value === '') {
        ElMessage('发了个寂寞')
        return;
    }
    // 模拟
    try {
        friendList.value[hoverIndex.value].messages.push({
            avatar: 'https://www.kecat.top/avatar.webp',
            name: '星野露比',
            content: content.value
        })
        nextTick(() => {
            handleScrollToBottom();
        })
    } catch (error: any) {
        ElMessage.error(error.message)
    } finally {
        content.value = '';
    }
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
})


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
  