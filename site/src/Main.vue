<script setup>

import Menu from './layouts/Header.vue';
import Slider from "./sections/Slider.vue";
import MainBanner from "./sections/MainBanner.vue";

</script>
<template>

<Slider />
<MainBanner />
</template>
<script>
  export default {
    data() {
      return {
        showMenu: false,
      }
    },
  }
</script>
