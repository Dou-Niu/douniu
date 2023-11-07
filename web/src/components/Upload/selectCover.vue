<template>
    <div class="upload_cover">
        <div class="title">视频封面</div>
        <div class="cover">
            <div class="set_cover" @click="hanldeClickSelect">
                <img class="fadeInLeft" v-if="curCover" :src="curCover" alt="">
                <div class="text" v-else>
                    上传视频，选择封面
                </div>
                <input ref="inputFileRef" type="file" class="inputFile" @change="handleSelectFile">
            </div>
            <div class="select_cover" v-if="videoCovers.length">
                <div class="select_cover_item fadeIn" v-for="item in videoCovers" @click="handleSelectCover(item)">
                    <img :src="item.url">
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { videoCovers, curCover, curCoverFile } from "@/utils/upload";
import { ref } from 'vue'
import { ElMessage } from 'element-plus';
const inputFileRef = ref<HTMLInputElement>()
const handleSelectCover = async (item: { url: string, blob: Blob }) => {
    curCover.value = item.url;
    curCoverFile.value = await (async () => {
        return new window.File([item.blob], item.url, { type: 'image/webp' })
    })()
}
const hanldeClickSelect = () => {
    if (!videoCovers.value.length) return;
    let inputFileDom = inputFileRef.value as HTMLInputElement;
    inputFileDom.click();
}
const handleSelectFile = (e: Event) => {
    let target = e.target as HTMLInputElement;
    let file = (target.files as FileList)[0];
    if (!file.type.includes('image')) {
        ElMessage.error('请选择图片文件')
        return;
    }
    curCover.value = URL.createObjectURL(file);
    curCoverFile.value = file;
}

</script>

<style lang="less" scoped>
@keyframes fadeIn {
    0% {
        transform: translateY(-30%);
        opacity: 0;
    }

    100% {
        transform: translateY(0);
        opacity: 1;
    }
}

@keyframes fadeInLeft {
    0% {
        transform: translateX(-30%);
        opacity: 0;
    }

    100% {
        transform: translateX(0);
        opacity: 1;
    }
}

.inputFile {
    transform: scale(0);
}

.fadeIn {
    animation: fadeIn;
    animation-duration: 1s;
}

.fadeInLeft {
    animation: fadeInLeft;
    animation-duration: 1s;
}

.title {
    margin: 0 0 4px;
    font-size: 14px;
    line-height: 20px;
    font-weight: bold;
    color: rgba(255, 255, 255, .9);
    display: flex;
    align-items: center;
    line-height: 20px;
}

.upload_cover {
    width: 100%;
    max-width: 1208px;
    display: flex;
    flex-direction: column;

    .cover {
        flex: 1;
        display: flex;
        gap: 20px;

        .set_cover {
            min-height: 50px;
            border-radius: 10px;
            overflow: hidden;
            width: 30%;
            max-height: 230px;

            .text {
                box-sizing: border-box;
                width: 100%;
                height: 100%;
                font-size: 14px;
                display: flex;
                align-items: center;
                color: rgba(255, 255, 255, .5);
                padding-left: 10px;
                background-color: rgba(255, 255, 255, .2);
            }

            img {
                width: 100%;
                height: 100%;
                object-fit: cover;
            }
        }

        .select_cover {
            position: relative;
            flex: 1;
            background-color: rgba(255, 255, 255, .2);
            display: flex;
            gap: 10px;
            padding: 30px 10px 20px 10px;
            border-radius: 10px;

            &::after {
                content: "推荐封面";
                position: absolute;
                top: 5px;
                left: 10px;
                font-weight: 100;
                font-size: 14px;
                color: rgba(255, 255, 255, .5);
            }

            .select_cover_item {
                flex: 1;
                cursor: pointer;
                max-height: 230px;

                img {
                    border-radius: 10px;
                    max-width: 255px;
                    width: 100%;
                    height: 100%;
                    object-fit: cover;
                }
            }
        }
    }
}
</style>