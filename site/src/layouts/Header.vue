<script setup>
    import { ref } from 'vue';
    import Menu from "../sections/Menu.vue";

    let mobileNav = ref(false);

    let toggleMobileNav = () => {
        mobileNav.value = !mobileNav.value
    }
</script>

<template>
  <div class="banner-1">
    <Menu />
  </div>
</template>