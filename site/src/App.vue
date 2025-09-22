<script setup>
import Header from './layouts/Header.vue';
import Footer from './layouts/Footer.vue';
</script>
<template>
  <Header />
  <transition name="fade">
  <router-view />
    </transition>
  <Footer />
</template>
<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>