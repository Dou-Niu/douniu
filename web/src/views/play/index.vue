<template>
    <VideoPlay />
</template>

<script setup>
import {onMounted} from 'vue'
import { useRoute } from 'vue-router'
const route = useRoute()

onMounted(() => {
    console.log(route.query.id);
})
</script>

<style scoped>

</style>
