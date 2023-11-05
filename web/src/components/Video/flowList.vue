<template>
    <div class="flow-container" ref="flowRef" :style="{ height: `${listInstance.maxHeight.value}px` }">
        <div class="flow-item" v-for="item in  listInstance.curRenderListInfo.value " :key="item"
            :style="{ width: `${width ? width + 'px' : itemWidth + '%'}`, transform: `translate(${item.left}px, ${item.top}px)` }">
            <slot :item="item" :listInstance="listInstance" name="item"></slot>
            <!-- 不使用插槽时规定list中默认为图片地址 -->
            <template v-if="!$slots.item">
                <img :src="item">
            </template>
        </div>
    </div>
</template>

<script setup lang="ts">
import { defineProps, onMounted, ref, nextTick, watch } from 'vue';
import { flowList } from '@/utils/flowList'
import { debounce } from '../../utils/debounce';


const props = defineProps<{
    list: any[],
    // 单项宽度
    // 默认使用百分比吧
    itemWidth: number
}>()

const emits = defineEmits(['updateItemWidth'])
let listInstance = new flowList(props.list, props.itemWidth);
let flowRef = ref<HTMLElement>();
let width = ref(0);

let getItemWidth = (containerWidth: number) => {
    let tempWidth = 0;
    if (containerWidth > 1600) {
        // 6列
        emits('updateItemWidth', 16);
        tempWidth = 16;
    }
    if (containerWidth > 1250 && containerWidth <= 1600) {
        // 5列
        emits('updateItemWidth', 19);
        tempWidth = 19;
    } else if (containerWidth > 900 && containerWidth <= 1250) {
        // 4列
        emits('updateItemWidth', 24);
        tempWidth = 24;
    } else if (containerWidth > 640 && containerWidth <= 900) {
        // 3列
        emits('updateItemWidth', 32.5);
        tempWidth = 32.5;
    } else if (containerWidth > 375 && containerWidth <= 640) {
        // 2列
        emits('updateItemWidth', 48);
        tempWidth = 48;
    } else if (containerWidth <= 375) {
        // 1列
        emits('updateItemWidth', 100);
        tempWidth = 100;
    }
    return tempWidth;
}

let listenResize = () => {
    let flowEl = flowRef.value as HTMLElement
    let onResize = debounce(() => {
        let containerWidth = flowEl.offsetWidth;
        let tempWidth = props.itemWidth;
        tempWidth = getItemWidth(containerWidth);
        width.value = containerWidth * tempWidth / 100;
        // 这里要让flow-item的高度更新后在更新布局
        nextTick(() => {
            listInstance.updateOnResize(containerWidth, tempWidth / 100);
        })
    }, 200);
    window.addEventListener('resize', onResize);
}

let init = () => {
    let flowEl = flowRef.value as HTMLElement;
    let containerWidth = flowEl.offsetWidth;
    let tempWidth = getItemWidth(containerWidth);
    width.value = containerWidth * tempWidth / 100;
    listInstance.calcSpaceCols(containerWidth, tempWidth / 100);
    listenResize();
}

onMounted(() => {
    init();
})
watch(() => props.list, () => {
    let index = listInstance.curRenderListInfo.value.length;
    listInstance.curRenderListInfo.value = [...listInstance.curRenderListInfo.value, ...props.list.slice(index)];
})
</script>

<style lang="less" scoped>
.flow-container {
    flex-grow: 1;
    position: relative;
    width: 100%;

    .flow-item {
        position: absolute;
        will-change: transform;
    }
}
</style>