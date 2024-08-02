<script>
  import PageTitle from '../common/pageTitle/page_title.vue'
  import GameItemList from '../games/game-item_list.vue'
  import ModalHeader from '../common/modal/modal_header.vue'
  import Button from '../common/button/button.vue'
  
  import axios from 'axios'

  export default {
    data(){
      return{
        pageTitle: "Games",
        search_game: '',
        games: [],
        showModal: false,
        showAlert: false,
        modalHeaderText: {
          title: "Creating new game"
        },
        newGame: {
          start_time: '',
          end_time: '',
          description: ''
        }
      }
    },
    components: {
      PageTitle,
      GameItemList,
      ModalHeader,
      Button
    },
    methods: {
      getGames(){
        axios.get(`https://ctf01d.ru/api/v1/games`)
            .then(response => (
                this.games = response.data
            ))
      },
      createGame(){
        
        const game = {
          start_time: this.newGame.start_time+':00.000Z',
          end_time: this.newGame.end_time+':00.000Z',
          description: this.newGame.description
        }
        axios.post(`https://ctf01d.ru/api/v1/games`, game)
            .then(response => (
                this.toggleModal(),
                this.getGames(),
                this.toggleAlert()
            ))
      },
      toggleModal() {
        if(!this.showModal) {
          this.showModal = true
        } else {
          this.showModal = false
        }
      },
      toggleAlert(){
        if(!this.showAlert) {
          this.showAlert = true
          setTimeout(() => {
            this.showAlert = false
          }, 5000);
        } else {
          this.showAlert = false
        }
      },
      formatedStart(){
        let date = new Date()
        this.newGame.start_time = date.toISOString().slice(0, 16)
      }
    },
    mounted(){
        this.getGames()
        this.formatedStart()
    },
    computed: {
      searchGames() {
            if(this.search_game){
                return this.games.filter(item => {
                    return item.description.toLowerCase().includes(this.search_game.toLowerCase())
                })
            }
            return this.games
        },
    }
  }
</script>

<template>
  <div class="alert reg_16" :class="{showAlert: showAlert}">Game created successfully</div>
  <div class="modal_wrapper" v-if="showModal">
    <div class="modal_create_game">
      <ModalHeader
        :text="modalHeaderText"
        :action="toggleModal"
      />
      <form @submit.prevent="createGame" class="modal_content">
        <div class="form_control">
          <label class="reg_16" for="start_time">Start time</label>
          <input
            type="datetime-local"
            name="start_time"
            v-model="newGame.start_time"
            :min="newGame.start_time"
            required
          />
        </div>
        <div class="form_control">
          <label class="reg_16" for="end_time">Finish time</label>
          <input
            type="datetime-local"
            name="end_time"
            v-model="newGame.end_time"
            :min="newGame.start_time"
            required
          />
        </div>
        <div class="form_control">
          <label class="reg_16" for="description">Description</label>
          <textarea v-model="newGame.description" name="description" id="description" cols="30" rows="2" placeholder="Couple words" required></textarea>
        </div>
        <div class="form_submit_wrapper">
          <input type="submit" value="create now">
        </div>
      </form>
    </div>
  </div>
  <div class="overlay" v-if="showModal"
    @click.prevent="toggleModal"
  ></div>
  <div class="container">
    <PageTitle
      :title="pageTitle"
      :action="toggleModal"
    />
    <div class="row">
      <div class="col6">
        <input type="text" name="search_game" id="search_game" v-model="search_game" placeholder="Search game">
      </div>
    </div>
    <div class="row">
      <div class="col12">
        <div class="games_list">
          <GameItemList
            v-for="game in searchGames"
            :key="game.id"
            :game_data="game"
            @updateGameList="getGames"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
    @import '/src/styles/vars.scss';

    .overlay {
      width: 100%;
      height: 100%;
      position: fixed;
      background-color: rgba(0, 0, 0, .7);
      z-index: 98;
    }

    .modal_wrapper {
      position: fixed;
      z-index: 99;
      top: 20%;
      left: 33%;
      width: 500px;
      background-color: $surface-2;
    }

    .modal_content {
      padding: 0 32px 32px 32px;
    }

    .form_control {
      margin-bottom: 24px;
    }

    .games_list {
      width: 100%;
    }

    .alert {
      position: fixed;
      right: 32px;
      top: -80px;
      background-color: #4bca86;
      z-index: 100;
      padding: 12px 24px;
      color: $on-accent;
      min-width: 600px;
    }

    .showAlert {
      top: 80px;
    }

</style>