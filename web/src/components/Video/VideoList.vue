<template>
    <flow-list :item-width="itemWidth" @updateItemWidth="handleUpdateItemWidth" :list="videoList">
        <template #item="{ item, listInstance }">
            <video-item :video-item="item" :list-instance="listInstance"></video-item>
        </template>
    </flow-list>
</template>

<script setup lang="ts">
import VideoItem from './VideoItem.vue';
import flowList from './flowList.vue';
import { Video } from '../../types/index';
import { ref } from 'vue'
defineProps<{
    videoList: Video[]
}>()

let itemWidth = ref(19);
let handleUpdateItemWidth = (newWidth: number) => {
    itemWidth.value = newWidth;
}

</script>