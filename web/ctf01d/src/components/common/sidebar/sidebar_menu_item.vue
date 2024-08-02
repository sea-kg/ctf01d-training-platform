<template>
  <div @click.prevent="goTo()" :class="{ 'sidebar_menu_item_active': $route.path === this.item.path }" class="sidebar_menu_item reg_16">{{ item.title }}</div>
</template>

<script>
  export default {
    props: {
      item: {
        type: Object,
        required: true
      }
    },
    methods: {
      goTo(){
        this.$router.push(`${this.item.path}`)
      },
    }
  }
</script>

<style lang="scss">
@import '../../../styles/vars.scss';
  .sidebar_menu_item {
    margin-bottom: 16px;
    border-bottom: 2px solid $surface_2;
    padding-bottom: 4px;
    color: $primary;
    &:hover {
      border-bottom: 2px solid $accent;
      cursor: pointer;
      transition: none;
    }
  }
  .sidebar_menu_item_active {
    border-bottom: 2px solid $accent;
  }
</style>