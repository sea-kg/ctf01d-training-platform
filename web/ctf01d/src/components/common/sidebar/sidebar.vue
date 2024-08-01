<template>
  <div class="sidebar_wrapper">
    <div class="sidebar_top">
      <div class="sidebar_logo">
        <RouterLink to="/">
          <img class="logo" src="/src/assets/img/header_logo.svg" alt="">
        </RouterLink>
      </div>
      <Menu />
    </div>
    <div class="sidebar_profile_info">
      <Button 
        :title="'sign in'"
        :isPrimary="true"
        :imgPath=loginButtonIconPath
      />
    </div>
  </div>
</template>

<script setup>
  import Menu from './sidebar_menu.vue'
  import Button from '../button/button.vue'
  const loginButtonIconPath = '/src/assets/icon/log-in.svg'
</script>

<style lang="scss">
  @import '/src/styles/vars.scss';

  .sidebar_logo {
    margin-bottom: 40px;
  }

  .sidebar_wrapper {
    display: flex;
    flex-direction: column;
    width: 200px;
    height: 100vh;
    padding: 16px;
    justify-content: space-between;
    position: fixed;
    left: 0;
    top: 0;
    background-color: $surface-2;
  }
</style>