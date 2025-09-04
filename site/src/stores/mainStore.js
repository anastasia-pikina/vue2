import { defineStore } from 'pinia';
import axios from "axios";

export const useMainStore = defineStore('mainStore', {
    state: () => {
        return {
            api_url: 'http://localhost:4000',
        }
    },

    actions: {
        async getData(url) {
            if (!url) {
                return null;
            }

            url = `${this.api_url}${url}`;

            try {
                const response = await axios.get(url);
                return response.data;
            } catch (error) {
                console.error(error);
                return null;
            }
        },
    },
})