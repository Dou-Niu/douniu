<template>
    <div class="upload_input">
        <el-upload drag v-if="status === 'beforeUpload'" :before-upload="handleUpload">
            <el-icon class="el-icon--upload"><upload-filled /></el-icon>
            <div class="el-upload__text" :style="{ color: 'rgba(255,255,255,.9)' }">
                拖拽视频到此处或者点击上传视频
            </div>
        </el-upload>
        <div class="progress" v-else-if="status === 'uploading'">
            <el-progress :width="300" type="dashboard" color="rgba(0,0,0,.5)" :percentage="uploadPercentage">
                <template #default="{ percentage }">
                    <span class="percentage-value">{{ percentage }}%</span>
                    <span class="percentage-label">正在上传</span>
                </template>
            </el-progress>
        </div>
        <video controls v-else-if="status === 'uploaded'" class="pre-video" :src="videoSrc"></video>
    </div>
</template>

<script setup lang="ts">
import { uploadVideo, uploadPercentage } from '@/utils/upload';
import { ElMessage } from 'element-plus';
import { ref, defineExpose } from 'vue';

type Status = 'beforeUpload' | 'uploaded' | 'uploading';
let status = ref<Status>('beforeUpload');
let videoSrc = ref('');

const handleUpload = async (file: File) => {
    try {
        if (!file.type.includes('video')) {
            ElMessage.error("请选择视频文件");
            return false;
        }
        status.value = 'uploading';
        const { videoURL } = await uploadVideo(file);
        status.value = 'uploaded';
        videoSrc.value = videoURL;
    } catch (error: any) {
        ElMessage.error(error);
    }
    return false
}

defineExpose({
    handleUpload
})
</script>



<style lang="less" scoped>
.upload_input {
    width: 100%;
    height: 50vh;
    max-width: 1208px;
    position: relative;

    .percentage-value {
        display: block;
        margin-top: 10px;
        font-size: 28px;
    }

    .percentage-label {
        display: block;
        margin-top: 10px;
        font-size: 12px;
    }

    .progress {
        width: 300px;
        height: 300px;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }

    .pre-video {
        width: 100%;
        height: 100%;
        object-fit: contain;
    }

    :deep(.el-upload) {
        width: 100%;
        height: 50vh;
    }

    :deep(.el-upload-dragger) {
        background-color: rgba(255, 255, 255, .2);
        height: 100%;
        width: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        border-color: rgba(255, 255, 255, .5);

        &:hover {
            border-color: rgba(255, 255, 255, .9);
        }
    }
}
</style>