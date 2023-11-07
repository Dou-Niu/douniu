<template>
  <div class="flex justify-center items-center my-6">
    <el-avatar :size="40" src="/img/logo.jpg" class="mr-4"/>
    <div class="flex justify-center items-center text-2xl loader">斗牛</div>
  </div>
</template>

<script setup>

</script>

<style scoped>
@keyframes animate8345 {

  0%,
  100% {
    filter: hue-rotate(0deg);
  }

  50% {
    filter: hue-rotate(360deg);
  }
}

.loader {
  color: rgb(0, 0, 0);
  background: linear-gradient(to right, #2d60ec, #3ccfda);
  font-size: 30px;
  -webkit-text-fill-color: transparent;
  -webkit-background-clip: text;
  animation: animate8345 9s linear infinite;
  font-weight: bold;
}
</style>
