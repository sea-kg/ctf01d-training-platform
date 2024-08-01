<script>
import Button from '../button/button.vue'
export default {
    props: {
        title: {
            type: String,
            required: true
        },
        action: {
            type: function(){}
        },
    },
    data(){
        return {
            button: {
                title: '',
                addButtonIconPath: '/src/assets/icon/plus-square.svg'
            }
        }
    },
    components: {
        Button
    }
  }
</script>

<template>
  <div class="row">
    <div class="wrapper_page_title">
        <div class="col12">
            <div class="page_title">
                <div class="title semi_40">{{ title }}</div>
                <Button
                    :title=button.title
                    :isSecondary="true"
                    :imgPath=button.addButtonIconPath
                    :action="action"
                />
            </div>
        </div>
    </div>
  </div>
</template>

<style lang="scss">
    @import '/src/styles/vars.scss';

    .wrapper_page_title{
        width: 100%;
        margin-top: 40px;
        margin-bottom: 32px;
    }
    .page_title{
        display: flex;
    }
    .title {
        margin-right: 16px;
        color: $primary;
    }
</style>