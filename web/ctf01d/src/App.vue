<script setup>
import Header from './components/common/header/header.vue'
import Welcome from './components/welcome/welcome.vue'
</script>

<template>
  <div>
    <Header />
    <div class="container">
      <div class="row">
        <div class="col7">
          <Welcome />
        </div>
        <div class="col5">
          <img src="./assets/img/welcome.svg" alt="" class="hero_img">
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
