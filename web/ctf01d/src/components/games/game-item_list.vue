<script>
import axios from 'axios'
import moment from 'moment';

import TeamAvatarList from '../games/game-team-item_list.vue'

export default {
    props: ['game_data'],
    data(){
      return{
        game: '',
        team_length: ''
      }
    },
    components: {
      TeamAvatarList
    },
    methods: {
      getGame(){
        axios.get(`https://ctf01d.ru/api/v1/games/${this.game_data.id}`)
            .then(response => (
                this.game = response.data,
                this.team_length = this.game.team_details.length
            ))
      }
    },
    mounted(){
        this.getGame()
    },
    computed: {
        dateStatus () {
            return moment(this.game.end_time).fromNow();
        },
        startDate(){
            return moment(this.game.start_time).format('LL');;
        },
        startTime(){
            return moment(this.game.start_time).format('LT');;
        },
        finishDate(){
            return moment(this.game.end_time).format('LL');;
        },
        finishTime(){
            return moment(this.game.end_time).format('LT');;
        }
    }
}
</script>

<template>
  <div class="game_list_item">
    <div class="game_list_item_description semi_14">{{ game.description }}</div>
    <div class="game_list_item_start reg_14">
      <div class="item_date">{{ startDate }}</div>
      <div class="item_time">{{ startTime }}</div>
    </div>
    <div class="game_list_item_finish reg_14">
      <div class="item_date">{{ finishDate }}</div>
      <div class="item_time">{{ finishTime }}</div>
    </div>
    <div class="game_list_item_status reg_14">{{ dateStatus }}</div>
    <div class="game_list_item_teams">
      <div class="team_length reg_14">teams: [{{ team_length || 0 }}]</div>
      <div class="teams_avatar_list">
        <TeamAvatarList 
          v-for="team in game.team_details"
          :key="team.id"
          :team="team"
        />
      </div>
    </div>
  </div>
  
</template>

<style lang="scss">
    @import '/src/styles/vars.scss';

    .game_list_item {
      padding: 12px;
      border-bottom: 1px solid $outline;
      display: flex;
      justify-content: space-between;
      align-items: center;
      &:hover {
        background-color: $surface-2;
        cursor: pointer;
      }
    }

    .game_list_item_description {
      width: 360px;
      color: $accent;
    }

    .game_list_item_start,
    .game_list_item_finish,
    .game_list_item_status {
      width: 160px;
      color: $primary;
    }

    .game_list_item_teams {
      width: 136px;
      color: $secondary;
    }

    .teams_avatar_list {
      display: flex;
      width: 100%;
      justify-content: space-between;
      margin-top: 4px;
    }
</style>