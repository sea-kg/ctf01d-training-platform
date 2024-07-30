<template>
  <div class="welcome_wrapper">
    <div class="welcome_title bold_60">Welcome to <span class="accent">ctf01d</span> training platform!</div>
    <div class="welcome_description reg_16">{{ welcomeText.description }}</div>
    <Button
      :title="'see on Github'"
      :isSecondary=true
      :action="toGitHub"
    />
  </div>
</template>

<script setup>
import Button from '../common/button/button.vue'
const welcomeText = {
  title: 'Welcome to <span class="accent">ctf01d</span> training platform!',
  description: 'This service can prepare training attack-defense game, based on ctf01d jury system'
}
const toGitHub = () => window.open('https://github.com/sea-kg/ctf01d-training-platform', '_blank').focus();
</script>

<style>
.welcome_wrapper {
  margin-left: 117px;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.welcome_title, .welcome_description {
  margin-bottom: 24px;
}
.accent {
  color: #C6FF00;
}
</style>