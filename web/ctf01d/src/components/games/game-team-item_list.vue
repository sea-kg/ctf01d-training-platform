<template>
    <div class="game_teams_item_list">
        <img :src="`https://ctf01d.ru/api/v1/avatar/${this.team.name}`" alt="" class="team_avatar">
        <div class="tooltip reg_16">{{ team.name }}</div>
    </div>
</template>

<script>
export default {
    props: ["team"],
    data(){
      return{
        avatar: ''
      }
    }
}
</script>

<style lang="scss">
    @import '/src/styles/vars.scss';

    .game_teams_item_list {
        position: relative;
        cursor: pointer;
        &:hover .tooltip {
            display: block;
        }
    }

    .game_teams_item_list img {
        width: 22px;
        height: 22px;
    }

    .tooltip {
        background-color: $surface-3;
        border-radius: 4px;
        padding: 0 4px;
        max-width: 120px;
        position: absolute;
        right: -11px;
        bottom: -22px;
        display: none;
        z-index: 999;
        text-align: left;
    }

</style>