<script>
import Button from '../common/button/button.vue'
import TeamListItem from './game-team-item.vue'
import axios from 'axios'
import moment from 'moment';

export default {
    props: ['game_data'],
    data(){
      return{
        game: '',
        team_length: ''
      }
    },
    components: {
        Button,
        TeamListItem
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
  <div class="col4">
    <div class="game_wrapper">
        <div class="game_title_wrapper">
            <div class="game_title semi_24">{{ game.description }}</div>
            <div class="game_subtitle reg_14">{{ dateStatus }}</div>
        </div>
        <div class="game_due_date_wrapper">
            <div class="game_time_wrapper game_time_start game_time_start_active">
                <div class="game_time_status reg_14">Start</div>
                <div class="game_timing semi_16">
                    <div class="game_date">{{ startDate }}</div>
                    <div class="game_time">{{ startTime }}</div>
                </div>
            </div>
            <img src="/src/assets/icon/chevrons-right.svg" alt="" class="game_time_icon">
            <div class="game_time_wrapper game_time_finish">
                <div class="game_time_status reg_14">Finish</div>
                <div class="game_timing semi_16">
                    <div class="game_date">{{ finishDate }}</div>
                    <div class="game_time">{{ finishTime }}</div>
                </div>
            </div>
        </div>
        <div class="game_teams_wrapper" v-if="team_length > 0">
            <div class="game_teams_title reg_14">teams [{{ team_length }}]</div>
            <div class="game_teams_list">
                <TeamListItem
                    v-for="team in game.team_details"
                    :key="team.id"
                    :team="team"
                />
                <div class="list_more" v-if="false">
                    <Button
                        :title="'see more'"
                        :isGhost="true"
                    />
                </div>
            </div>
        </div>
    </div>
  </div>
</template>

<style lang="scss">
    @import '/src/styles/vars.scss';

    .game_wrapper {
        border-radius: 8px;
        background-color: $surface-2;
        width: 100%;
        padding: 32px 16px;
        margin-bottom: 24px;
    }

    .game_title_wrapper {
        margin: 0 8px;
        margin-bottom: 24px;
    }

    .game_title {
        color: $accent;
        margin-bottom: 4px;
    }

    .game_subtitle {
        color: $primary;
    }

    .game_due_date_wrapper {
        margin: 0 8px;
        margin-bottom: 24px;
        display: flex;
    }

    .game_time_wrapper{
        padding: 8px 12px;
        border-radius: 8px;
    }

    .game_time_status {
        margin-bottom: 8px;
    }

    .game_time_start {
        background-color: $surface-1;
        color: $tertiary;
    }

    .game_time_start_active {
        background-color: $accent;
        color: $on-accent;
    }

    .game_time_finish {
        border: 1px solid $outline;
        color: $primary;
    }

    .game_time_finish_active {
        background-color: $surface-1;
    }

    .game_time_icon {
        margin: 0 24px;
    }

    .game_time_icon_disabled {
        opacity: .1;
    }

    .game_teams_title {
        color: $secondary;
        margin: 0 8px;
        margin-bottom: 4px;
    }

    

    .list_more {
        padding-top: 16px;
        width: 100%;
        display: flex;
        justify-content: center;
    }

</style>