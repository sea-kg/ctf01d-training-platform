<script setup>
  import Item from './header_menu_item.vue'
  let list = [
    {
      title: 'games',
      path: '/games'
    },
    {
      title: 'services',
      path: '/services'
    },
    {
      title: 'teams',
      path: '/teams'
    },
    {
      title: 'users',
      path: '/users'
    },
  ]
</script>

<template>
  <div class="header_menu">
    <Item 
      v-for="item in list"
      :key="item.name"
      :item="item"
      />
  </div>
</template>

<style>
  .header_menu{
    display: flex;
  }
</style>