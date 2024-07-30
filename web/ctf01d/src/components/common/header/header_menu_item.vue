<template>
  <div class="header_menu_item reg_16">{{ title }}</div>
</template>

<script>
  export default {
    props: {
      title: {
        type: String,
        required: true
      },
    }
  }
</script>

<style lang="scss">
@import '../../../styles/vars.scss';
  .header_menu_item {
    margin-right: 64px;
    border-bottom: 2px solid $surface_1;
    padding-bottom: 4px;
    &:hover {
      border-bottom: 2px solid $accent;
      cursor: pointer;
      transition: none;
    }
  }
</style>