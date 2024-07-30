<template>
  <header class="header_wrapper">
    <div class="header_logo">
      <RouterLink to="/">
        <img class="logo" src="../../../assets/img/header_logo.svg" alt="">
      </RouterLink>
    </div>
    <Menu />
    <div class="header_profile_info">
      <Button 
        :title="'log in'"
        :isPrimary="true"
        :imgPath=loginButtonIconPath
      />
    </div>
  </header>
</template>

<script setup>
  import Menu from './header_menu.vue'
  import Button from '../button/button.vue'
  const loginButtonIconPath = '/src/assets/icon/log-in.svg'
</script>

<style>
  .header_wrapper {
    display: flex;
    width: 100%;
    height: 80px;
    padding: 16px 32px;
    justify-content: space-between;
    align-items: center;
  }
</style>