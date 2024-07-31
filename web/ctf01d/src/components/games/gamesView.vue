<script>
  import PageTitle from '../common/pageTitle/page_title.vue'
  import Game from './game-item.vue'
  import axios from 'axios'

  export default {
    data(){
      return{
        pageTitle: "Games",
        games: []
      }
    },
    components: {
      PageTitle,
      Game
    },
    methods: {
      getGames(){
        axios.get(`http://ctf01d.ru:4102/api/v1/games`)
            .then(response => (
                this.games = response.data
            ))
      }
    },
    mounted(){
        this.getGames()
    },
  }
</script>

<template>
  <div class="container">
    <PageTitle :title="pageTitle"/>
    <div class="row">
      <Game
        v-for="game in games"
        :key="game.id"
        :game_data="game"
      />
    </div>
  </div>
</template>

<style>

</style>