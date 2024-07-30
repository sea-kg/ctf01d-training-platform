<template>
  <button @click.prevent="action()"
    class="button button_M semi_16"
    :class="{ primary: isPrimary, secondary: isSecondary}"
  >{{ title }}
    <img class="button_icon" src="../../../assets/icon/log-in.svg" alt="">
  </button>
</template>

<script>
export default {
    props: {
      title: {
        type: String,
        required: true
      },
      isPrimary: {
        type: Boolean
      },
      isSecondary: {
        type: Boolean
      },
      action: {
        type: function(){}
      }
    }
  }
</script>

<style lang="scss">
  @import '../../../styles/vars.scss';
  .button{
    display: flex;
    align-items: center;
    width: fit-content;
  }
  .button_M {
    padding: 12px 24px;
    &:hover {
      cursor: pointer;
    }
  }
  .primary {
    border: none;
    background-color: $accent;
    color: $on-accent;
    &:hover {
      background-color: #E8FF99;
    }
  }
  .secondary {
    border: none;
    background-color: $primary;
    color: $on-accent;
    &:hover {
      background-color: #e2e2e2;
    }
  }
  .button_icon {
    margin-left: 6px;
  }
</style>