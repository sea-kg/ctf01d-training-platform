<script setup>
  import Item from './header_menu_item.vue'
  const list = ['games', 'services', 'teams', 'users']
</script>

<template>
  <div class="header_menu">
    <Item 
      v-for="(item, index) in list"
      :key="index"
      :title="item"
      />
  </div>
</template>

<style>
  .header_menu{
    display: flex;
  }
</style>