<script setup>
import { ref, onMounted } from 'vue'
import {useMainStore} from './stores/mainStore';
const store = useMainStore();
const contacts = ref([]);

onMounted(async () => {
  contacts.value = await store.getData('/contacts') ?? [];
});
</script>

<template>
  <div class="banner-2 md:flex items-center md:justify-evenly" id="about">
    <img
        class="md:h-[500px] h-[400px]"
        src="./assets/about-female.png"
        alt="user image"
    />
    <div class="space-y-5 py-8 px-8 md:py-16 md:px-20 md:w-1/2" v-for="(contact, index) in contacts" :key="index">
      <h4 class="project-title item">{{contact.address}}</h4>
      <p class="font-work_sans" v-for="(phoneItem, phoneIndex) in contact.phone" :key="phoneIndex">
        {{phoneItem}}
      </p>
    </div>
  </div>
</template>

<style scoped>

</style>