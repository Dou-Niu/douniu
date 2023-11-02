<template>
    <el-scrollbar>
        <div class="upload_container">
            <div class="top_title">
                <span>发布视频</span>
                <el-button v-if="videoCovers.length" class="btn" type="success" round
                    @click="handleUploadAgain">重新上传</el-button>
                <input type="file" ref="inputRef" @change="handleFileChange" style="transform: scale(0);">
            </div>
            <upload-video ref="uploader"></upload-video>
            <div class="upload_title">
                <div class="title">视频标题</div>
                <el-input :rows="5" placeholder="写一个合适的标题，能让更多人看到" v-model="videoTitle" type="textarea" :max="30"
                    :show-word-limit="true"></el-input>
            </div>
            <select-cover></select-cover>
            <div class="upload_category">
                <div class="title">视频分类</div>
                <el-select></el-select>
            </div>
            <div class="upload_publish">
                <el-button type="success" @click="handlePublish">发布</el-button>
                <el-button>取消</el-button>
            </div>
        </div>
    </el-scrollbar>
</template>

<script setup lang="ts">
import uploadVideo from '@/components/Upload/uploadVideo.vue';
import { ref } from 'vue';
import { curCover, videoCovers, curVideoURL } from '@/utils/upload';
import selectCover from '@/components/Upload/selectCover.vue'
import { uploadVideo as uploadFile, curCoverFile } from '../../utils/upload/uploadVideo';
const videoTitle = ref('');

const uploader = ref();
const inputRef = ref();
const handleUploadAgain = () => {
    (inputRef.value as HTMLInputElement).click();
}
const handleFileChange = (e: Event) => {
    curCover.value = "";
    videoCovers.value.length = 0;
    let target = e.target as HTMLInputElement;
    uploader.value.handleUpload((target.files as FileList)[0])
    target.value = ""
}
const handlePublish = async () => {
    let file = curCoverFile.value as File;
    let videoURL = curVideoURL.value;
    let imgURL = await uploadFile(file);
    console.log('封面', imgURL.videoURL);
    console.log('视频地址', videoURL);
}

</script>

<style scoped lang="less">
.upload_container {
    margin: auto;
    width: 90%;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 20px;

    .top_title {
        position: relative;
        width: 100%;
        font-style: normal;
        font-weight: bold;
        font-size: 24px;
        line-height: 28px;
        color: rgba(255, 255, 255, .9);
        max-width: 1208px;

        .btn {
            position: absolute;
            right: 10px;
            top: 6px;
        }
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

    .upload_title {
        width: 100%;
        max-width: 1208px;

        :deep(.el-input) {
            width: 50%;
        }

        :deep(.el-textarea__inner) {
            color: rgba(255, 255, 255, .9);
            background-color: rgba(255, 255, 255, .2);
            font-size: 15px;
            box-shadow: 0 0 0 1px rgba(255, 255, 255, .5) inset;

            &:focus {
                box-shadow: 0 0 0 1px rgba(255, 255, 255, .9) inset;
            }
        }
    }



    .upload_category {
        width: 100%;
        max-width: 1208px;

        :deep(.el-select .el-input.is-focus .el-input__wrapper) {
            box-shadow: 0 0 0 1px rgba(255, 255, 255, .9) inset !important;
        }

        :deep(.el-input__wrapper) {
            background-color: rgba(255, 255, 255, .2);
            box-shadow: 0 0 0 1px rgba(255, 255, 255, .5) inset !important;
        }
    }

    .upload_publish {
        width: 100%;
        max-width: 1208px;
    }
}
</style>
