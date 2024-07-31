<template>
  <div @click.prevent="goTo()" class="header_menu_item reg_16">{{ item.title }}</div>
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
  .header_menu_item {
    margin-right: 64px;
    border-bottom: 2px solid $surface_1;
    padding-bottom: 4px;
    color: $primary;
    &:hover {
      border-bottom: 2px solid $accent;
      cursor: pointer;
      transition: none;
    }
  }
</style>