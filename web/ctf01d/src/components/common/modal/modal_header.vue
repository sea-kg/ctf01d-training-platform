<script>
  export default {
    props: ["text", "action"],
    data(){
      return{
        
      }
    }
  }
</script>

<template>
  <div class="modal_header">
    <div class="modal_header_text">
      <div class="modal_header_title semi_24">{{ text.title }}</div>
      <div class="modal_header_subtitle reg_14">{{ text.subtitle }}</div>
    </div>
    <div class="close_icon" @click.prevent="action">
      <img src="/src/assets/icon/x-square.svg" alt="">
    </div>
  </div>
</template>

<style lang="scss">
  @import '/src/styles/vars.scss';

  .modal_header {
    padding: 32px;
    background-color: $surface-2;
    display: flex;
    justify-content: space-between;
    min-width: 360px;
    max-width: 800px;
  }

  .modal_header_title {
    color: $primary;
  }

  .modal_header_subtitle {
    color: $secondary;
  }

  .close_icon img {
    cursor: pointer;
  }

</style>