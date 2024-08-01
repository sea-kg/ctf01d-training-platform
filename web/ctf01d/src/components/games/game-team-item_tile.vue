<template>
    <div class="game_teams_item">
        <img :src="`https://ctf01d.ru/api/v1/avatar/${this.team.name}`" alt="" class="team_avatar">
        <div class="game_team_info">
            <div class="team_title semi_14">{{ team.name }}</div>
            <div class="team_description reg_14">{{ team.description }}</div>
        </div>
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

    .game_teams_item {
        padding: 16px 8px;
        display: flex;
        border-bottom: 1px solid $outline;
        &:hover {
            background-color: $surface-3;
            border-bottom: 1px solid $surface-3;
            border-radius: 8px;
            cursor: pointer;
        }
    }

    .game_team_info {
        margin-left: 16px;
    }

    .team_avatar {
        object-fit: contain;
        width: 42px;
        height: 42px;
    }

    .team_title {
        color: $primary;
    }

    .team_description {
        color: $secondary;
    }
</style>