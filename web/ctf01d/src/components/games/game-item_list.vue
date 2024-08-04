<script>
// import axios from 'axios';
import axios from 'https://cdn.jsdelivr.net/npm/axios@1.3.5/+esm';

import TeamAvatarList from '../games/game-team-item_list.vue'
import ModalHeader from '../common/modal/modal_header.vue'
import Button from '../common/button/button.vue'

export default {
    props: ['game_data'],
    data(){
      return{
        game: '',
        team_length: '',
        showEditModal: false,
        showDeleteModal: false,
        modalEditHeaderText: {
          title: "Edit game"
        },
        modalDeleteHeaderText: {
          title: "Delete game?"
        },
        editGameObj: {
          start_time: '',
          end_time: '',
          description: ''
        },
        minStartDate: '',
        showAlertEdit: false,
        alertEditData: ''
      }
    },
    components: {
      TeamAvatarList,
      ModalHeader,
      Button
    },
    methods: {
      getGame(){
        axios.get(`https://ctf01d.ru/api/v1/games/${this.game_data.id}`)
            .then(response => (
                this.game = response.data,
                this.team_length = this.game.team_details.length
            ))
      },
      editGame(){
          
        const game = {
          start_time: this.editGameObj.start_time+':00.000Z',
          end_time: this.editGameObj.end_time+':00.000Z',
          description: this.editGameObj.description
        }
        axios.put(`https://ctf01d.ru/api/v1/games/${this.game_data.id}`, game)
            .then(response => (
                this.alertEditData = response.data.data,
                this.toggleModalEdit(),
                this.$emit("updateGameList"),
                this.toggleAlertEdit()
            ))
      },
      deleteGame(){
        axios.delete(`https://ctf01d.ru/api/v1/games/${this.game_data.id}`)
            .then(response => (
              this.$emit("updateGameList")
            ))
      },
      toggleModalEdit() {
        if(!this.showEditModal) {
          this.showEditModal = true
        } else {
          this.showEditModal = false
        }
      },
      toggleModalDelete() {
        if(!this.showDeleteModal) {
          this.showDeleteModal = true
        } else {
          this.showDeleteModal = false
        }
      },
      formatedObj(){
        this.editGameObj.start_time = this.game_data.start_time.slice(0, 16)
        this.editGameObj.end_time = this.game_data.end_time.slice(0, 16)
        this.editGameObj.description = this.game_data.description
      },
      formatedMinStart(){
        let date = new Date()
        this.minStartDate = date.toISOString().slice(0, 16)
      },
      toggleAlertEdit(){
        if(!this.showAlertEdit) {
          this.showAlertEdit = true
          setTimeout(() => {
            this.showAlertEdit = false
          }, 5000);
        } else {
          this.showAlertEdit = false
        }
      },
    },
    mounted(){
        this.getGame(),
        this.formatedMinStart(),
        this.formatedObj()
    },
    computed: {
        dateStatus () {
            return this.game.end_time;
        },
        startDate(){
            return this.game.start_time;
        },
        startTime(){
            return this.game.start_time;
        },
        finishDate(){
            return this.game.end_time;
        },
        finishTime(){
            return this.game.end_time;
        }
    }
}
</script>

<template>
  <div class="alert reg_16" :class="{showAlert: showAlertEdit}">{{ alertEditData }}</div>
  <div class="modal_wrapper" v-if="showEditModal">
    <div class="modal_create_game">
      <ModalHeader
        :text="modalEditHeaderText"
        :action="toggleModalEdit"
      />
      <form @submit.prevent="editGame" class="modal_content">
        <div class="form_control">
          <label class="reg_16" for="start_time">Start time</label>
          <input
            type="datetime-local"
            name="start_time"
            v-model="editGameObj.start_time"
            :min="minStartDate"
            required
          />
        </div>
        <div class="form_control">
          <label class="reg_16" for="end_time">Finish time</label>
          <input
            type="datetime-local"
            name="end_time"
            v-model="editGameObj.end_time"
            :min="editGameObj.start_time"
            required
          />
        </div>
        <div class="form_control">
          <label class="reg_16" for="description">Description</label>
          <textarea v-model="editGameObj.description" name="description" id="description" cols="30" rows="2" placeholder="Couple words" required></textarea>
        </div>
        <div class="form_submit_wrapper">
          <input type="submit" value="save changes">
        </div>
      </form>
    </div>
  </div>

  <div class="modal_wrapper" v-if="showDeleteModal">
    <div class="modal_create_game">
      <ModalHeader
        :text="modalDeleteHeaderText"
        :action="toggleModalDelete"
      />
      <form @submit.prevent="deleteGame" class="modal_content">
        <div class="form_submit_wrapper">
          <input type="submit" value="Yes, delete">
          <Button
            :isSecondary="true"
            :title="'No, leave it'"
            :action="toggleModalDelete"
          />
        </div>
      </form>
    </div>
  </div>

  <div class="overlay" v-if="showEditModal"
    @click.prevent="toggleModalEdit"
  ></div>
  <div class="overlay" v-if="showDeleteModal"
    @click.prevent="toggleModalDelete"
  ></div>


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
    <div class="game_list_item_actions">
      <img @click.prevent="toggleModalEdit()" src="/src/assets/icon/edit.svg" alt="" class="action_button action_edit">
      <img @click.prevent="toggleModalDelete()" src="/src/assets/icon/trash.svg" alt="" class="action_button action_delete">
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

    .game_list_item_actions {
      display: flex;
    }

    .action_button {
      padding: 4px;
      border-radius: 4px;
      &:hover {
        background-color: $surface-3;
      }
    }

    .action_edit {
      margin-right: 6px;
    }

    .overlay {
      width: 100%;
      height: 100%;
      position: fixed;
      background-color: rgba(0, 0, 0, .7);
      z-index: 98;
      left: 0;
      top: 0;
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

    .form_submit_wrapper {
      display: flex;
    }

</style>
