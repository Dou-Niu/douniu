<template>
  <div class="w-full max-h-full overflow-auto! comments px-4">
    <div class="total">
      全部评论
    </div>
    <el-scrollbar max-height="800px">
      <div v-for="item in 10" class="flex my-10" v-infinite-scroll="load">
        <el-avatar :size="50" :src="circleUrl" class="mr-2  " />
        <div class="flex flex-col">
          <span class="text-#4C4D4F mb-3">名字</span>
          <span class="mb-3">评论</span>
          <div class="flex justify-between items-center mb-3">
            <div class="text-#4C4D4F mx-4">回复</div>
            <div class="text-#4C4D4F mx-4">分享</div>
            <div class="text-#4C4D4F mx-4">评论</div>
          </div>
          <div class="flex items-center justify-center">
            <el-divider class="w-15! border-#4C4D4F!" />
            <div class="text-#4C4D4F! mx-4 flex whitespace-nowrap">
              展开回复<el-icon>
                <ArrowDown />
              </el-icon>
            </div>
            <el-divider class="w-15! border-#4C4D4F!" />
          </div>
          <div>
            回复列表
          </div>
        </div>
      </div>
    </el-scrollbar>
  </div>
</template>

<script setup>

</script>

<style scoped>

</style>
