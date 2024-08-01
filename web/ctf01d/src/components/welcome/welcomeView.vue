<template>
  <div class="container">
    <div class="row">
      <div class="col6">
        <div class="welcome_wrapper">
          <div class="welcome_title bold_60">Welcome to <span class="accent">ctf01d</span> training platform!</div>
          <div class="welcome_description reg_16">{{ welcomeText.description }}</div>
          <Button
            :title="'see on Github'"
            :isSecondary=true
            :action="toGitHub"
            :imgPath=githubIconPath
          />
        </div>
      </div>
      <div class="col6">
        <img src="/src/assets/img/welcome.svg" alt="" class="hero_img">
      </div>
    </div>
  </div>
  
</template>

<script setup>
  import Button from '../common/button/button.vue'
  const welcomeText = {
    description: 'This service can prepare training attack-defense game, based on ctf01d jury system'
  }
  const toGitHub = () => window.open('https://github.com/sea-kg/ctf01d-training-platform', '_blank').focus();
  const githubIconPath = '/src/assets/icon/github.svg'
</script>

<style lang="scss">
@import '/src/styles/vars.scss';
.welcome_wrapper {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: $primary;
}

.welcome_title, .welcome_description {
  margin-bottom: 24px;
}
.accent {
  color: #C6FF00;
}
.hero_img {
  height: calc(100vh - 100px);
}
</style>