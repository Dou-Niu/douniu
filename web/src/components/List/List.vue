<template>
    <div class="list-item" v-for="item in list" :key="item.id">
        <Item></Item>
    </div>
</template>

<script setup lang="ts">
import item from './Item.vue'
const props = defineProps<{
    list: any[]
}>()

</script>

<style lang="less" scoped></style>